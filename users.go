package gonextcloud

import (
	"encoding/json"
	"github.com/fatih/structs"
	req "github.com/levigross/grequests"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
)

//Users contains all Users available actions
type Users struct {
	c *Client
}

// List return the Nextcloud'user list
func (u *Users) List() ([]string, error) {
	res, err := u.c.baseRequest(http.MethodGet, routes.users, nil)
	//res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

//ListDetails return a map of user with details
func (u *Users) ListDetails() (map[string]types.UserDetails, error) {
	res, err := u.c.baseRequest(http.MethodGet, routes.users, nil, "details")
	//res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.UserListDetailsResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

// Get return the details about the specified user
func (u *Users) Get(name string) (*types.UserDetails, error) {
	if name == "" {
		return nil, &types.APIError{Message: "name cannot be empty"}
	}
	res, err := u.c.baseRequest(http.MethodGet, routes.users, nil, name)
	if err != nil {
		return nil, err
	}
	var r types.UserResponse
	js := res.String()
	// Nextcloud does not encode JSON properly
	js = reformatJSON(js)
	if err := json.Unmarshal([]byte(js), &r); err != nil {
		return nil, err
	}
	return &r.Ocs.Data, nil
}

// Search returns the users whose name match the search string
func (u *Users) Search(search string) ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"search": search},
	}
	res, err := u.c.baseRequest(http.MethodGet, routes.users, ro)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

// Create create a new user
func (u *Users) Create(username string, password string, user *types.UserDetails) error {
	// Create base Users
	ro := &req.RequestOptions{
		Data: map[string]string{
			"userid":   username,
			"password": password,
		},
	}
	if err := u.baseRequest(http.MethodPost, ro); err != nil {
		return err
	}
	// Check if we need to add user details information
	if user == nil {
		return nil
	}
	// Add user details information
	return u.Update(user)
}

// CreateWithoutPassword create a user without provisioning a password, the email address must be provided to send
// an init password email
func (u *Users) CreateWithoutPassword(username, email, displayName, quota, language string, groups ...string) error {
	if u.c.version.Major < 14 {
		return errors.New("unsupported method: requires Nextcloud 14+")
	}
	if username == "" || email == "" {
		return errors.New("username and email cannot be empty")
	}

	data := map[string]string{
		"userid":      username,
		"email":       email,
		"displayName": displayName,
		"quota":       quota,
		"language":    language,
	}

	ro := &req.RequestOptions{}
	f := url.Values{}
	for k, v := range data {
		if v != "" {
			f.Add(k, v)
		}
	}
	for _, g := range groups {
		f.Add("groups[]", g)
	}
	ro.RequestBody = strings.NewReader(f.Encode())
	ro.Headers = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	if err := u.baseRequest(http.MethodPost, ro); err != nil {
		return err
	}
	return nil
}

//CreateBatchWithoutPassword create multiple users and send them the init password email
func (u *Users) CreateBatchWithoutPassword(users []types.User) error {
	var wg sync.WaitGroup
	errs := make(chan error)
	for _, us := range users {
		wg.Add(1)
		go func(user types.User) {
			logrus.Debugf("creating user %s", user.Username)
			defer wg.Done()
			if err := u.CreateWithoutPassword(
				user.Username, user.Email, user.DisplayName, "", "", user.Groups...,
			); err != nil {
				errs <- err
			}
		}(us)
	}
	go func() {
		wg.Wait()
		close(errs)
	}()
	var es []error
	for err := range errs {
		es = append(es, err)
	}
	if len(es) > 0 {
		return errors.Errorf("errors occurred while creating users: %v", es)
	}
	return nil
}

//Delete delete the user
func (u *Users) Delete(name string) error {
	return u.baseRequest(http.MethodDelete, nil, name)
}

//Enable enables the user
func (u *Users) Enable(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	return u.baseRequest(http.MethodPut, ro, name, "enable")
}

//Disable disables the user
func (u *Users) Disable(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	return u.baseRequest(http.MethodPut, ro, name, "disable")
}

//SendWelcomeEmail (re)send the welcome mail to the user (return an error if the user has not configured his email)
func (u *Users) SendWelcomeEmail(name string) error {
	return u.baseRequest(http.MethodPost, nil, name, "welcome")
}

