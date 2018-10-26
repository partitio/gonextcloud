# gonextcloud

```go
import "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
```

Package gonextcloud is a Go client for the Nextcloud Provisioning API.

For more information about the Provisioning API, see the documentation:
https://docs.nextcloud.com/server/13/admin_manual/configuration_user/user_provisioning_api.html


### Usage

You use the library by creating a client object and calling methods on it.

For example, to list all the Nextcloud's instance users:
```go
package main

import (
    "fmt"
    "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/client"
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

    users, err := c.Users.List()
    if err != nil {
        panic(err)
    }
    fmt.Println("Users :", users)
}
```
#### type Client

```go
type Client struct {
	Apps          *Apps
	AppsConfig    *AppsConfig
	GroupFolders  *GroupFolders
	Notifications *Notifications
	Shares        *Shares
	Users         *Users
	Groups        *Groups
}
```

Client is the API client that performs all operations against a Nextcloud
server.

#### func  NewClient

```go
func NewClient(hostname string) (*Client, error)
```
NewClient create a new Client from the Nextcloud Instance URL

#### func (*Client) Login

```go
func (c *Client) Login(username string, password string) error
```
Login perform login and create a session with the Nextcloud API.

#### func (*Client) Logout

```go
func (c *Client) Logout() error
```
Logout logs out from the Nextcloud API, close the session and delete session's
cookie

#### type Apps

```go
type Apps struct {
}
```

Apps contains all Apps available actions

#### func (*Apps) Disable

```go
func (a *Apps) Disable(name string) error
```
Disable disables an app

#### func (*Apps) Enable

```go
func (a *Apps) Enable(name string) error
```
Enable enables an app

#### func (*Apps) Infos

```go
func (a *Apps) Infos(name string) (types.App, error)
```
Infos return the app's details

#### func (*Apps) List

```go
func (a *Apps) List() ([]string, error)
```
List return the list of the Nextcloud Apps

#### func (*Apps) ListDisabled

```go
func (a *Apps) ListDisabled() ([]string, error)
```
ListDisabled lists the disabled apps

#### func (*Apps) ListEnabled

```go
func (a *Apps) ListEnabled() ([]string, error)
```
ListEnabled lists the enabled apps

#### type AppsConfig

```go
type AppsConfig struct {
}
```

AppsConfig contains all Apps Configuration available actions

#### func (*AppsConfig) DeleteValue

```go
func (a *AppsConfig) DeleteValue(id, key, value string) error
```
DeleteValue delete the config value and (!! be careful !!) the key

#### func (*AppsConfig) Details

```go
func (a *AppsConfig) Details(appID string) (map[string]string, error)
```
Details returns all the config's key, values pair of the app

#### func (*AppsConfig) Get

```go
func (a *AppsConfig) Get() (map[string]map[string]string, error)
```
Get returns all apps AppConfigDetails

#### func (*AppsConfig) Keys

```go
func (a *AppsConfig) Keys(id string) (keys []string, err error)
```
Keys returns the app's config keys

#### func (*AppsConfig) List

```go
func (a *AppsConfig) List() (apps []string, err error)
```
List lists all the available apps

#### func (*AppsConfig) SetValue

```go
func (a *AppsConfig) SetValue(id, key, value string) error
```
SetValue set the config value for the given app's key

#### func (*AppsConfig) Value

```go
func (a *AppsConfig) Value(id, key string) (string, error)
```
Value get the config value for the given app's key

#### type AppsConfigI

```go
type AppsConfigI interface {
	List() (apps []string, err error)
	Keys(id string) (keys []string, err error)
	Value(id, key string) (string, error)
	SetValue(id, key, value string) error
	DeleteValue(id, key, value string) error
	Get() (map[string]map[string]string, error)
	Details(appID string) (map[string]string, error)
}
```

AppsConfigI available methods

#### type AppsI

```go
type AppsI interface {
	List() ([]string, error)
	ListEnabled() ([]string, error)
	ListDisabled() ([]string, error)
	Infos(name string) (types.App, error)
	Enable(name string) error
	Disable(name string) error
}
```
AppsI available methods

