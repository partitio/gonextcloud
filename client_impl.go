package gonextcloud

import (
	"net/url"

	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/adphi/gowebdav"
)

// client is the API client that performs all operations against a Nextcloud server.
type client struct {
	baseURL      *url.URL
	username     string
	password     string
	session      *req.Session
	headers      map[string]string
	capabilities *Capabilities
	version      *Version

	apps          *apps
	appsConfig    *appsConfig
	groupFolders  *groupFolders
	notifications *notifications
	shares        *shares
	users         *users
	groups        *groups
	webdav        *webDav
	passwords      *passwords
}

// newClient create a new client from the Nextcloud Instance URL
func newClient(hostname string) (*client, error) {
	baseURL, err := url.ParseRequestURI(hostname)
	if err != nil {
		baseURL, err = url.ParseRequestURI("https://" + hostname)
		if err != nil {
			return nil, err
		}
	}

	c := &client{
		baseURL: baseURL,
		headers: map[string]string{
			"OCS-APIREQUEST": "true",
			"Accept":         "application/json",
		},
	}

	c.apps = &apps{c}
	c.appsConfig = &appsConfig{c}
	c.groupFolders = &groupFolders{c}
	c.notifications = &notifications{c}
	c.shares = &shares{c}
	c.users = &users{c}
	c.groups = &groups{c}
	c.passwords = &passwords{c}
	// Create empty webdav client
	// It will be replaced after login
	c.webdav = &webDav{Client: &gowebdav.Client{}}
	return c, nil
}

// apps return the apps client Interface
func (c *client) Apps() Apps {
	return c.apps
}

// appsConfig return the appsConfig client Interface
func (c *client) AppsConfig() AppsConfig {
	return c.appsConfig
}

// groupFolders return the groupFolders client Interface
func (c *client) GroupFolders() GroupFolders {
	return c.groupFolders
}

// notifications return the notifications client Interface
func (c *client) Notifications() Notifications {
	return c.notifications
}

// shares return the shares client Interface
func (c *client) Shares() Shares {
	return c.shares
}

// users return the users client Interface
func (c *client) Users() Users {
	return c.users
}

// groups return the groups client Interface
func (c *client) Groups() Groups {
	return c.groups
}

// WebDav return the WebDav client Interface
func (c *client) WebDav() WebDav {
	return c.webdav
}

// Password return the Password client Interface
func (c *client) Passwords() Passwords {
	return c.passwords
}
