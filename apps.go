package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
)

//AppsI available methods
type AppsI interface {
	List() ([]string, error)
	ListEnabled() ([]string, error)
	ListDisabled() ([]string, error)
	Infos(name string) (types.App, error)
	Enable(name string) error
	Disable(name string) error
}

//Apps contains all Apps available actions
type Apps struct {
	c *Client
}

//List return the list of the Nextcloud Apps
func (a *Apps) List() ([]string, error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, nil)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//ListEnabled lists the enabled apps
func (a *Apps) ListEnabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "enabled"},
	}
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, ro)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//ListDisabled lists the disabled apps
func (a *Apps) ListDisabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "disabled"},
	}
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, ro)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//Infos return the app's details
func (a *Apps) Infos(name string) (types.App, error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, nil, name)
	if err != nil {
		return types.App{}, err
	}
	var r types.AppResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Enable enables an app
func (a *Apps) Enable(name string) error {
	_, err := a.c.baseRequest(http.MethodPost, routes.apps, nil, name)
	return err
}

//Disable disables an app
func (a *Apps) Disable(name string) error {
	_, err := a.c.baseRequest(http.MethodDelete, routes.apps, nil, name)
	return err
}