#### func (*Client) Monitoring

```go
func (c *Client) Monitoring() (*types.Monitoring, error)
```
Monitoring return nextcloud monitoring statistics

#### type GroupFolders

```go
type GroupFolders struct {
}
```

GroupFolders contains all Groups Folders available actions

#### func (*GroupFolders) AddGroup

```go
func (g *GroupFolders) AddGroup(folderID int, groupName string) error
```
AddGroup adds group to folder

#### func (*GroupFolders) Create

```go
func (g *GroupFolders) Create(name string) (id int, err error)
```
Create creates a group folder

#### func (*GroupFolders) Get

```go
func (g *GroupFolders) Get(id int) (types.GroupFolder, error)
```
Get returns the group folder details

#### func (*GroupFolders) List

```go
func (g *GroupFolders) List() (map[int]types.GroupFolder, error)
```
List returns the groups folders

#### func (*GroupFolders) RemoveGroup

```go
func (g *GroupFolders) RemoveGroup(folderID int, groupName string) error
```
RemoveGroup remove a group from the group folder

#### func (*GroupFolders) Rename

```go
func (g *GroupFolders) Rename(groupID int, name string) error
```
Rename renames the group folder

#### func (*GroupFolders) SetGroupPermissions

```go
func (g *GroupFolders) SetGroupPermissions(folderID int, groupName string, permission types.SharePermission) error
```
SetGroupPermissions set groups permissions

#### func (*GroupFolders) SetQuota

```go
func (g *GroupFolders) SetQuota(folderID int, quota int) error
```
SetQuota set quota on the group folder. quota in bytes, use -3 for unlimited

#### type GroupFoldersI

```go
type GroupFoldersI interface {
	List() (map[int]types.GroupFolder, error)
	Get(id int) (types.GroupFolder, error)
	Create(name string) (id int, err error)
	Rename(groupID int, name string) error
	AddGroup(folderID int, groupName string) error
	RemoveGroup(folderID int, groupName string) error
	SetGroupPermissions(folderID int, groupName string, permission types.SharePermission) error
	SetQuota(folderID int, quota int) error
}
```

GroupFoldersI available methods

#### type Groups

```go
type Groups struct {
}
```

Groups contains all Groups available actions

#### func (*Groups) Create

```go
func (g *Groups) Create(name string) error
```
Create creates a group

#### func (*Groups) Delete

```go
func (g *Groups) Delete(name string) error
```
Delete deletes the group

#### func (*Groups) List

```go
func (g *Groups) List() ([]string, error)
```
List lists the Nextcloud groups

#### func (*Groups) ListDetails

```go
func (g *Groups) ListDetails() ([]types.Group, error)
```
ListDetails lists the Nextcloud groups

#### func (*Groups) Search

```go
func (g *Groups) Search(search string) ([]string, error)
```
Search return the list of groups matching the search string

#### func (*Groups) SubAdminList

```go
func (g *Groups) SubAdminList(name string) ([]string, error)
```
SubAdminList lists the group's subadmins

#### func (*Groups) Users

```go
func (g *Groups) Users(name string) ([]string, error)
```
Users list the group's users

#### type GroupsI

```go
type GroupsI interface {
	List() ([]string, error)
	ListDetails() ([]types.Group, error)
	Users(name string) ([]string, error)
	Search(search string) ([]string, error)
	Create(name string) error
	Delete(name string) error
	SubAdminList(name string) ([]string, error)
}
```

GroupsI available methods

#### type Notifications

```go
type Notifications struct {
}
```

Notifications contains all Notifications available actions

#### func (*Notifications) AdminAvailable

```go
func (n *Notifications) AdminAvailable() error
```
AdminAvailable returns an error if the admin-notifications app is not installed

#### func (*Notifications) Available

```go
func (n *Notifications) Available() error
```
Available returns an error if the notifications app is not installed

#### func (*Notifications) Create

```go
func (n *Notifications) Create(userID, title, message string) error
```
Create creates a notification (if the user is an admin)

#### func (*Notifications) Delete

```go
func (n *Notifications) Delete(id int) error
```
Delete deletes the notification corresponding to the id

