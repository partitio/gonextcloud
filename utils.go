package gonextcloud

import (
	"encoding/json"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func (c *Client) baseRequest(route *url.URL, name string, subroute string, ro *req.RequestOptions, method string) (*req.Response, error) {
	if !c.loggedIn() {
		return nil, errUnauthorized
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
	// As we cannot read the ReaderCloser twice, we use the string content
	js := res.String()
	var r types.BaseResponse
	json.Unmarshal([]byte(js), &r)
	if r.Ocs.Meta.Statuscode != 100 {
		err := types.ErrorFromMeta(r.Ocs.Meta)
		return nil, err
	}
	return res, nil
}

func reformatJSON(json string) string {
	// Nextcloud encode boolean as string
	json = strings.Replace(json, "\"true\"", "true", -1)
	json = strings.Replace(json, "\"false\"", "false", -1)
	// Nextcloud encode quota as an empty array for never connected users
	json = strings.Replace(json, "\"quota\":[],", "", -1)
	return json
}
