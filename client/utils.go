package client

import (
	req "github.com/levigross/grequests"
	"net/http"
	"net/url"
	"path"
)

func (c *Client) baseRequest(route *url.URL, name string, subroute string, ro *req.RequestOptions, method string) (*req.Response, error) {
	if !c.loggedIn() {
		return nil, unauthorized
	}
	u := c.baseURL.ResolveReference(route)
	if name != "" {
		u.Path = path.Join(u.Path, name)
	}
	if subroute != "" {
		u.Path = path.Join(u.Path, subroute)
	}
	var (
		res *req.Response
		err error
	)
	if method == http.MethodGet {
		res, err = c.session.Get(u.String(), ro)
	} else if method == http.MethodPost {
		res, err = c.session.Post(u.String(), ro)
	} else if method == http.MethodPut {
		res, err = c.session.Put(u.String(), ro)
	} else if method == http.MethodDelete {
		res, err = c.session.Delete(u.String(), ro)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
