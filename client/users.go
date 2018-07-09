package client

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/client/types"
	"net/http"
	"path"
)

func (c *Client) UserList() ([]string, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var ul types.UserListResponse
	res.JSON(&ul)
	return ul.Ocs.Data.Users, nil
}

func (c *Client) User(name string) (*types.User, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name)
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return nil, fmt.Errorf("%d : %s", ur.Ocs.Meta.Statuscode, ur.Ocs.Meta.Message)
	}
	return &ur.Ocs.Data, nil
}

func (c *Client) UserSearch(search string) ([]string, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	ro := &req.RequestOptions{
		Params: map[string]string{"search": search},
	}
	res, err := c.session.Get(u.String(), ro)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	if r.Ocs.Meta.Statuscode != 100 {
		return nil, fmt.Errorf("%d : %s", r.Ocs.Meta.Statuscode, r.Ocs.Meta.Message)
	}
	return r.Ocs.Data.Users, nil
}

func (c *Client) UserCreate(username string, password string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"userid":   username,
			"password": password,
		},
	}
	return c.userBaseRequest("", "", ro, http.MethodPost)
}

func (c *Client) UserDelete(name string) error {
	return c.userBaseRequest(name, "", nil, http.MethodDelete)
}

func (c *Client) UserEnable(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	return c.userBaseRequest(name, "enable", ro, http.MethodPut)
}

func (c *Client) UserDisable(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	return c.userBaseRequest(name, "disable", ro, http.MethodPut)
}

func (c *Client) UserSendWelcomeEmail(name string) error {
	return c.userBaseRequest(name, "welcome", nil, http.MethodPost)
}

func (c *Client) UserUpdateEmail(name string, email string) error {
	return c.userUpdateAttribute(name, "email", email)
}

func (c *Client) UserUpdateDisplayName(name string, displayName string) error {
	return c.userUpdateAttribute(name, "displayname", displayName)
}

func (c *Client) UserUpdatePhone(name string, phone string) error {
	return c.userUpdateAttribute(name, "phone", phone)
}

func (c *Client) UserUpdateAddress(name string, address string) error {
	return c.userUpdateAttribute(name, "address", address)
}

func (c *Client) UserUpdateWebSite(name string, website string) error {
	return c.userUpdateAttribute(name, "website", website)
}

func (c *Client) UserUpdateTwitter(name string, twitter string) error {
	return c.userUpdateAttribute(name, "twitter", twitter)
}

func (c *Client) UserUpdatePassword(name string, password string) error {
	return c.userUpdateAttribute(name, "password", password)
}

func (c *Client) UserUpdateQuota(name string, quota string) error {
	return c.userUpdateAttribute(name, "quota", quota)
}

func (c *Client) UserGroupList(name string) ([]string, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name, "groups")
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	if r.Ocs.Meta.Statuscode != 100 {
		return nil, fmt.Errorf("%d : %s", r.Ocs.Meta.Statuscode, r.Ocs.Meta.Message)
	}
	return r.Ocs.Data.Groups, nil
}

func (c *Client) UserGroupAdd(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(name, "groups", ro, http.MethodPost)
}

func (c *Client) UserGroupRemove(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(name, "groups", ro, http.MethodDelete)
}

func (c *Client) UserGroupPromote(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(name, "subadmins", ro, http.MethodPost)
}

func (c *Client) UserGroupDemote(name string, group string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": group,
		},
	}
	return c.userBaseRequest(name, "subadmins", ro, http.MethodDelete)
}

func (c *Client) UserGroupSubAdminList(name string) ([]string, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name, "subadmins")
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.BaseResponse
	res.JSON(&r)
	if r.Ocs.Meta.Statuscode != 100 {
		return nil, fmt.Errorf("%d : %s", r.Ocs.Meta.Statuscode, r.Ocs.Meta.Message)
	}
	return r.Ocs.Data, nil
}

func (c *Client) userUpdateAttribute(name string, key string, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"key":   key,
			"value": value,
		},
	}
	return c.userBaseRequest(name, "", ro, http.MethodPut)
}

func (c *Client) userBaseRequest(name string, route string, ro *req.RequestOptions, method string) error {
	res, err := c.baseRequest(routes.users, name, route, ro, method)
	if err != nil {
		return err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf("%d : %s", ur.Ocs.Meta.Statuscode, ur.Ocs.Meta.Message)
	}
	return nil
}
