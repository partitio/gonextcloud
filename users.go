package gonextcloud

import (
	"encoding/json"
	"github.com/fatih/structs"
	req "github.com/levigross/grequests"
	"github.com/pkg/errors"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"
)

// UserList return the Nextcloud'user list
func (c *Client) UserList() ([]string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.users, nil)
	//res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

//UserListDetails return a map of user with details
func (c *Client) UserListDetails() (map[string]types.User, error) {
	res, err := c.baseRequest(http.MethodGet, routes.users, nil, "details")
	//res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.UserListDetailsResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

// User return the details about the specified user
func (c *Client) User(name string) (*types.User, error) {
	if name == "" {
		return nil, &types.APIError{Message: "name cannot be empty"}
	}
	res, err := c.baseRequest(http.MethodGet, routes.users, nil, name)
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

// UserSearch returns the users whose name match the search string
func (c *Client) UserSearch(search string) ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"search": search},
	}
	res, err := c.baseRequest(http.MethodGet, routes.users, ro)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

// UserCreate create a new user
func (c *Client) UserCreate(username string, password string, user *types.User) error {
	// Create base User
	ro := &req.RequestOptions{
		Data: map[string]string{
			"userid":   username,
			"password": password,
		},
	}
	if err := c.userBaseRequest(http.MethodPost, ro); err != nil {
		return err
	}
	// Check if we need to add user details information
	if user == nil {
		return nil
	}
	// Add user details information
	return c.UserUpdate(user)
}

// UserCreateWithoutPassword create a user without provisioning a password, the email address must be provided to send
// an init password email
func (c *Client) UserCreateWithoutPassword(username, email, displayName string) error {
	if c.version.Major < 14 {
		return errors.New("unsupported method: requires Nextcloud 14+")
	}
	if username == "" || email == "" {
		return errors.New("username and email cannot be empty")
	}
	ro := &req.RequestOptions{
		Data: map[string]string{
			"userid":      username,
			"email":       email,
			"displayName": displayName,
		},
	}

	if err := c.userBaseRequest(http.MethodPost, ro); err != nil {
		return err
	}
	return nil
}

//UserDelete delete the user
func (c *Client) UserDelete(name string) error {
	return c.userBaseRequest(http.MethodDelete, nil, name)
}

//UserEnable enables the user
func (c *Client) UserEnable(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	return c.userBaseRequest(http.MethodPut, ro, name, "enable")
}

//UserDisable disables the user
func (c *Client) UserDisable(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	return c.userBaseRequest(http.MethodPut, ro, name, "disable")
}

//UserSendWelcomeEmail (re)send the welcome mail to the user (return an error if the user has not configured his email)
func (c *Client) UserSendWelcomeEmail(name string) error {
	return c.userBaseRequest(http.MethodPost, nil, name, "welcome")
}

//UserUpdate takes a *types.User struct to update the user's information
func (c *Client) UserUpdate(user *types.User) error {
	m := structs.Map(user)
	errs := make(chan types.UpdateError)
	var wg sync.WaitGroup
	for k := range m {
		if !ignoredUserField(k) && m[k].(string) != "" {
			wg.Add(1)
			go func(key string, value string) {
				defer wg.Done()
				if err := c.userUpdateAttribute(user.ID, strings.ToLower(key), value); err != nil {
					errs <- types.UpdateError{
						Field: key,
						Error: err,
					}
				}
			}(k, m[k].(string))
		}
	}
	go func() {
		wg.Wait()
		close(errs)
	}()
	return types.NewUpdateError(errs)
}

//UserUpdateEmail update the user's email
func (c *Client) UserUpdateEmail(name string, email string) error {
	return c.userUpdateAttribute(name, "email", email)
}

//UserUpdateDisplayName update the user's display name
func (c *Client) UserUpdateDisplayName(name string, displayName string) error {
	return c.userUpdateAttribute(name, "displayname", displayName)
}

//UserUpdatePhone update the user's phone
func (c *Client) UserUpdatePhone(name string, phone string) error {
	return c.userUpdateAttribute(name, "phone", phone)
}

//UserUpdateAddress update the user's address
func (c *Client) UserUpdateAddress(name string, address string) error {
	return c.userUpdateAttribute(name, "address", address)
}

//UserUpdateWebSite update the user's website
func (c *Client) UserUpdateWebSite(name string, website string) error {
	return c.userUpdateAttribute(name, "website", website)
}

//UserUpdateTwitter update the user's twitter
func (c *Client) UserUpdateTwitter(name string, twitter string) error {
	return c.userUpdateAttribute(name, "twitter", twitter)
}

//UserUpdatePassword update the user's password
func (c *Client) UserUpdatePassword(name string, password string) error {
	return c.userUpdateAttribute(name, "password", password)
}

//UserUpdateQuota update the user's quota (bytes)
func (c *Client) UserUpdateQuota(name string, quota int) error {
	return c.userUpdateAttribute(name, "quota", strconv.Itoa(quota))
}

//UserGroupList lists the user's groups
func (c *Client) UserGroupList(name string) ([]string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.users, nil, name, "groups")
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//UserGroupAdd adds a the user to the group
func (c *Client) UserGroupAdd(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(http.MethodPost, ro, name, "groups")
}

//UserGroupRemove removes the user from the group
func (c *Client) UserGroupRemove(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(http.MethodDelete, ro, name, "groups")
}

//UserGroupPromote promotes the user as group admin
func (c *Client) UserGroupPromote(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(http.MethodPost, ro, name, "subadmins")
}

//UserGroupDemote demotes the user
func (c *Client) UserGroupDemote(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(http.MethodDelete, ro, name, "subadmins")
}

//UserGroupSubAdminList lists the groups where he is subadmin
func (c *Client) UserGroupSubAdminList(name string) ([]string, error) {
	if !c.loggedIn() {
		return nil, errUnauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name, "subadmins")
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.BaseResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

func (c *Client) userUpdateAttribute(name string, key string, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"key":   key,
			"value": value,
		},
	}
	return c.userBaseRequest(http.MethodPut, ro, name)
}

func (c *Client) userBaseRequest(method string, ro *req.RequestOptions, subRoutes ...string) error {
	_, err := c.baseRequest(method, routes.users, ro, subRoutes...)
	return err
}

func ignoredUserField(key string) bool {
	keys := []string{"ID", "Quota", "Enabled", "Groups", "Language"}
	for _, k := range keys {
		if key == k {
			return true
		}
	}
	return false
}
