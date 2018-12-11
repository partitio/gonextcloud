package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/url"
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
	return c, nil
}

func (c *Client) Apps() types.Apps {
	return c.apps
}

func (c *Client) AppsConfig() types.AppsConfig {
	return c.appsConfig
}

func (c *Client) GroupFolders() types.GroupFolders {
	return c.groupFolders
}

func (c *Client) Notifications() types.Notifications {
	return c.notifications
}

func (c *Client) Shares() types.Shares {
	return c.shares
}

func (c *Client) Users() types.Users {
	return c.users
}

func (c *Client) Groups() types.Groups {
	return c.groups
}
