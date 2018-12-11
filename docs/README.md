# gonextcloud
--
    import "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"

Package gonextcloud is a Go client for the Nextcloud Provisioning API.

For more information about the Provisioning API, see the documentation:
https://docs.nextcloud.com/server/13/admin_manual/configuration_user/user_provisioning_api.html


### Usage

You use the library by creating a client object and calling methods on it.

For example, to list all the Nextcloud's instance users:

    package main

    import (
    	"fmt"
    	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
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

    	users, err := c.Users().List()
    	if err != nil {
    		panic(err)
    	}
    	fmt.Println("Users :", users)
    }

## Usage

#### type Apps

    type Apps struct {
    }


Apps contains all Apps available actions

#### func (*Apps) Disable

    func (a *Apps) Disable(name string) error

Disable disables an app

#### func (*Apps) Enable

    func (a *Apps) Enable(name string) error

Enable enables an app

#### func (*Apps) Infos

    func (a *Apps) Infos(name string) (types.App, error)

Infos return the app's details

#### func (*Apps) List

    func (a *Apps) List() ([]string, error)

List return the list of the Nextcloud Apps

#### func (*Apps) ListDisabled

    func (a *Apps) ListDisabled() ([]string, error)

ListDisabled lists the disabled apps

#### func (*Apps) ListEnabled

    func (a *Apps) ListEnabled() ([]string, error)

ListEnabled lists the enabled apps

#### type AppsConfig

    type AppsConfig struct {
    }


AppsConfig contains all Apps Configuration available actions

#### func (*AppsConfig) DeleteValue

    func (a *AppsConfig) DeleteValue(id, key, value string) error

DeleteValue delete the config value and (!! be careful !!) the key

#### func (*AppsConfig) Details

    func (a *AppsConfig) Details(appID string) (map[string]string, error)

Details returns all the config's key, values pair of the app

#### func (*AppsConfig) Get

    func (a *AppsConfig) Get() (map[string]map[string]string, error)

Get returns all apps AppConfigDetails

#### func (*AppsConfig) Keys

    func (a *AppsConfig) Keys(id string) (keys []string, err error)

Keys returns the app's config keys

#### func (*AppsConfig) List

    func (a *AppsConfig) List() (apps []string, err error)

List lists all the available apps

#### func (*AppsConfig) SetValue

    func (a *AppsConfig) SetValue(id, key, value string) error

SetValue set the config value for the given app's key

#### func (*AppsConfig) Value

    func (a *AppsConfig) Value(id, key string) (string, error)

Value get the config value for the given app's key

#### type Client

    type Client struct {
    }


Client is the API client that performs all operations against a Nextcloud
server.

#### func  NewClient

    func NewClient(hostname string) (*Client, error)

NewClient create a new Client from the Nextcloud Instance URL

#### func (*Client) Apps

    func (c *Client) Apps() types.Apps


#### func (*Client) AppsConfig

    func (c *Client) AppsConfig() types.AppsConfig


#### func (*Client) GroupFolders

    func (c *Client) GroupFolders() types.GroupFolders


#### func (*Client) Groups

    func (c *Client) Groups() types.Groups


#### func (*Client) Login

    func (c *Client) Login(username string, password string) error

Login perform login and create a session with the Nextcloud API.

#### func (*Client) Logout

    func (c *Client) Logout() error

Logout logs out from the Nextcloud API, close the session and delete session's
cookie

#### func (*Client) Monitoring

    func (c *Client) Monitoring() (*types.Monitoring, error)

Monitoring return nextcloud monitoring statistics

#### func (*Client) Notifications

    func (c *Client) Notifications() types.Notifications


#### func (*Client) Shares

    func (c *Client) Shares() types.Shares


#### func (*Client) Users

    func (c *Client) Users() types.Users


#### type GroupFolders

    type GroupFolders struct {
    }


GroupFolders contains all Groups Folders available actions

#### func (*GroupFolders) AddGroup

    func (g *GroupFolders) AddGroup(folderID int, groupName string) error

AddGroup adds group to folder

#### func (*GroupFolders) Create

    func (g *GroupFolders) Create(name string) (id int, err error)

