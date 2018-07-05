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
	fmt.Println(res.String())
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
	fmt.Println(res.String())
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf(ur.Ocs.Meta.Message)
	}
	return nil
}
