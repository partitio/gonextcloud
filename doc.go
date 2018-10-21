/*
Package gonextcloud is a Go client for the Nextcloud Provisioning API.

For more information about the Provisioning API, see the documentation:
https://docs.nextcloud.com/server/13/admin_manual/configuration_user/user_provisioning_api.html

Usage

You use the library by creating a client object and calling methods on it.

For example, to list all the Nextcloud's instance users:

	package main

	import (
		"fmt"
		"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/client"
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

		users, err := c.List()
		if err != nil {
			panic(err)
		}
		fmt.Println("Users :", users)
	}
*/
package gonextcloud