Create creates a group folder

#### func (*GroupFolders) Get

    func (g *GroupFolders) Get(id int) (types.GroupFolder, error)

Get returns the group folder details

#### func (*GroupFolders) List

    func (g *GroupFolders) List() (map[int]types.GroupFolder, error)

List returns the groups folders

#### func (*GroupFolders) RemoveGroup

    func (g *GroupFolders) RemoveGroup(folderID int, groupName string) error

RemoveGroup remove a group from the group folder

#### func (*GroupFolders) Rename

    func (g *GroupFolders) Rename(groupID int, name string) error

Rename renames the group folder

#### func (*GroupFolders) SetGroupPermissions

    func (g *GroupFolders) SetGroupPermissions(folderID int, groupName string, permission types.SharePermission) error

SetGroupPermissions set groups permissions

#### func (*GroupFolders) SetQuota

    func (g *GroupFolders) SetQuota(folderID int, quota int) error

SetQuota set quota on the group folder. quota in bytes, use -3 for unlimited

#### type Groups

    type Groups struct {
    }


Groups contains all Groups available actions

#### func (*Groups) Create

    func (g *Groups) Create(name string) error

Create creates a group

#### func (*Groups) Delete

    func (g *Groups) Delete(name string) error

Delete deletes the group

#### func (*Groups) List

    func (g *Groups) List() ([]string, error)

List lists the Nextcloud groups

#### func (*Groups) ListDetails

    func (g *Groups) ListDetails() ([]types.Group, error)

ListDetails lists the Nextcloud groups

#### func (*Groups) Search

    func (g *Groups) Search(search string) ([]string, error)

Search return the list of groups matching the search string

#### func (*Groups) SubAdminList

    func (g *Groups) SubAdminList(name string) ([]string, error)

SubAdminList lists the group's subadmins

#### func (*Groups) Users

    func (g *Groups) Users(name string) ([]string, error)

Users list the group's users

#### type Notifications

    type Notifications struct {
    }


Notifications contains all Notifications available actions

#### func (*Notifications) AdminAvailable

    func (n *Notifications) AdminAvailable() error

AdminAvailable returns an error if the admin-notifications app is not installed

#### func (*Notifications) Available

    func (n *Notifications) Available() error

Available returns an error if the notifications app is not installed

#### func (*Notifications) Create

    func (n *Notifications) Create(userID, title, message string) error

Create creates a notification (if the user is an admin)

#### func (*Notifications) Delete

    func (n *Notifications) Delete(id int) error

Delete deletes the notification corresponding to the id

#### func (*Notifications) DeleteAll

    func (n *Notifications) DeleteAll() error

DeleteAll deletes all notifications

#### func (*Notifications) Get

    func (n *Notifications) Get(id int) (types.Notification, error)

Get returns the notification corresponding to the id

#### func (*Notifications) List

    func (n *Notifications) List() ([]types.Notification, error)

List returns all the notifications

#### type Routes

    type Routes struct {
    }


Routes references the available routes

#### type Shares

    type Shares struct {
    }


Shares contains all Shares available actions

#### func (*Shares) Create

    func (s *Shares) Create(
    	path string,
    	shareType types.ShareType,
    	permission types.SharePermission,
    	shareWith string,
    	publicUpload bool,
    	password string,
    ) (types.Share, error)

Create create a share

#### func (*Shares) Delete

    func (s *Shares) Delete(shareID int) error

Delete Remove the given share.

#### func (*Shares) Get

    func (s *Shares) Get(shareID string) (types.Share, error)

Get information about a known Share

#### func (*Shares) GetFromPath

    func (s *Shares) GetFromPath(path string, reshares bool, subfiles bool) ([]types.Share, error)

GetFromPath return shares from a specific file or folder

#### func (*Shares) List

    func (s *Shares) List() ([]types.Share, error)

List list all shares of the logged in user

#### func (*Shares) Update

    func (s *Shares) Update(shareUpdate types.ShareUpdate) error

Update update share details expireDate expireDate expects a well formatted date
string, e.g. ‘YYYY-MM-DD’

