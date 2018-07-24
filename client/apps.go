package client

import (
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/client/types"
)

func (c *Client) AppList() ([]string, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(routes.apps)
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var r types.AppListResponse
	res.JSON(&r)
	return r.Ocs.Data.Apps, nil
	return nil, nil
}

func (c *Client) appsBaseRequest(name string, route string, ro *req.RequestOptions, method string) error {
	res, err := c.baseRequest(routes.apps, name, route, ro, method)
	if err != nil {
		return err
	}
	var r types.UserResponse
	res.JSON(&r)
	if r.Ocs.Meta.Statuscode != 100 {
		e := types.ErrorFromMeta(r.Ocs.Meta)
		return &e
	}
	return nil
}
