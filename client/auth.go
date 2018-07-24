package client

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/client/types"
)

var unauthorized = fmt.Errorf("login first")

func (c *Client) Login(username string, password string) error {
	c.username = username
	c.password = password
	options := req.RequestOptions{
		Headers: c.headers,
		Auth:    []string{c.username, c.password},
	}
	c.session = req.NewSession(&options)
	// TODO What to do with capabilities ? (other thant connection validation)
	u := c.baseURL.ResolveReference(routes.capabilities)
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return err
	}
	var r types.CapabilitiesResponse
	res.JSON(&r)
	if r.Ocs.Meta.Statuscode != 100 {
		e := types.ErrorFromMeta(r.Ocs.Meta)
		return &e
	}
	c.capabilities = &r.Ocs.Data.Capabilities
	// Check if authentication failed
	if !c.loggedIn() {
		e := types.APIError{Message: "authentication failed"}
		return &e
	}
	return nil
}

func (c *Client) Logout() error {
	c.session.CloseIdleConnections()
	return nil
}

func (c *Client) loggedIn() bool {
	// When authentication failed, capabilities doesn't contains core information
	return c.capabilities.Core.WebdavRoot != ""
}