#### func (*Shares) UpdateExpireDate

    func (s *Shares) UpdateExpireDate(shareID int, expireDate string) error

UpdateExpireDate updates the share's expire date expireDate expects a well
formatted date string, e.g. ‘YYYY-MM-DD’

#### func (*Shares) UpdatePassword

    func (s *Shares) UpdatePassword(shareID int, password string) error

UpdatePassword updates share password

#### func (*Shares) UpdatePermissions

    func (s *Shares) UpdatePermissions(shareID int, permissions types.SharePermission) error

UpdatePermissions update permissions

#### func (*Shares) UpdatePublicUpload

    func (s *Shares) UpdatePublicUpload(shareID int, public bool) error

UpdatePublicUpload enable or disable public upload

#### type Users

    type Users struct {
    }


Users contains all Users available actions

#### func (*Users) Create

    func (u *Users) Create(username string, password string, user *types.User) error

Create create a new user

#### func (*Users) CreateWithoutPassword

    func (u *Users) CreateWithoutPassword(username, email, displayName string) error

CreateWithoutPassword create a user without provisioning a password, the email
address must be provided to send an init password email

#### func (*Users) Delete

    func (u *Users) Delete(name string) error

Delete delete the user

#### func (*Users) Disable

    func (u *Users) Disable(name string) error

Disable disables the user

#### func (*Users) Enable

    func (u *Users) Enable(name string) error

Enable enables the user

#### func (*Users) Get

    func (u *Users) Get(name string) (*types.User, error)

Get return the details about the specified user

#### func (*Users) GroupAdd

    func (u *Users) GroupAdd(name string, group string) error

GroupAdd adds a the user to the group

#### func (*Users) GroupDemote

    func (u *Users) GroupDemote(name string, group string) error

GroupDemote demotes the user

#### func (*Users) GroupList

    func (u *Users) GroupList(name string) ([]string, error)

GroupList lists the user's groups

#### func (*Users) GroupPromote

    func (u *Users) GroupPromote(name string, group string) error

GroupPromote promotes the user as group admin

#### func (*Users) GroupRemove

    func (u *Users) GroupRemove(name string, group string) error

GroupRemove removes the user from the group

#### func (*Users) GroupSubAdminList

    func (u *Users) GroupSubAdminList(name string) ([]string, error)

GroupSubAdminList lists the groups where he is subadmin

#### func (*Users) List

    func (u *Users) List() ([]string, error)

List return the Nextcloud'user list

#### func (*Users) ListDetails

    func (u *Users) ListDetails() (map[string]types.User, error)

ListDetails return a map of user with details

#### func (*Users) Search

    func (u *Users) Search(search string) ([]string, error)

Search returns the users whose name match the search string

#### func (*Users) SendWelcomeEmail

    func (u *Users) SendWelcomeEmail(name string) error

SendWelcomeEmail (re)send the welcome mail to the user (return an error if the
user has not configured his email)

#### func (*Users) Update

    func (u *Users) Update(user *types.User) error

Update takes a *types.Users struct to update the user's information

#### func (*Users) UpdateAddress

    func (u *Users) UpdateAddress(name string, address string) error

UpdateAddress update the user's address

#### func (*Users) UpdateDisplayName

    func (u *Users) UpdateDisplayName(name string, displayName string) error

UpdateDisplayName update the user's display name

#### func (*Users) UpdateEmail

    func (u *Users) UpdateEmail(name string, email string) error

UpdateEmail update the user's email

#### func (*Users) UpdatePassword

    func (u *Users) UpdatePassword(name string, password string) error

UpdatePassword update the user's password

#### func (*Users) UpdatePhone

    func (u *Users) UpdatePhone(name string, phone string) error

UpdatePhone update the user's phone

#### func (*Users) UpdateQuota

    func (u *Users) UpdateQuota(name string, quota int) error

UpdateQuota update the user's quota (bytes)

#### func (*Users) UpdateTwitter

    func (u *Users) UpdateTwitter(name string, twitter string) error

UpdateTwitter update the user's twitter

