package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
)

//GroupList lists the Nextcloud groups
func (c *Client) GroupList() ([]string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.groups, nil)
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//GroupUsers list the group's users
func (c *Client) GroupUsers(name string) ([]string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.groups, nil, name)
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
	res, err := c.baseRequest(http.MethodGet, routes.groups, ro)
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
	return c.groupBaseRequest(http.MethodPost, ro)
}

//GroupDelete deletes the group
func (c *Client) GroupDelete(name string) error {
	return c.groupBaseRequest(http.MethodDelete, nil, name)
}

//GroupSubAdminList lists the group's subadmins
func (c *Client) GroupSubAdminList(name string) ([]string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.groups, nil, name, "subadmins")
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

func (c *Client) groupBaseRequest(method string, ro *req.RequestOptions, subRoute ...string) error {
	_, err := c.baseRequest(method, routes.groups, ro, subRoute...)
	return err
}
