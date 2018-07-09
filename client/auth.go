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
	r, err := c.session.Get(u.String(), nil)
	if err != nil {
		return err
	}
	var cs types.CapabilitiesResponse
	r.JSON(&cs)
	if cs.Ocs.Meta.Statuscode != 100 {
		return fmt.Errorf("%d : %s", cs.Ocs.Meta.Statuscode, cs.Ocs.Meta.Message)
	}
	c.capabilities = &cs.Ocs.Data.Capabilities
	return nil
}

func (c *Client) Logout() error {
	c.session.CloseIdleConnections()
	return nil
}

func (c *Client) loggedIn() bool {
	return c.capabilities != nil
}
