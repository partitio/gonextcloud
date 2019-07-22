package gonextcloud

import (
	"net/http"

	req "github.com/levigross/grequests"
)

//apps contains all apps available actions
type apps struct {
	c *client
}

//List return the list of the Nextcloud apps
func (a *apps) List() ([]string, error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, nil)
	if err != nil {
		return nil, err
	}
	var r AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//ListEnabled lists the enabled apps
func (a *apps) ListEnabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "enabled"},
	}
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, ro)
	if err != nil {
		return nil, err
	}
	var r AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//ListDisabled lists the disabled apps
func (a *apps) ListDisabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "disabled"},
	}
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, ro)
	if err != nil {
		return nil, err
	}
	var r AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//Infos return the app's details
func (a *apps) Infos(name string) (App, error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.apps, nil, name)
	if err != nil {
		return App{}, err
	}
	var r AppResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Enable enables an app
func (a *apps) Enable(name string) error {
	_, err := a.c.baseRequest(http.MethodPost, routes.apps, nil, name)
	return err
}

//Disable disables an app
func (a *apps) Disable(name string) error {
	_, err := a.c.baseRequest(http.MethodDelete, routes.apps, nil, name)
	return err
}
