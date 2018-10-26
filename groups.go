package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
)

//GroupsI available methods
type GroupsI interface {
	List() ([]string, error)
	ListDetails() ([]types.Group, error)
	Users(name string) ([]string, error)
	Search(search string) ([]string, error)
	Create(name string) error
	Delete(name string) error
	SubAdminList(name string) ([]string, error)
}

//Groups contains all Groups available actions
type Groups struct {
	c *Client
}

//List lists the Nextcloud groups
func (g *Groups) List() ([]string, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil)
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//ListDetails lists the Nextcloud groups
func (g *Groups) ListDetails() ([]types.Group, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil, "details")
	if err != nil {
		return nil, err
	}
	var r types.GroupListDetailsResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//Users list the group's users
func (g *Groups) Users(name string) ([]string, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil, name)
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

//Search return the list of groups matching the search string
func (g *Groups) Search(search string) ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"search": search},
	}
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, ro)
	if err != nil {
		return nil, err
	}
	var r types.GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//Create creates a group
func (g *Groups) Create(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": name,
		},
	}
	return g.baseRequest(http.MethodPost, ro)
}

//Delete deletes the group
func (g *Groups) Delete(name string) error {
	return g.baseRequest(http.MethodDelete, nil, name)
}

//SubAdminList lists the group's subadmins
func (g *Groups) SubAdminList(name string) ([]string, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil, name, "subadmins")
	if err != nil {
		return nil, err
	}
	var r types.UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

func (g *Groups) baseRequest(method string, ro *req.RequestOptions, subRoute ...string) error {
	_, err := g.c.baseRequest(method, routes.groups, ro, subRoute...)
	return err
}
