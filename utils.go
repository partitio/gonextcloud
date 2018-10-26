package gonextcloud

import (
	"encoding/json"
	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func (c *Client) baseRequest(method string, route *url.URL, ro *req.RequestOptions, subRoutes ...string) (*req.Response, error) {
	if !c.loggedIn() {
		return nil, errUnauthorized
	}
	u := c.baseURL.ResolveReference(route)

	for _, sr := range subRoutes {
		if sr != "" {
			u.Path = path.Join(u.Path, sr)
		}
	}
	var (
		res *req.Response
		err error
	)
	switch method {
	case http.MethodGet:
		res, err = c.session.Get(u.String(), ro)
	case http.MethodPost:
		res, err = c.session.Post(u.String(), ro)
	case http.MethodPut:
		res, err = c.session.Put(u.String(), ro)
	case http.MethodDelete:
		res, err = c.session.Delete(u.String(), ro)
	}
	if err != nil {
		return nil, err
	}
	// As we cannot read the ReaderCloser twice, we use the string content
	js := res.String()
	var r types.BaseResponse
	json.Unmarshal([]byte(js), &r)
	if r.Ocs.Meta.Statuscode == 200 || r.Ocs.Meta.Statuscode == 100 {
		return res, nil
	}
	err = types.ErrorFromMeta(r.Ocs.Meta)
	return nil, err
}

func reformatJSON(json string) string {
	// Nextcloud encode boolean as string
	json = strings.Replace(json, "\"true\"", "true", -1)
	json = strings.Replace(json, "\"false\"", "false", -1)
	// Nextcloud encode quota as an empty array for never connected users
	json = strings.Replace(json, "\"quota\":[],", "", -1)
	return json
}
