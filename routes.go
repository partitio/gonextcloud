package gonextcloud

import "net/url"

// Route references the available routes
type Routes struct {
	capabilities *url.URL
	users        *url.URL
	groups       *url.URL
	apps         *url.URL
}

const badRequest = 998

var (
	apiPath = &url.URL{Path: "/ocs/v1.php/cloud"}
	routes  = Routes{
		capabilities: &url.URL{Path: apiPath.Path + "/capabilities"},
		users:        &url.URL{Path: apiPath.Path + "/users"},
		groups:       &url.URL{Path: apiPath.Path + "/groups"},
		apps:         &url.URL{Path: apiPath.Path + "/apps"},
	}
)
