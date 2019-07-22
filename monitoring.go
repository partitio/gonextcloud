package gonextcloud

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

type ActiveUsers struct {
	Last5Minutes int `json:"last5minutes"`
	Last1Hour    int `json:"last1hour"`
	Last24Hours  int `json:"last24hours"`
}

type Storage struct {
	NumUsers         int `json:"num_users"`
	NumFiles         int `json:"num_files"`
	NumStorages      int `json:"num_storages"`
	NumStoragesLocal int `json:"num_storages_local"`
	NumStoragesHome  int `json:"num_storages_home"`
	NumStoragesOther int `json:"num_storages_other"`
}
