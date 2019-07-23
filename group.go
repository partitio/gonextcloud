package gonextcloud

// Group is a Nextcloud group
type Group struct {
	ID          string `json:"id"`
	Displayname string `json:"displayname"`
	UserCount   int    `json:"usercount"`
	Disabled    int    `json:"disabled"`
	CanAdd      bool   `json:"canAdd"`
	CanRemove   bool   `json:"canRemove"`
}
