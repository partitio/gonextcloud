# [WIP] Nextcloud Go API Client

A simple Client for Nextcloud's API in Go.

## TODO
- [Auth](#authentication)
  - ~~login~~
  - ~~logout~~
- [Users](#users)
  - ~~search~~
  - ~~list~~
  - ~~get infos~~
  - ~~create~~
  - update
  - ~~delete~~
  - enable
  - disable
  - get groups
  - add to group
  - remove from group
  - get subadmin group
  - promote subadmin
  - demote subadmin
  - send welcome mail
- [Groups](#groups)
  - create
  - delete
  - get members
  - get subadmins
- [Apps](#apps)
  - list
  - get infos
  - enable
  - disable

# Getting started
## Authentication
```go
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
}
````
##Users
List :
```go
func (c *Client) UserList() ([]string, error)
```
Search
```go
func (c *Client) UserSearch(search string) ([]string, error)
```
Get
```go
func (c *Client) User(name string) (*types.User, error)
```
Create
```go
func (c *Client) UserCreate(username string, password string) error
```
Delete
```go
func (c *Client) UserDelete(name string) error 
```
## Groups
TODO

## Apps
TODO