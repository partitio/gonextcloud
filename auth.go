package gonextcloud

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
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
	// No need to check for Ocs.Meta.StatusCode as capabilities are always returned
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
	c.session.HTTPClient.Jar = nil
	// Clear capabilities as it is used to check for valid authentication
	c.capabilities = nil
	return nil
}

func (c *Client) loggedIn() bool {
	// When authentication failed, capabilities doesn't contains core information
	if c.capabilities == nil {
		return false
	}
	return c.capabilities.Core.WebdavRoot != ""
}
