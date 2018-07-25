package gonextcloud

import (
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
	"net/http"
)

func (c *Client) AppList() ([]string, error) {
	res, err := c.baseRequest(routes.apps, "", "", nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

func (c *Client) AppListEnabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "enabled"},
	}
	res, err := c.baseRequest(routes.apps, "", "", ro, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

func (c *Client) AppListDisabled() ([]string, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{"filter": "disabled"},
	}
	res, err := c.baseRequest(routes.apps, "", "", ro, http.MethodGet)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
}

func (c *Client) AppInfos(name string) (types.App, error) {
	res, err := c.baseRequest(routes.apps, name, "", nil, http.MethodGet)
	if err != nil {
		return types.App{}, err
	}
	var r types.AppResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

func (c *Client) AppEnable(name string) error {
	res, err := c.baseRequest(routes.apps, name, "", nil, http.MethodPut)
	if err != nil {
		return err
	}
	var r types.BaseResponse
	res.JSON(&r)
	return nil
}

func (c *Client) AppDisable(name string) error {
	res, err := c.baseRequest(routes.apps, name, "", nil, http.MethodDelete)
	if err != nil {
		return err
	}
	var r types.BaseResponse
	res.JSON(&r)
	return nil
}

func (c *Client) appsBaseRequest(name string, route string, ro *req.RequestOptions, method string) error {
	_, err := c.baseRequest(routes.apps, name, route, ro, method)
	return err
}