#### func (*Notifications) DeleteAll

```go
func (n *Notifications) DeleteAll() error
```
DeleteAll deletes all notifications

#### func (*Notifications) Get

```go
func (n *Notifications) Get(id int) (types.Notification, error)
```
Get returns the notification corresponding to the id

#### func (*Notifications) List

```go
func (n *Notifications) List() ([]types.Notification, error)
```
List returns all the notifications

#### type NotificationsI

```go
type NotificationsI interface {
	List() ([]types.Notification, error)
	Get(id int) (types.Notification, error)
	Delete(id int) error
	DeleteAll() error
	Create(userID, title, message string) error
	AdminAvailable() error
	Available() error
}
```

NotificationsI available methods

#### type Routes

```go
type Routes struct {
}
```

Routes references the available routes

#### type Shares

```go
type Shares struct {
}
```

Shares contains all Shares available actions

#### func (*Shares) Create

```go
func (s *Shares) Create(
	path string,
	shareType types.ShareType,
	permission types.SharePermission,
	shareWith string,
	publicUpload bool,
	password string,
) (types.Share, error)
```
Create create a share

#### func (*Shares) Delete

```go
func (s *Shares) Delete(shareID int) error
```
Delete Remove the given share.

#### func (*Shares) Get

```go
func (s *Shares) Get(shareID string) (types.Share, error)
```
Get information about a known Share

#### func (*Shares) GetFromPath

```go
func (s *Shares) GetFromPath(path string, reshares bool, subfiles bool) ([]types.Share, error)
```
GetFromPath return shares from a specific file or folder

#### func (*Shares) List

```go
func (s *Shares) List() ([]types.Share, error)
```
List list all shares of the logged in user

#### func (*Shares) Update

```go
func (s *Shares) Update(shareUpdate types.ShareUpdate) error
```
Update update share details expireDate expireDate expects a well formatted date
string, e.g. ‘YYYY-MM-DD’

#### func (*Shares) UpdateExpireDate

```go
func (s *Shares) UpdateExpireDate(shareID int, expireDate string) error
```
UpdateExpireDate updates the share's expire date expireDate expects a well
formatted date string, e.g. ‘YYYY-MM-DD’

#### func (*Shares) UpdatePassword

```go
func (s *Shares) UpdatePassword(shareID int, password string) error
```
UpdatePassword updates share password

#### func (*Shares) UpdatePermissions

```go
func (s *Shares) UpdatePermissions(shareID int, permissions types.SharePermission) error
```
UpdatePermissions update permissions

#### func (*Shares) UpdatePublicUpload

```go
func (s *Shares) UpdatePublicUpload(shareID int, public bool) error
```
UpdatePublicUpload enable or disable public upload

#### type SharesI

```go
type SharesI interface {
	List() ([]types.Share, error)
	GetFromPath(path string, reshares bool, subfiles bool) ([]types.Share, error)
	Get(shareID string) (types.Share, error)
	Create(
		path string,
		shareType types.ShareType,
		permission types.SharePermission,
		shareWith string,
		publicUpload bool,
		password string,
	) (types.Share, error)
	Delete(shareID int) error
	Update(shareUpdate types.ShareUpdate) error
	UpdateExpireDate(shareID int, expireDate string) error
	UpdatePublicUpload(shareID int, public bool) error
	UpdatePassword(shareID int, password string) error
	UpdatePermissions(shareID int, permissions types.SharePermission) error
}
```

SharesI available methods

#### type Users

```go
type Users struct {
}
```

Users contains all Users available actions

#### func (*Users) Create

```go
func (u *Users) Create(username string, password string, user *types.User) error
```
Create create a new user

#### func (*Users) CreateWithoutPassword

```go
func (u *Users) CreateWithoutPassword(username, email, displayName string) error
```
CreateWithoutPassword create a user without provisioning a password, the email
address must be provided to send an init password email

#### func (*Users) Delete

```go
func (u *Users) Delete(name string) error
```
Delete delete the user

#### func (*Users) Disable

```go
func (u *Users) Disable(name string) error
```
Disable disables the user

#### func (*Users) Enable