#### func (*Users) UpdateWebSite

    func (u *Users) UpdateWebSite(name string, website string) error

UpdateWebSite update the user's website

# types
--
    import "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"


## Usage

```go
const (
	UserShare           ShareType = 0
	GroupShare          ShareType = 1
	PublicLinkShare     ShareType = 3
	FederatedCloudShare ShareType = 6

	ReadPermission    SharePermission = 1
	UpdatePermission  SharePermission = 2
	CreatePermission  SharePermission = 4
	DeletePermission  SharePermission = 8
	ReSharePermission SharePermission = 16
	AllPermissions    SharePermission = 31
)
```

#### type APIError

```go
type APIError struct {
	Code    int
	Message string
}
```

APIError contains the returned error code and message from the Nextcloud's API

#### func  ErrorFromMeta

```go
func ErrorFromMeta(meta Meta) *APIError
```
ErrorFromMeta return a types.APIError from the Response's types.Meta

#### func (*APIError) Error

```go
func (e *APIError) Error() string
```
Error return the types.APIError string

#### type ActiveUsers

```go
type ActiveUsers struct {
	Last5Minutes int `json:"last5minutes"`
	Last1Hour    int `json:"last1hour"`
	Last24Hours  int `json:"last24hours"`
}
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

App

#### type AppConfigResponse

```go
type AppConfigResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Data []string `json:"data"`
		} `json:"data"`
	} `json:"ocs"`
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

AppListResponse

#### type AppResponse

```go
type AppResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data App  `json:"data"`
	} `json:"ocs"`
}
```

AppResponse

#### type AppcConfigValueResponse

```go
type AppcConfigValueResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Data string `json:"data"`
		} `json:"data"`
	} `json:"ocs"`
}
```


#### type Apps

```go
type Apps interface {
	List() ([]string, error)
	ListEnabled() ([]string, error)
	ListDisabled() ([]string, error)
	Infos(name string) (App, error)
	Enable(name string) error
	Disable(name string) error
}
```

Apps available methods

#### type AppsConfig

```go
type AppsConfig interface {
	List() (apps []string, err error)
	Keys(id string) (keys []string, err error)
	Value(id, key string) (string, error)
	SetValue(id, key, value string) error
	DeleteValue(id, key, value string) error
	Get() (map[string]map[string]string, error)
	Details(appID string) (map[string]string, error)
}
```

AppsConfig available methods

#### type Auth

```go
type Auth interface {
	Login(username string, password string) error
	Logout() error
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

BaseResponse

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
	Ocm struct {
		Enabled    bool   `json:"enabled"`
		APIVersion string `json:"apiVersion"`
		EndPoint   string `json:"endPoint"`
		ShareTypes []struct {
			Name      string `json:"name"`
			Protocols struct {
				Webdav string `json:"webdav"`
			} `json:"protocols"`
		} `json:"shareTypes"`
	} `json:"ocm"`
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
		DefaultPermissions int `json:"default_permissions"`
		Federation         struct {
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
		OcsEndpoints       []string `json:"ocs-endpoints"`
		Push               []string `json:"push"`
		AdminNotifications []string `json:"admin-notifications"`
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
	Registration struct {
		Enabled  bool   `json:"enabled"`
		APIRoot  string `json:"apiRoot"`
		APILevel string `json:"apiLevel"`
	} `json:"registration"`
}
```

Capabilities

#### type CapabilitiesResponse

```go
type CapabilitiesResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Version      Version      `json:"version"`
			Capabilities Capabilities `json:"capabilities"`
		} `json:"data"`
	} `json:"ocs"`
}
```

CapabilitiesResponse

#### type Client

```go
type Client interface {
	Apps() Apps
	AppsConfig() AppsConfig
	GroupFolders() GroupFolders
	Notifications() Notifications
	Shares() Shares
	Users() Users
	Groups() Groups
}
```

Client is the main client interface

#### type ErrorResponse

```go
type ErrorResponse struct {
	Ocs struct {
		Meta Meta          `json:"meta"`
		Data []interface{} `json:"data"`
	} `json:"ocs"`
}
```

