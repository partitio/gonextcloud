![Nextcloud](https://upload.wikimedia.org/wikipedia/commons/thumb/6/60/Nextcloud_Logo.svg/640px-Nextcloud_Logo.svg.png)

[![pipeline status](http://gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/badges/master/pipeline.svg)](http://gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/commits/master)
[![coverage report](http://gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/badges/master/coverage.svg)](http://gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud)](https://goreportcard.com/report/gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud)
[![GoDoc](https://godoc.org/gitlab.com/gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud?status.svg)](https://godoc.org/gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud)
# GoNextcloud

A simple Client for Nextcloud's Provisioning API in Go.


```go
import "gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud"
```


## Usage

```go
package main

import (
	"fmt"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud"
)

func main() {
	url := "https://www.mynextcloud.com"
	username := "admin"
	password := "password"
	c, err := gonextcloud.NewClient(url)
	if err != nil {
		panic(err)
	}
	if err := c.Login(username, password); err != nil {
		panic(err)
	}
	defer c.Logout()
}
```
## Run tests
Configure the tests for your instance by editing [example.config.yml](example.config.yml) and renaming it config.yml

then run the tests :
```bash
$ go test -v .
```