```go
func (u *Users) Enable(name string) error
```
Enable enables the user

#### func (*Users) Get

```go
func (u *Users) Get(name string) (*types.User, error)
```
Get return the details about the specified user

#### func (*Users) GroupAdd

```go
func (u *Users) GroupAdd(name string, group string) error
```
GroupAdd adds a the user to the group

#### func (*Users) GroupDemote

```go
func (u *Users) GroupDemote(name string, group string) error
```
GroupDemote demotes the user

#### func (*Users) GroupList

```go
func (u *Users) GroupList(name string) ([]string, error)
```
GroupList lists the user's groups

#### func (*Users) GroupPromote

```go
func (u *Users) GroupPromote(name string, group string) error
```
GroupPromote promotes the user as group admin

#### func (*Users) GroupRemove

```go
func (u *Users) GroupRemove(name string, group string) error
```
GroupRemove removes the user from the group

#### func (*Users) GroupSubAdminList

```go
func (u *Users) GroupSubAdminList(name string) ([]string, error)
```
GroupSubAdminList lists the groups where he is subadmin

#### func (*Users) List

```go
func (u *Users) List() ([]string, error)
```
List return the Nextcloud'user list

#### func (*Users) ListDetails

```go
func (u *Users) ListDetails() (map[string]types.User, error)
```
ListDetails return a map of user with details

#### func (*Users) Search

```go
func (u *Users) Search(search string) ([]string, error)
```
Search returns the users whose name match the search string

#### func (*Users) SendWelcomeEmail

```go
func (u *Users) SendWelcomeEmail(name string) error
```
SendWelcomeEmail (re)send the welcome mail to the user (return an error if the
user has not configured his email)

#### func (*Users) Update

```go
func (u *Users) Update(user *types.User) error
```
Update takes a *types.Users struct to update the user's information

#### func (*Users) UpdateAddress

```go
func (u *Users) UpdateAddress(name string, address string) error
```
UpdateAddress update the user's address

#### func (*Users) UpdateDisplayName

```go
func (u *Users) UpdateDisplayName(name string, displayName string) error
```
UpdateDisplayName update the user's display name

#### func (*Users) UpdateEmail

```go
func (u *Users) UpdateEmail(name string, email string) error
```
UpdateEmail update the user's email

#### func (*Users) UpdatePassword

```go
func (u *Users) UpdatePassword(name string, password string) error
```
UpdatePassword update the user's password

#### func (*Users) UpdatePhone

```go
func (u *Users) UpdatePhone(name string, phone string) error
```
UpdatePhone update the user's phone

#### func (*Users) UpdateQuota

```go
func (u *Users) UpdateQuota(name string, quota int) error
```
UpdateQuota update the user's quota (bytes)

#### func (*Users) UpdateTwitter

```go
func (u *Users) UpdateTwitter(name string, twitter string) error
```
UpdateTwitter update the user's twitter

#### func (*Users) UpdateWebSite

```go
func (u *Users) UpdateWebSite(name string, website string) error
```
UpdateWebSite update the user's website

#### type UsersI

```go
type UsersI interface {
	List() ([]string, error)
	ListDetails() (map[string]types.User, error)
	Get(name string) (*types.User, error)
	Search(search string) ([]string, error)
	Create(username string, password string, user *types.User) error
	CreateWithoutPassword(username, email, displayName string) error
	Delete(name string) error
	Enable(name string) error
	Disable(name string) error
	SendWelcomeEmail(name string) error
	Update(user *types.User) error
	UpdateEmail(name string, email string) error
	UpdateDisplayName(name string, displayName string) error
	UpdatePhone(name string, phone string) error
	UpdateAddress(name string, address string) error
	UpdateWebSite(name string, website string) error
	UpdateTwitter(name string, twitter string) error
	UpdatePassword(name string, password string) error
	UpdateQuota(name string, quota int) error
	GroupList(name string) ([]string, error)
	GroupAdd(name string, group string) error
	GroupRemove(name string, group string) error
	GroupPromote(name string, group string) error
	GroupDemote(name string, group string) error
	GroupSubAdminList(name string) ([]string, error)
}
```

UsersI available methods
