package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
)

//AppList return the list of the Nextcloud Apps
func (c *Client) AppList() ([]string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.apps, nil)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//AppListEnabled lists the enabled apps
func (c *Client) AppListEnabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "enabled"},
	}
	res, err := c.baseRequest(http.MethodGet, routes.apps, ro)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//AppListDisabled lists the disabled apps
func (c *Client) AppListDisabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "disabled"},
	}
	res, err := c.baseRequest(http.MethodGet, routes.apps, ro)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

//AppInfos return the app's details
func (c *Client) AppInfos(name string) (types.App, error) {
	res, err := c.baseRequest(http.MethodGet, routes.apps, nil, name)
	if err != nil {
		return types.App{}, err
	}
	var r types.AppResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//AppEnable enables an app
func (c *Client) AppEnable(name string) error {
	_, err := c.baseRequest(http.MethodPost, routes.apps, nil, name)
	return err
}

//AppDisable disables an app
func (c *Client) AppDisable(name string) error {
	_, err := c.baseRequest(http.MethodDelete, routes.apps, nil, name)
	return err
}
