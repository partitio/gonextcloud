![Nextcloud](https://upload.wikimedia.org/wikipedia/commons/thumb/6/60/Nextcloud_Logo.svg/640px-Nextcloud_Logo.svg.png)

[![pipeline status](https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/badges/master/pipeline.svg)](https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/commits/master)
[![coverage report](https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/badges/master/coverage.svg)](https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud)](https://goreportcard.com/report/gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud)
[![GoDoc](https://godoc.org/gitlab.com/gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud?status.svg)](https://godoc.org/gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud)
# GoNextcloud

A simple Client for Nextcloud's Provisioning API in Go.

For more information about the Provisioning API, see the documentation:
https://docs.nextcloud.com/server/13/admin_manual/configuration_user/user_provisioning_api.html

## Usage

```go
import "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
```

You use the library by creating a client object and calling methods on it.

For example, to list all the Nextcloud's instance users:
```go
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
```

## Run tests
Configure the tests for your instance by editing [example.config.yml](example.config.yml) and renaming it config.yml

then run the tests :
```bash
$ go test -v .
```

## Docs

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
	WebDav() WebDav
	Monitoring() (*Monitoring, error)
	Login(username string, password string) error
	Logout() error
}
```

Client is the main client interface

#### func  NewClient

```go
func NewClient(hostname string) (Client, error)
```
NewClient create a new client

#### ShareType and SharePermission

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

#### type Groups

```go
type Groups interface {
	List() ([]string, error)
	ListDetails(search string) ([]Group, error)
	Users(name string) ([]string, error)
	Search(search string) ([]string, error)
	Create(name string) error
	Delete(name string) error
	SubAdminList(name string) ([]string, error)
}
```

Groups available methods

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

#### type Quota

```go
type Quota struct {
	Free     int64   `json:"free"`
	Used     int64   `json:"used"`
	Total    int64   `json:"total"`
	Relative float64 `json:"relative"`
	Quota    int64   `json:"quota"`
}
```


#### func (*Quota) String

```go
func (q *Quota) String() string
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

```

#### type Users

```go
type Users interface {
	List() ([]string, error)
	ListDetails() (map[string]UserDetails, error)
	Get(name string) (*UserDetails, error)
	Search(search string) ([]string, error)
	Create(username string, password string, user *UserDetails) error
	CreateWithoutPassword(username, email, displayName, quota, language string, groups ...string) error
	CreateBatchWithoutPassword(users []User) error
	Delete(name string) error
	Enable(name string) error
	Disable(name string) error
	SendWelcomeEmail(name string) error
	Update(user *UserDetails) error
	UpdateEmail(name string, email string) error
	UpdateDisplayName(name string, displayName string) error
	UpdatePhone(name string, phone string) error
	UpdateAddress(name string, address string) error
	UpdateWebSite(name string, website string) error
	UpdateTwitter(name string, twitter string) error
	UpdatePassword(name string, password string) error
	UpdateQuota(name string, quota int64) error
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


#### type WebDav

```go
type WebDav interface {
	// ReadDir reads the contents of a remote directory
	ReadDir(path string) ([]os.FileInfo, error)
	// Stat returns the file stats for a specified path
	Stat(path string) (os.FileInfo, error)
	// Remove removes a remote file
	Remove(path string) error
	// RemoveAll removes remote files
	RemoveAll(path string) error
	// Mkdir makes a directory
	Mkdir(path string, _ os.FileMode) error
	// MkdirAll like mkdir -p, but for webdav
	MkdirAll(path string, _ os.FileMode) error
	// Rename moves a file from A to B
	Rename(oldpath, newpath string, overwrite bool) error
	// Copy copies a file from A to B
	Copy(oldpath, newpath string, overwrite bool) error
	// Read reads the contents of a remote file
	Read(path string) ([]byte, error)
	// ReadStream reads the stream for a given path
	ReadStream(path string) (io.ReadCloser, error)
	// Write writes data to a given path
	Write(path string, data []byte, _ os.FileMode) error
	// WriteStream writes a stream
	WriteStream(path string, stream io.Reader, _ os.FileMode) error

	// Walk walks the file tree rooted at root, calling walkFn for each file or
	// directory in the tree, including root.
	Walk(path string, walkFunc filepath.WalkFunc) error
}
```

WebDav available methods


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
	Username    string
	Email       string
	DisplayName string
	Quota       string
	Language    string
	Groups      []string
}
```

User encapsulate the data needed to create a new Nextcloud's User

#### type UserDetails

```go
type UserDetails struct {
	Enabled     bool     `json:"enabled"`
	ID          string   `json:"id"`
	Quota       Quota    `json:"quota"`
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

UserDetails is the raw Nextcloud User response

#### type UserUpdateError

```go
type UserUpdateError struct {
	Errors map[string]error
}
```

UpdateError contains the errors resulting from a UserUpdate or a UserCreateFull
call

#### func (*UserUpdateError) Error

```go
func (e *UserUpdateError) Error() string
