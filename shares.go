package gonextcloud

// ShareType is the nextcloud shares types enum :
type ShareType int

// SharePermission is the nextcloud share permissions enum
type SharePermission int

const (
	// UserShare is a file or folder shared with other user(s)
	UserShare ShareType = 0
	// GroupShare is a file or folder shared with a group
	GroupShare ShareType = 1
	// PublicLinkShare is a file or folder shared through public link
	PublicLinkShare ShareType = 3
	// FederatedCloudShare is a file or folder shared through federated cloud
	FederatedCloudShare ShareType = 6

	// ReadPermission grant read permission
	ReadPermission SharePermission = 1
	// UpdatePermission grant update permission
	UpdatePermission SharePermission = 2
	// CreatePermission grant create permission
	CreatePermission SharePermission = 4
	// DeletePermission grant delete permission
	DeletePermission SharePermission = 8
	// ReSharePermission grant resharing permission
	ReSharePermission SharePermission = 16
	// AllPermissions grant all permissions
	AllPermissions SharePermission = 31
)

// ShareUpdate contains the data required in order to update a nextcloud share
type ShareUpdate struct {
	ShareID      int
	Permissions  SharePermission
	Password     string
	PublicUpload bool
	ExpireDate   string
}

// Share is a nextcloud share
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