//Update takes a *types.Users struct to update the user's information
func (u *Users) Update(user *types.UserDetails) error {
	m := structs.Map(user)
	errs := make(chan *types.UpdateError)
	var wg sync.WaitGroup
	for k := range m {
		// Filter updatable fields
		if ignoredUserField(k) {
			continue
		}
		var value string
		// Quota is a special case
		if k == "Quota" {
			// If empty
			if user.Quota == (types.Quota{}) {
				value = "default"
			} else {
				value = user.Quota.String()
			}
		} else {
			value = m[k].(string)
		}
		if value != "" {
			wg.Add(1)
			// All other non ignored values are strings
			go func(key string, value string) {
				defer wg.Done()
				if err := u.updateAttribute(user.ID, strings.ToLower(key), value); err != nil {
					errs <- &types.UpdateError{
						Field: key,
						Error: err,
					}
				}
			}(k, value)
		}
	}
	go func() {
		wg.Wait()
		close(errs)
	}()
	// Warning : we actually need to check the *err
	if err := types.NewUpdateError(errs); err != nil {
		return err
	}
	return nil
}

//UpdateEmail update the user's email
func (u *Users) UpdateEmail(name string, email string) error {
	return u.updateAttribute(name, "email", email)
}

//UpdateDisplayName update the user's display name
func (u *Users) UpdateDisplayName(name string, displayName string) error {
	return u.updateAttribute(name, "displayname", displayName)
}

//UpdatePhone update the user's phone
func (u *Users) UpdatePhone(name string, phone string) error {
	return u.updateAttribute(name, "phone", phone)
}

//UpdateAddress update the user's address
func (u *Users) UpdateAddress(name string, address string) error {
	return u.updateAttribute(name, "address", address)
}

//UpdateWebSite update the user's website
func (u *Users) UpdateWebSite(name string, website string) error {
	return u.updateAttribute(name, "website", website)
}

//UpdateTwitter update the user's twitter
func (u *Users) UpdateTwitter(name string, twitter string) error {
	return u.updateAttribute(name, "twitter", twitter)
}

//UpdatePassword update the user's password
func (u *Users) UpdatePassword(name string, password string) error {
	return u.updateAttribute(name, "password", password)
}

//UpdateQuota update the user's quota (bytes). Set negative quota for unlimited
func (u *Users) UpdateQuota(name string, quota int64) error {
	q := types.Quota{Quota: quota}
	return u.updateAttribute(name, "quota", q.String())
}

//GroupList lists the user's groups
func (u *Users) GroupList(name string) ([]string, error) {
	res, err := u.c.baseRequest(http.MethodGet, routes.users, nil, name, "groups")
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//GroupAdd adds a the user to the group
func (u *Users) GroupAdd(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return u.baseRequest(http.MethodPost, ro, name, "groups")
}

//GroupRemove removes the user from the group
func (u *Users) GroupRemove(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return u.baseRequest(http.MethodDelete, ro, name, "groups")
}

//GroupPromote promotes the user as group admin
func (u *Users) GroupPromote(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return u.baseRequest(http.MethodPost, ro, name, "subadmins")
}

//GroupDemote demotes the user
func (u *Users) GroupDemote(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return u.baseRequest(http.MethodDelete, ro, name, "subadmins")
}

//GroupSubAdminList lists the groups where he is subadmin
func (u *Users) GroupSubAdminList(name string) ([]string, error) {
	if !u.c.loggedIn() {
		return nil, errUnauthorized
	}
	ur := u.c.baseURL.ResolveReference(routes.users)
	ur.Path = path.Join(ur.Path, name, "subadmins")
	res, err := u.c.session.Get(ur.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.BaseResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

func (u *Users) updateAttribute(name string, key string, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"key":   key,
			"value": value,
		},
	}
	return u.baseRequest(http.MethodPut, ro, name)
}

func (u *Users) baseRequest(method string, ro *req.RequestOptions, subRoutes ...string) error {
	_, err := u.c.baseRequest(method, routes.users, ro, subRoutes...)
	return err
}

func ignoredUserField(key string) bool {
	keys := []string{"Email", "Displayname", "Phone", "Address", "Website", "Twitter", "Quota"}
	for _, k := range keys {
		if key == k {
			return false
		}
	}
	return true
}
