package client

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/client/types"
	"net/http"
)

func (c *Client) GroupList() ([]string, error) {
	res, err := c.baseRequest(routes.groups, "", "", nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var gr types.GroupListResponse
	res.JSON(&gr)
	if gr.Ocs.Meta.Statuscode != 100 {
		return nil, fmt.Errorf("%d : %s", gr.Ocs.Meta.Statuscode, gr.Ocs.Meta.Message)
	}
	return gr.Ocs.Data.Groups, nil
}

func (c *Client) GroupUsers(name string) ([]string, error) {
	res, err := c.baseRequest(routes.groups, name, "", nil, http.MethodGet)
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

func (c *Client) GroupSearch(search string) ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"search": search},
	}
	res, err := c.baseRequest(routes.groups, "", "", ro, http.MethodGet)
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

func (c *Client) GroupCreate(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": name,
		},
	}
	if err := c.groupBaseRequest("", "", ro, http.MethodPost); err != nil {
		return err
	}
	return nil
}

func (c *Client) GroupDelete(name string) error {
	if err := c.groupBaseRequest(name, "", nil, http.MethodDelete); err != nil {
		return err
	}
	return nil
}

func (c *Client) GroupSubAdminList(name string) ([]string, error) {
	res, err := c.baseRequest(routes.groups, name, "subadmins", nil, http.MethodGet)
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

func (c *Client) groupBaseRequest(name string, route string, ro *req.RequestOptions, method string) error {
	res, err := c.baseRequest(routes.groups, name, route, ro, method)
	if err != nil {
		return err
	}
	var ur types.GroupListResponse
	res.JSON(&ur)
	if ur.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf("%d : %s", ur.Ocs.Meta.Statuscode, ur.Ocs.Meta.Message)
	}
	return nil
}