ErrorResponse

#### type Group

```go
type Group struct {
	ID          string `json:"id"`
	Displayname string `json:"displayname"`
	UserCount   int    `json:"usercount"`
	Disabled    int    `json:"disabled"`
	CanAdd      bool   `json:"canAdd"`
	CanRemove   bool   `json:"canRemove"`
}
```

Group

#### type GroupFolder

```go
type GroupFolder struct {
	ID         int                        `json:"id"`
	MountPoint string                     `json:"mount_point"`
	Groups     map[string]SharePermission `json:"groups"`
	Quota      int                        `json:"quota"`
	Size       int                        `json:"size"`
}
```


#### type GroupFolderBadFormatGroups

```go
type GroupFolderBadFormatGroups struct {
	ID         int               `json:"id"`
	MountPoint string            `json:"mount_point"`
	Groups     map[string]string `json:"groups"`
	Quota      string            `json:"quota"`
	Size       int               `json:"size"`
}
```


#### func (*GroupFolderBadFormatGroups) FormatGroupFolder

```go
func (gf *GroupFolderBadFormatGroups) FormatGroupFolder() GroupFolder
```

#### type GroupFolderBadFormatIDAndGroups

```go
type GroupFolderBadFormatIDAndGroups struct {
	ID         string            `json:"id"`
	MountPoint string            `json:"mount_point"`
	Groups     map[string]string `json:"groups"`
	Quota      string            `json:"quota"`
	Size       int               `json:"size"`
}
```


#### func (*GroupFolderBadFormatIDAndGroups) FormatGroupFolder

```go
func (gf *GroupFolderBadFormatIDAndGroups) FormatGroupFolder() GroupFolder
```

#### type GroupFolders

```go
type GroupFolders interface {
	List() (map[int]GroupFolder, error)
	Get(id int) (GroupFolder, error)
	Create(name string) (id int, err error)
	Rename(groupID int, name string) error
	AddGroup(folderID int, groupName string) error
	RemoveGroup(folderID int, groupName string) error
	SetGroupPermissions(folderID int, groupName string, permission SharePermission) error
	SetQuota(folderID int, quota int) error
}
```

GroupFolders available methods

#### type GroupFoldersCreateResponse

```go
type GroupFoldersCreateResponse struct {
	Ocs struct {
		Meta Meta                            `json:"meta"`
		Data GroupFolderBadFormatIDAndGroups `json:"data"`
	} `json:"ocs"`
}
```


#### type GroupFoldersListResponse

```go
type GroupFoldersListResponse struct {
	Ocs struct {
		Meta Meta                                       `json:"meta"`
		Data map[string]GroupFolderBadFormatIDAndGroups `json:"data"`
	} `json:"ocs"`
}
```


#### type GroupFoldersResponse

```go
type GroupFoldersResponse struct {
	Ocs struct {
		Meta Meta                       `json:"meta"`
		Data GroupFolderBadFormatGroups `json:"data"`
	} `json:"ocs"`
}
```


#### type GroupListDetailsResponse

```go
type GroupListDetailsResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Groups []Group `json:"groups"`
		} `json:"data"`
	} `json:"ocs"`
}
```

GroupListDetailsResponse

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

GroupListResponse

#### type Groups

```go
type Groups interface {
	List() ([]string, error)
	ListDetails() ([]Group, error)
	Users(name string) ([]string, error)
	Search(search string) ([]string, error)
	Create(name string) error
	Delete(name string) error
	SubAdminList(name string) ([]string, error)
}
```

Groups available methods

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

Meta

#### type Monitoring

