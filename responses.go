package gonextcloud

//meta
type meta struct {
	Status       string `json:"status"`
	Statuscode   int    `json:"statuscode"`
	Message      string `json:"message"`
	Totalitems   string `json:"totalitems"`
	Itemsperpage string `json:"itemsperpage"`
}

//errorResponse
type errorResponse struct {
	Ocs struct {
		Meta meta          `json:"meta"`
		Data []interface{} `json:"data"`
	} `json:"ocs"`
}

//userListResponse
type userListResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Users []string `json:"users"`
		} `json:"data"`
	} `json:"ocs"`
}

type userListDetailsResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Users map[string]UserDetails `json:"users"`
		} `json:"data"`
	} `json:"ocs"`
}

//userResponse
type userResponse struct {
	Ocs struct {
		Meta meta        `json:"meta"`
		Data UserDetails `json:"data"`
	} `json:"ocs"`
}

//baseResponse
type baseResponse struct {
	Ocs struct {
		Meta meta     `json:"meta"`
		Data []string `json:"data"`
	} `json:"ocs"`
}

//groupListResponse
type groupListResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Groups []string `json:"groups"`
		} `json:"data"`
	} `json:"ocs"`
}

//groupListDetailsResponse
type groupListDetailsResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Groups []Group `json:"groups"`
		} `json:"data"`
	} `json:"ocs"`
}

//appListResponse
type appListResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Apps []string `json:"apps"`
		} `json:"data"`
	} `json:"ocs"`
}

//appResponse
type appResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data App  `json:"data"`
	} `json:"ocs"`
}

type appConfigResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Data []string `json:"data"`
		} `json:"data"`
	} `json:"ocs"`
}

type appcConfigValueResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Data string `json:"data"`
		} `json:"data"`
	} `json:"ocs"`
}

//capabilitiesResponse
type capabilitiesResponse struct {
	Ocs struct {
		Meta meta `json:"meta"`
		Data struct {
			Version      Version      `json:"version"`
			Capabilities Capabilities `json:"capabilities"`
		} `json:"data"`
	} `json:"ocs"`
}

// Version contains the nextcloud version informations
type Version struct {
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Micro   int    `json:"micro"`
	String  string `json:"string"`
	Edition string `json:"edition"`
}

type monitoringResponse struct {
	Ocs struct {
		Meta meta       `json:"meta"`
		Data Monitoring `json:"data"`
	} `json:"ocs"`
}

type sharesListResponse struct {
	Ocs struct {
		Meta meta    `json:"meta"`
		Data []Share `json:"data"`
	} `json:"ocs"`
}

type sharesResponse struct {
	Ocs struct {
		Meta meta  `json:"meta"`
		Data Share `json:"data"`
	} `json:"ocs"`
}

type groupFoldersListResponse struct {
	Ocs struct {
		Meta meta                                       `json:"meta"`
		Data map[string]groupFolderBadFormatIDAndGroups `json:"data"`
	} `json:"ocs"`
}

type groupFoldersCreateResponse struct {
	Ocs struct {
		Meta meta                            `json:"meta"`
		Data groupFolderBadFormatIDAndGroups `json:"data"`
	} `json:"ocs"`
}

type groupFoldersResponse struct {
	Ocs struct {
		Meta meta                       `json:"meta"`
		Data groupFolderBadFormatGroups `json:"data"`
	} `json:"ocs"`
}

type notificationsListResponse struct {
	Ocs struct {
		Meta meta           `json:"meta"`
		Data []Notification `json:"data"`
	} `json:"ocs"`
}

type notificationResponse struct {
	Ocs struct {
		Meta meta         `json:"meta"`
		Data Notification `json:"data"`
	} `json:"ocs"`
}
