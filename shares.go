package gonextcloud

type ShareType int
type SharePermission int

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

type ShareUpdate struct {
	ShareID      int
	Permissions  SharePermission
	Password     string
	PublicUpload bool
	ExpireDate   string
}

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