```go
type Monitoring struct {
	Nextcloud struct {
		System  System  `json:"system"`
		Storage Storage `json:"storage"`
		Shares  struct {
			NumShares               int `json:"num_shares"`
			NumSharesUser           int `json:"num_shares_user"`
			NumSharesGroups         int `json:"num_shares_groups"`
			NumSharesLink           int `json:"num_shares_link"`
			NumSharesLinkNoPassword int `json:"num_shares_link_no_password"`
			NumFedSharesSent        int `json:"num_fed_shares_sent"`
			NumFedSharesReceived    int `json:"num_fed_shares_received"`
		} `json:"shares"`
	} `json:"nextcloud"`
	Server struct {
		Webserver string `json:"webserver"`
		Php       struct {
			Version           string `json:"version"`
			MemoryLimit       int    `json:"memory_limit"`
			MaxExecutionTime  int    `json:"max_execution_time"`
			UploadMaxFilesize int    `json:"upload_max_filesize"`
		} `json:"php"`
		Database struct {
			Type    string `json:"type"`
			Version string `json:"version"`
			Size    int    `json:"size"`
		} `json:"database"`
	} `json:"server"`
	ActiveUsers ActiveUsers `json:"activeUsers"`
}
```


#### type MonitoringResponse

```go
type MonitoringResponse struct {
	Ocs struct {
		Meta Meta       `json:"meta"`
		Data Monitoring `json:"data"`
	} `json:"ocs"`
}
```


#### type Notification

```go
type Notification struct {
	NotificationID        int           `json:"notification_id"`
	App                   string        `json:"app"`
	User                  string        `json:"user"`
	Datetime              time.Time     `json:"datetime"`
	ObjectType            string        `json:"object_type"`
	ObjectID              string        `json:"object_id"`
	Subject               string        `json:"subject"`
	Message               string        `json:"message"`
	Link                  string        `json:"link"`
	SubjectRich           string        `json:"subjectRich"`
	SubjectRichParameters []interface{} `json:"subjectRichParameters"`
	MessageRich           string        `json:"messageRich"`
	MessageRichParameters []interface{} `json:"messageRichParameters"`
	Icon                  string        `json:"icon"`
	Actions               []interface{} `json:"actions"`
}
```


#### type NotificationResponse

```go
type NotificationResponse struct {
	Ocs struct {
		Meta Meta         `json:"meta"`
		Data Notification `json:"data"`
	} `json:"ocs"`
}
```


#### type Notifications

```go
type Notifications interface {
	List() ([]Notification, error)
	Get(id int) (Notification, error)
	Delete(id int) error
	DeleteAll() error
	Create(userID, title, message string) error
	AdminAvailable() error
	Available() error
}
```

Notifications available methods

#### type NotificationsListResponse

```go
type NotificationsListResponse struct {
	Ocs struct {
		Meta Meta           `json:"meta"`
		Data []Notification `json:"data"`
	} `json:"ocs"`
}
```


#### type Share

```go
type Share struct {
	ID                   string      `json:"id"`
	ShareType            int         `json:"share_type"`
	UIDOwner             string      `json:"uid_owner"`
	DisplaynameOwner     string      `json:"displayname_owner"`
	Permissions          int         `json:"permissions"`
	Stime                int         `json:"stime"`
	Parent               interface{} `json:"parent"`
	Expiration           string      `json:"expiration"`
	Token                string      `json:"token"`
	UIDFileOwner         string      `json:"uid_file_owner"`
	DisplaynameFileOwner string      `json:"displayname_file_owner"`
	Path                 string      `json:"path"`
	ItemType             string      `json:"item_type"`
	Mimetype             string      `json:"mimetype"`
	StorageID            string      `json:"storage_id"`
	Storage              int         `json:"storage"`
	ItemSource           int         `json:"item_source"`
	FileSource           int         `json:"file_source"`
	FileParent           int         `json:"file_parent"`
	FileTarget           string      `json:"file_target"`
	ShareWith            string      `json:"share_with"`
	ShareWithDisplayname string      `json:"share_with_displayname"`
	MailSend             int         `json:"mail_send"`
	Tags                 []string    `json:"tags"`
}
```


#### type SharePermission

```go
type SharePermission int
```


#### type ShareType

```go
type ShareType int
```


#### type ShareUpdate

```go
type ShareUpdate struct {
	ShareID      int
	Permissions  SharePermission
	Password     string
	PublicUpload bool
	ExpireDate   string
}
```


#### type Shares

