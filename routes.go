package gonextcloud

import "net/url"

// Routes references the available routes
type Routes struct {
	capabilities *url.URL
	users        *url.URL
	groups       *url.URL
	apps         *url.URL
	monitor      *url.URL
}

const badRequest = 998

var (
	apiPath = &url.URL{Path: "/ocs/v1.php"}
	routes  = Routes{
		capabilities: &url.URL{Path: apiPath.Path + "/cloud/capabilities"},
		users:        &url.URL{Path: apiPath.Path + "/cloud/users"},
		groups:       &url.URL{Path: apiPath.Path + "/cloud/groups"},
		apps:         &url.URL{Path: apiPath.Path + "/cloud/apps"},
		monitor:      &url.URL{Path: apiPath.Path + "/apps/serverinfo/api/v1/info"},
	}
)
