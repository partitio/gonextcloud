package gonextcloud

import (
	"io"
	"os"
	"path/filepath"
)

// NewClient create a new client
func NewClient(hostname string) (Client, error) {
	return newClient(hostname)
}

// Client is the main client interface
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

type Auth interface {
	Login(username string, password string) error
	Logout() error
}

// Apps available methods
type Apps interface {
	List() ([]string, error)
	ListEnabled() ([]string, error)
	ListDisabled() ([]string, error)
	Infos(name string) (App, error)
	Enable(name string) error
	Disable(name string) error
}

// AppsConfig available methods
type AppsConfig interface {
	List() (apps []string, err error)
	Keys(id string) (keys []string, err error)
	Value(id, key string) (string, error)
	SetValue(id, key, value string) error
	DeleteValue(id, key, value string) error
	Get() (map[string]map[string]string, error)
	Details(appID string) (map[string]string, error)
}

// Groups available methods
type Groups interface {
	List() ([]string, error)
	ListDetails(search string) ([]Group, error)
	Users(name string) ([]string, error)
	Search(search string) ([]string, error)
	Create(name string) error
	Delete(name string) error
	SubAdminList(name string) ([]string, error)
}

// GroupFolders available methods
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

// Notifications available methods
type Notifications interface {
	List() ([]Notification, error)
	Get(id int) (Notification, error)
	Delete(id int) error
	DeleteAll() error
	Create(userID, title, message string) error
	AdminAvailable() error
	Available() error
}

// Shares available methods
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

// Users available methods
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

// WebDav available methods
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
