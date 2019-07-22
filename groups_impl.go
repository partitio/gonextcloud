package gonextcloud

import (
	"net/http"

	req "github.com/levigross/grequests"
)

//groups contains all groups available actions
type groups struct {
	c *client
}

//List lists the Nextcloud groups
func (g *groups) List() ([]string, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil)
	if err != nil {
		return nil, err
	}
	var r GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//ListDetails lists the Nextcloud groups
func (g *groups) ListDetails(search string) ([]Group, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{
			"search": search,
		},
	}
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, ro, "details")
	if err != nil {
		return nil, err
	}
	var r GroupListDetailsResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//users list the group's users
func (g *groups) Users(name string) ([]string, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil, name)
	if err != nil {
		return nil, err
	}
	var r UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

//Search return the list of groups matching the search string
func (g *groups) Search(search string) ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"search": search},
	}
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, ro)
	if err != nil {
		return nil, err
	}
	var r GroupListResponse
	res.JSON(&r)
	return r.Ocs.Data.Groups, nil
}

//Create creates a group
func (g *groups) Create(name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"groupid": name,
		},
	}
	return g.baseRequest(http.MethodPost, ro)
}

//Delete deletes the group
func (g *groups) Delete(name string) error {
	return g.baseRequest(http.MethodDelete, nil, name)
}

//SubAdminList lists the group's subadmins
func (g *groups) SubAdminList(name string) ([]string, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groups, nil, name, "subadmins")
	if err != nil {
		return nil, err
	}
	var r UserListResponse
	res.JSON(&r)
	return r.Ocs.Data.Users, nil
}

func (g *groups) baseRequest(method string, ro *req.RequestOptions, subRoute ...string) error {
	_, err := g.c.baseRequest(method, routes.groups, ro, subRoute...)
	return err
}