```go
type Shares interface {
	List() ([]Share, error)
	GetFromPath(path string, reshares bool, subfiles bool) ([]Share, error)
	Get(shareID string) (Share, error)
	Create(
		path string,
		shareType ShareType,
		permission SharePermission,
		shareWith string,
		publicUpload bool,
		password string,
	) (Share, error)
	Delete(shareID int) error
	Update(shareUpdate ShareUpdate) error
	UpdateExpireDate(shareID int, expireDate string) error
	UpdatePublicUpload(shareID int, public bool) error
	UpdatePassword(shareID int, password string) error
	UpdatePermissions(shareID int, permissions SharePermission) error
}
```

Shares available methods

#### type SharesListResponse

```go
type SharesListResponse struct {
	Ocs struct {
		Meta Meta    `json:"meta"`
		Data []Share `json:"data"`
	} `json:"ocs"`
}
```


#### type SharesResponse

```go
type SharesResponse struct {
	Ocs struct {
		Meta Meta  `json:"meta"`
		Data Share `json:"data"`
	} `json:"ocs"`
}
```


#### type Storage

```go
type Storage struct {
	NumUsers         int `json:"num_users"`
	NumFiles         int `json:"num_files"`
	NumStorages      int `json:"num_storages"`
	NumStoragesLocal int `json:"num_storages_local"`
	NumStoragesHome  int `json:"num_storages_home"`
	NumStoragesOther int `json:"num_storages_other"`
}
```


#### type System

```go
type System struct {
	Version             string    `json:"version"`
	Theme               string    `json:"theme"`
	EnableAvatars       string    `json:"enable_avatars"`
	EnablePreviews      string    `json:"enable_previews"`
	MemcacheLocal       string    `json:"memcache.local"`
	MemcacheDistributed string    `json:"memcache.distributed"`
	FilelockingEnabled  string    `json:"filelocking.enabled"`
	MemcacheLocking     string    `json:"memcache.locking"`
	Debug               string    `json:"debug"`
	Freespace           int64     `json:"freespace"`
	Cpuload             []float32 `json:"cpuload"`
	MemTotal            int       `json:"mem_total"`
	MemFree             int       `json:"mem_free"`
	SwapTotal           int       `json:"swap_total"`
	SwapFree            int       `json:"swap_free"`
}
```


#### type UpdateError

```go
type UpdateError struct {
	Field string
	Error error
}
```

UpdateError contains the user's field and corresponding error

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
	Language    string   `json:"language,omitempty"`

	StorageLocation string        `json:"storageLocation,omitempty"`
	LastLogin       int64         `json:"lastLogin,omitempty"`
	Backend         string        `json:"backend,omitempty"`
	Subadmin        []interface{} `json:"subadmin,omitempty"`
	Locale          string        `json:"locale,omitempty"`
}
```

Users

#### type UserListDetailsResponse

```go
type UserListDetailsResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Users map[string]User `json:"users"`
		} `json:"data"`
	} `json:"ocs"`
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

UserListResponse

#### type UserResponse

```go
type UserResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data User `json:"data"`
	} `json:"ocs"`
}
```

UserResponse

#### type UserUpdateError

```go
type UserUpdateError struct {
	Errors map[string]error
}
```

UpdateError contains the errors resulting from a UserUpdate or a UserCreateFull
call

#### func  NewUpdateError

```go
func NewUpdateError(errors chan UpdateError) *UserUpdateError
```
NewUpdateError returns an UpdateError based on an UpdateError channel

#### func (*UserUpdateError) Error

```go
func (e *UserUpdateError) Error() string
```

#### type Users

```go
type Users interface {
	List() ([]string, error)
	ListDetails() (map[string]User, error)
	Get(name string) (*User, error)
	Search(search string) ([]string, error)
	Create(username string, password string, user *User) error
	CreateWithoutPassword(username, email, displayName, quota, language string, groups ...string) error
	Delete(name string) error
	Enable(name string) error
	Disable(name string) error
	SendWelcomeEmail(name string) error
	Update(user *User) error
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

Users available methods

#### type Version

```go
type Version struct {
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Micro   int    `json:"micro"`
	String  string `json:"string"`
	Edition string `json:"edition"`
}
```

