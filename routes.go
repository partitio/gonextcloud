package gonextcloud

import "net/url"

// apiRoutes references the available routes
type apiRoutes struct {
	capabilities       *url.URL
	users              *url.URL
	groups             *url.URL
	apps               *url.URL
	monitor            *url.URL
	shares             *url.URL
	groupfolders       *url.URL
	appsConfig         *url.URL
	notifications      *url.URL
	adminNotifications *url.URL
}

const badRequest = 998

var (
	apiPath = &url.URL{Path: "/ocs/v2.php"}
	routes  = apiRoutes{
		capabilities:       &url.URL{Path: apiPath.Path + "/cloud/capabilities"},
		users:              &url.URL{Path: apiPath.Path + "/cloud/users"},
		groups:             &url.URL{Path: apiPath.Path + "/cloud/groups"},
		apps:               &url.URL{Path: apiPath.Path + "/cloud/apps"},
		monitor:            &url.URL{Path: apiPath.Path + "/apps/serverinfo/api/v1/info"},
		shares:             &url.URL{Path: apiPath.Path + "/apps/files_sharing/api/v1/shares"},
		groupfolders:       &url.URL{Path: "/apps/groupfolders/folders"},
		appsConfig:         &url.URL{Path: apiPath.Path + "/apps/provisioning_api/api/v1/config/apps"},
		notifications:      &url.URL{Path: apiPath.Path + "/apps/notifications/api/v2/notifications"},
		adminNotifications: &url.URL{Path: apiPath.Path + "/apps/admin_notifications/api/v2/notifications"},
	}
)
