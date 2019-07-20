package gonextcloud

import (
	"net/url"

	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/adphi/gowebdav"

	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
)

// Client is the API client that performs all operations against a Nextcloud server.
type Client struct {
	baseURL      *url.URL
	username     string
	password     string
	session      *req.Session
	headers      map[string]string
	capabilities *types.Capabilities
	version      *types.Version

	apps          *Apps
	appsConfig    *AppsConfig
	groupFolders  *GroupFolders
	notifications *Notifications
	shares        *Shares
	users         *Users
	groups        *Groups
	webdav        *webDav
}

// NewClient create a new Client from the Nextcloud Instance URL
func NewClient(hostname string) (*Client, error) {
	baseURL, err := url.ParseRequestURI(hostname)
	if err != nil {
		baseURL, err = url.ParseRequestURI("https://" + hostname)
		if err != nil {
			return nil, err
		}
	}

	c := &Client{
		baseURL: baseURL,
		headers: map[string]string{
			"OCS-APIREQUEST": "true",
			"Accept":         "application/json",
		},
	}

	c.apps = &Apps{c}
	c.appsConfig = &AppsConfig{c}
	c.groupFolders = &GroupFolders{c}
	c.notifications = &Notifications{c}
	c.shares = &Shares{c}
	c.users = &Users{c}
	c.groups = &Groups{c}
	// Create empty webdav client
	// It will be replaced after login
	c.webdav = &webDav{Client: &gowebdav.Client{}}
	return c, nil
}

// Apps return the Apps client Interface
func (c *Client) Apps() types.Apps {
	return c.apps
}

// AppsConfig return the AppsConfig client Interface
func (c *Client) AppsConfig() types.AppsConfig {
	return c.appsConfig
}

// GroupFolders return the GroupFolders client Interface
func (c *Client) GroupFolders() types.GroupFolders {
	return c.groupFolders
}

// Notifications return the Notifications client Interface
func (c *Client) Notifications() types.Notifications {
	return c.notifications
}

// Shares return the Shares client Interface
func (c *Client) Shares() types.Shares {
	return c.shares
}

// Users return the Users client Interface
func (c *Client) Users() types.Users {
	return c.users
}

// Groups return the Groups client Interface
func (c *Client) Groups() types.Groups {
	return c.groups
}

// WebDav return the WebDav client Interface
func (c *Client) WebDav() types.WebDav {
	return c.webdav
}
