package gonextcloud

import (
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
	"net/http"
)

//GroupList lists the Nextcloud groups
func (c *Client) GroupList() ([]string, error) {
	res, err := c.baseRequest(routes.groups, "", "", nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//GroupUsers list the group's users
func (c *Client) GroupUsers(name string) ([]string, error) {
	res, err := c.baseRequest(routes.groups, name, "", nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

//GroupSearch return the list of groups matching the search string
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
	return r.Ocs.Data.Groups, nil
}

//GroupCreate creates a group
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

//GroupDelete deletes the group
func (c *Client) GroupDelete(name string) error {
	if err := c.groupBaseRequest(name, "", nil, http.MethodDelete); err != nil {
		return err
	}
	return nil
}

//GroupSubAdminList lists the group's subadmins
func (c *Client) GroupSubAdminList(name string) ([]string, error) {
	res, err := c.baseRequest(routes.groups, name, "subadmins", nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

func (c *Client) groupBaseRequest(name string, route string, ro *req.RequestOptions, method string) error {
	_, err := c.baseRequest(routes.groups, name, route, ro, method)
	return err
}
