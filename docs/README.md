![Nextcloud](https://upload.wikimedia.org/wikipedia/commons/thumb/6/60/Nextcloud_Logo.svg/2000px-Nextcloud_Logo.svg.png)

# gonextcloud

```go
import "github.com/partitio/gonextcloud/client"
```


## Usage

```go
package main

import (
	"fmt"
	"github.com/partitio/gonextcloud"
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
