/*
Package client is a Go client for the Nextcloud Provisioning API.

For more information about the Provisioning API, see the documentation:
https://docs.nextcloud.com/server/13/admin_manual/configuration_user/user_provisioning_api.html

Usage

You use the library by creating a client object and calling methods on it.

For example, to list all the Nextcloud's instance users:

	package main

	import (
		"fmt"
		"github.com/partitio/gonextcloud/client"
	)

	func main() {
		url := "https://www.mynextcloud.com"
		username := "admin"
		password := "password"
		c, err := client.NewClient(url)
		if err != nil {
			panic(err)
		}
		if err := c.Login(username, password); err != nil {
			panic(err)
		}
		defer c.Logout()

		users, err := c.UserList()
		if err != nil {
			panic(err)
		}
		fmt.Println("Users :", users)
	}
*/

package gonextcloud

import (
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
	"net/url"
)

type Client struct {
	baseURL      *url.URL
	username     string
	password     string
	session      *req.Session
	headers      map[string]string
	capabilities *types.Capabilities
}

func NewClient(hostname string) (*Client, error) {
	baseURL, err := url.ParseRequestURI(hostname)
	if err != nil {
		baseURL, err = url.ParseRequestURI("https://" + hostname)
		if err != nil {
			return nil, err
		}
	}

	c := Client{
		baseURL: baseURL,
		headers: map[string]string{
			"OCS-APIREQUEST": "true",
			"Accept":         "application/json",
		},
	}
	return &c, nil
}
