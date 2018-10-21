package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
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

	Apps          *Apps
	AppsConfig    *AppsConfig
	GroupFolders  *GroupFolders
	Notifications *Notifications
	Shares        *Shares
	Users         *Users
	Groups        *Groups
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
	c.Apps = &Apps{c}
	c.AppsConfig = &AppsConfig{c}
	c.GroupFolders = &GroupFolders{c}
	c.Notifications = &Notifications{c}
	c.Shares = &Shares{c}
	c.Users = &Users{c}
	c.Groups = &Groups{c}
	return c, nil
}
