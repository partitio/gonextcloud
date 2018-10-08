![Nextcloud](https://upload.wikimedia.org/wikipedia/commons/thumb/6/60/Nextcloud_Logo.svg/2000px-Nextcloud_Logo.svg.png)

[![Go Report Card](https://goreportcard.com/badge/gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud)](https://goreportcard.com/report/gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud)

# gonextcloud

```go
import "gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/client"
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


#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(hostname string) (*Client, error)
```

#### func (*Client) AppDisable

```go
func (c *Client) AppDisable(name string) error
```

#### func (*Client) AppEnable

```go
func (c *Client) AppEnable(name string) error
```

#### func (*Client) AppInfos

```go
func (c *Client) AppInfos(name string) (types.App, error)
```

#### func (*Client) AppList

```go
func (c *Client) AppList() ([]string, error)
```

#### func (*Client) AppListDisabled

```go
func (c *Client) AppListDisabled() ([]string, error)
```

#### func (*Client) AppListEnabled

```go
func (c *Client) AppListEnabled() ([]string, error)
```

#### func (*Client) GroupCreate

```go
func (c *Client) GroupCreate(name string) error
```

#### func (*Client) GroupDelete

```go
func (c *Client) GroupDelete(name string) error
```

#### func (*Client) GroupList

```go
func (c *Client) GroupList() ([]string, error)
```

#### func (*Client) GroupSearch

```go
func (c *Client) GroupSearch(search string) ([]string, error)
```

#### func (*Client) GroupSubAdminList

```go
func (c *Client) GroupSubAdminList(name string) ([]string, error)
```

#### func (*Client) GroupUsers

```go
func (c *Client) GroupUsers(name string) ([]string, error)
```

#### func (*Client) Login

```go
func (c *Client) Login(username string, password string) error
```

#### func (*Client) Logout

```go
func (c *Client) Logout() error
```

#### func (*Client) User

```go
func (c *Client) User(name string) (*types.User, error)
```

#### func (*Client) UserCreate

```go
func (c *Client) UserCreate(username string, password string, user *types.User) error
```

#### func (*Client) UserDelete

```go
func (c *Client) UserDelete(name string) error
```

#### func (*Client) UserDisable

```go
func (c *Client) UserDisable(name string) error
```

#### func (*Client) UserEnable

```go
func (c *Client) UserEnable(name string) error
```

#### func (*Client) UserGroupAdd

```go
func (c *Client) UserGroupAdd(name string, group string) error
```

#### func (*Client) UserGroupDemote

```go
func (c *Client) UserGroupDemote(name string, group string) error
```

#### func (*Client) UserGroupList

```go
func (c *Client) UserGroupList(name string) ([]string, error)
```

#### func (*Client) UserGroupPromote

```go
func (c *Client) UserGroupPromote(name string, group string) error
```

#### func (*Client) UserGroupRemove

```go
func (c *Client) UserGroupRemove(name string, group string) error
```

#### func (*Client) UserGroupSubAdminList

```go
func (c *Client) UserGroupSubAdminList(name string) ([]string, error)
```

#### func (*Client) UserList

```go
func (c *Client) UserList() ([]string, error)
```

#### func (*Client) UserSearch

```go
func (c *Client) UserSearch(search string) ([]string, error)
```

#### func (*Client) UserSendWelcomeEmail

```go
func (c *Client) UserSendWelcomeEmail(name string) error
```

#### func (*Client) UserUpdate

```go
func (c *Client) UserUpdate(user *types.User) error
```

#### func (*Client) UserUpdateAddress

```go
func (c *Client) UserUpdateAddress(name string, address string) error
```

#### func (*Client) UserUpdateDisplayName

```go
func (c *Client) UserUpdateDisplayName(name string, displayName string) error
```

#### func (*Client) UserUpdateEmail

```go
func (c *Client) UserUpdateEmail(name string, email string) error
```

#### func (*Client) UserUpdatePassword

```go
func (c *Client) UserUpdatePassword(name string, password string) error
```

#### func (*Client) UserUpdatePhone

```go
func (c *Client) UserUpdatePhone(name string, phone string) error
```

#### func (*Client) UserUpdateQuota

```go
func (c *Client) UserUpdateQuota(name string, quota int) error
```

#### func (*Client) UserUpdateTwitter

```go
func (c *Client) UserUpdateTwitter(name string, twitter string) error
```

#### func (*Client) UserUpdateWebSite

```go
func (c *Client) UserUpdateWebSite(name string, website string) error
```

#### type Routes

```go
type Routes struct {
}
```
# types

```go
import "gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
```

#### type APIError

```go
type APIError struct {
	Code    int
	Message string
}
```


#### func  ErrorFromMeta

```go
func ErrorFromMeta(meta Meta) *APIError
```

#### func (*APIError) Error

```go
func (e *APIError) Error() string
```

#### type App

```go
type App struct {
	ID            string   `json:"id"`
	Ocsid         string   `json:"ocsid"`
	Name          string   `json:"name"`
	Summary       string   `json:"summary"`
	Description   string   `json:"description"`
	Licence       string   `json:"licence"`
	Author        string   `json:"author"`
	Version       string   `json:"version"`
	Namespace     string   `json:"namespace"`
	Types         []string `json:"types"`
	Documentation struct {
		Admin     string `json:"admin"`
		Developer string `json:"developer"`
		User      string `json:"user"`
	} `json:"documentation"`
	Category   []string `json:"category"`
	Website    string   `json:"website"`
	Bugs       string   `json:"bugs"`
	Repository struct {
		Attributes struct {
			Type string `json:"type"`
		} `json:"@attributes"`
		Value string `json:"@value"`
	} `json:"repository"`
	Screenshot   []interface{} `json:"screenshot"`
	Dependencies struct {
		Owncloud struct {
			Attributes struct {
				MinVersion string `json:"min-version"`
				MaxVersion string `json:"max-version"`
			} `json:"@attributes"`
		} `json:"owncloud"`
		Nextcloud struct {
			Attributes struct {
				MinVersion string `json:"min-version"`
				MaxVersion string `json:"max-version"`
			} `json:"@attributes"`
		} `json:"nextcloud"`
	} `json:"dependencies"`
	Settings struct {
		Admin           []string      `json:"admin"`
		AdminSection    []string      `json:"admin-section"`
		Personal        []interface{} `json:"personal"`
		PersonalSection []interface{} `json:"personal-section"`
	} `json:"settings"`
	Info        []interface{} `json:"info"`
	Remote      []interface{} `json:"remote"`
	Public      []interface{} `json:"public"`
	RepairSteps struct {
		Install       []interface{} `json:"install"`
		PreMigration  []interface{} `json:"pre-migration"`
		PostMigration []interface{} `json:"post-migration"`
		LiveMigration []interface{} `json:"live-migration"`
		Uninstall     []interface{} `json:"uninstall"`
	} `json:"repair-steps"`
	BackgroundJobs     []interface{} `json:"background-jobs"`
	TwoFactorProviders []interface{} `json:"two-factor-providers"`
	Commands           []interface{} `json:"commands"`
	Activity           struct {
		Filters   []interface{} `json:"filters"`
		Settings  []interface{} `json:"settings"`
		Providers []interface{} `json:"providers"`
	} `json:"activity"`
}
```


#### type AppListResponse

```go
type AppListResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Apps []string `json:"apps"`
		} `json:"data"`
	} `json:"ocs"`
}
```


#### type AppResponse

```go
type AppResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data App  `json:"data"`
	} `json:"ocs"`
}
```


#### type BaseResponse

```go
type BaseResponse struct {
	Ocs struct {
		Meta Meta     `json:"meta"`
		Data []string `json:"data"`
	} `json:"ocs"`
}
```


#### type Capabilities

```go
type Capabilities struct {
	Core struct {
		Pollinterval int    `json:"pollinterval"`
		WebdavRoot   string `json:"webdav-root"`
	} `json:"core"`
	Bruteforce struct {
		Delay int `json:"delay"`
	} `json:"bruteforce"`
	Activity struct {
		Apiv2 []string `json:"apiv2"`
	} `json:"activity"`
	Dav struct {
		Chunking string `json:"chunking"`
	} `json:"dav"`
	FilesSharing struct {
		APIEnabled bool `json:"api_enabled"`
		Public     struct {
			Enabled  bool `json:"enabled"`
			Password struct {
				Enforced bool `json:"enforced"`
			} `json:"password"`
			ExpireDate struct {
				Enabled bool `json:"enabled"`
			} `json:"expire_date"`
			SendMail        bool `json:"send_mail"`
			Upload          bool `json:"upload"`
			UploadFilesDrop bool `json:"upload_files_drop"`
		} `json:"public"`
		Resharing bool `json:"resharing"`
		User      struct {
			SendMail   bool `json:"send_mail"`
			ExpireDate struct {
				Enabled bool `json:"enabled"`
			} `json:"expire_date"`
		} `json:"user"`
		GroupSharing bool `json:"group_sharing"`
		Group        struct {
			Enabled    bool `json:"enabled"`
			ExpireDate struct {
				Enabled bool `json:"enabled"`
			} `json:"expire_date"`
		} `json:"group"`
		Federation struct {
			Outgoing   bool `json:"outgoing"`
			Incoming   bool `json:"incoming"`
			ExpireDate struct {
				Enabled bool `json:"enabled"`
			} `json:"expire_date"`
		} `json:"federation"`
		Sharebymail struct {
			Enabled         bool `json:"enabled"`
			UploadFilesDrop struct {
				Enabled bool `json:"enabled"`
			} `json:"upload_files_drop"`
			Password struct {
				Enabled bool `json:"enabled"`
			} `json:"password"`
			ExpireDate struct {
				Enabled bool `json:"enabled"`
			} `json:"expire_date"`
		} `json:"sharebymail"`
	} `json:"files_sharing"`
	Notifications struct {
		OcsEndpoints []string `json:"ocs-endpoints"`
		Push         []string `json:"push"`
	} `json:"notifications"`
	PasswordPolicy struct {
		MinLength                int  `json:"minLength"`
		EnforceNonCommonPassword bool `json:"enforceNonCommonPassword"`
		EnforceNumericCharacters bool `json:"enforceNumericCharacters"`
		EnforceSpecialCharacters bool `json:"enforceSpecialCharacters"`
		EnforceUpperLowerCase    bool `json:"enforceUpperLowerCase"`
	} `json:"password_policy"`
	Theming struct {
		Name              string `json:"name"`
		URL               string `json:"url"`
		Slogan            string `json:"slogan"`
		Color             string `json:"color"`
		ColorText         string `json:"color-text"`
		ColorElement      string `json:"color-element"`
		Logo              string `json:"logo"`
		Background        string `json:"background"`
		BackgroundPlain   bool   `json:"background-plain"`
		BackgroundDefault bool   `json:"background-default"`
	} `json:"theming"`
	Files struct {
		Bigfilechunking  bool     `json:"bigfilechunking"`
		BlacklistedFiles []string `json:"blacklisted_files"`
		Undelete         bool     `json:"undelete"`
		Versioning       bool     `json:"versioning"`
	} `json:"files"`
}
```


#### type CapabilitiesResponse

```go
type CapabilitiesResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Version struct {
				Major   int    `json:"major"`
				Minor   int    `json:"minor"`
				Micro   int    `json:"micro"`
				String  string `json:"string"`
				Edition string `json:"edition"`
			} `json:"version"`
			Capabilities Capabilities `json:"capabilities"`
		} `json:"data"`
	} `json:"ocs"`
}
```


#### type ErrorResponse

```go
type ErrorResponse struct {
	Ocs struct {
		Meta Meta          `json:"meta"`
		Data []interface{} `json:"data"`
	} `json:"ocs"`
}
```


#### type GroupListResponse

```go
type GroupListResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Groups []string `json:"groups"`
		} `json:"data"`
	} `json:"ocs"`
}
```


#### type Meta

```go
type Meta struct {
	Status       string `json:"status"`
	Statuscode   int    `json:"statuscode"`
	Message      string `json:"message"`
	Totalitems   string `json:"totalitems"`
	Itemsperpage string `json:"itemsperpage"`
}
```


#### type UpdateError

```go
type UpdateError struct {
	Field string
	Error error
}
```


#### type User

```go
type User struct {
	Enabled bool   `json:"enabled"`
	ID      string `json:"id"`
	Quota   struct {
		Free     int64   `json:"free"`
		Used     int     `json:"used"`
		Total    int64   `json:"total"`
		Relative float64 `json:"relative"`
		Quota    int     `json:"quota"`
	} `json:"quota"`
	Email       string   `json:"email"`
	Displayname string   `json:"displayname"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	Website     string   `json:"website"`
	Twitter     string   `json:"twitter"`
	Groups      []string `json:"groups"`
	Language    string   `json:"language"`
}
```


#### type UserListResponse

```go
type UserListResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Users []string `json:"users"`
		} `json:"data"`
	} `json:"ocs"`
}
```


#### type UserResponse

```go
type UserResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data User `json:"data"`
	} `json:"ocs"`
}
```


#### type UserUpdateError

```go
type UserUpdateError struct {
	Errors map[string]error
}
```


#### func  NewUpdateError

```go
func NewUpdateError(errors chan UpdateError) *UserUpdateError
```

#### func (*UserUpdateError) Error

```go
func (e *UserUpdateError) Error() string
```
