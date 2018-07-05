package client

import "net/url"

type Routes struct {
	capabilities *url.URL
	users        *url.URL
	groups       *url.URL
}

var (
	apiPath = &url.URL{Path: "/ocs/v1.php/cloud"}
	routes  = Routes{
		capabilities: &url.URL{Path: apiPath.Path + "/capabilities"},
		users:        &url.URL{Path: apiPath.Path + "/users"},
		groups:       &url.URL{Path: apiPath.Path + "/groups"},
	}
	badRequest = 998
)
