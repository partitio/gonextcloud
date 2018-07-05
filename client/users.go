package client

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/client/types"
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
		return nil, fmt.Errorf(ur.Ocs.Meta.Message)
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
		return nil, fmt.Errorf(r.Ocs.Meta.Message)
	}
	return r.Ocs.Data.Users, nil
}

func (c *Client) UserCreate(username string, password string) error {
	if !c.loggedIn() {
		return unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	ro := &req.RequestOptions{
		Data: map[string]string{
			"userid":   username,
			"password": password,
		},
	}
	res, err := c.session.Post(u.String(), ro)
	if err != nil {
		return err
	}
	var r types.UserResponse
	res.JSON(&r)
	if r.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(r.Ocs.Meta.Message)
	}
	return nil
}

func (c *Client) UserDelete(name string) error {
	if !c.loggedIn() {
		return unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name)
	res, err := c.session.Delete(u.String(), nil)
	if err != nil {
		return err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(ur.Ocs.Meta.Message)
	}
	return nil
}

func (c *Client) UserEnable(name string) error {
	if !c.loggedIn() {
		return unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name, "enable")
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	res, err := c.session.Put(u.String(), ro)
	if err != nil {
		return err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(ur.Ocs.Meta.Message)
	}
	return nil
}

func (c *Client) UserDisable(name string) error {
	if !c.loggedIn() {
		return unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name, "disable")
	ro := &req.RequestOptions{
		Data: map[string]string{},
	}
	res, err := c.session.Put(u.String(), ro)
	if err != nil {
		return err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(ur.Ocs.Meta.Message)
	}
	return nil
}

func (c *Client) UserSendWelcomeEmail(name string) error {
	if !c.loggedIn() {
		return unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name, "welcome")
	res, err := c.session.Post(u.String(), nil)
	if err != nil {
		return err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(ur.Ocs.Meta.Message)
	}
	return nil
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

func (c *Client) userUpdateAttribute(name string, key string, value string) error {
	if !c.loggedIn() {
		return unauthorized
	}
	u := c.baseURL.ResolveReference(routes.users)
	u.Path = path.Join(u.Path, name)
	ro := &req.RequestOptions{
		Data: map[string]string{
			"key":   key,
			"value": value,
		},
	}
	res, err := c.session.Put(u.String(), ro)
	if err != nil {
		return err
	}
	var ur types.UserResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(ur.Ocs.Meta.Message)
	}
	return nil
}
