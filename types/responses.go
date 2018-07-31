package types

//Meta
type Meta struct {
	Status       string `json:"status"`
	Statuscode   int    `json:"statuscode"`
	Message      string `json:"message"`
	Totalitems   string `json:"totalitems"`
	Itemsperpage string `json:"itemsperpage"`
}

//ErrorResponse
type ErrorResponse struct {
	Ocs struct {
		Meta Meta          `json:"meta"`
		Data []interface{} `json:"data"`
	} `json:"ocs"`
}

//UserListResponse
type UserListResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Users []string `json:"users"`
		} `json:"data"`
	} `json:"ocs"`
}

//UserResponse
type UserResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data User `json:"data"`
	} `json:"ocs"`
}

//BaseResponse
type BaseResponse struct {
	Ocs struct {
		Meta Meta     `json:"meta"`
		Data []string `json:"data"`
	} `json:"ocs"`
}

//GroupListResponse
type GroupListResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Groups []string `json:"groups"`
		} `json:"data"`
	} `json:"ocs"`
}

//AppListResponse
type AppListResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Apps []string `json:"apps"`
		} `json:"data"`
	} `json:"ocs"`
}

//AppResponse
type AppResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data App  `json:"data"`
	} `json:"ocs"`
}

//CapabilitiesResponse
type CapabilitiesResponse struct {
	Ocs struct {
		Meta Meta `json:"meta"`
		Data struct {
			Version struct {
				Major   int    `json:"major"`
				Minor   int    `json:"minor"`
				Micro   int    `json:"micro"`
				String  string `json:"string"`
				Edition string `json:"edition"`
			} `json:"version"`
			Capabilities Capabilities `json:"capabilities"`
		} `json:"data"`
	} `json:"ocs"`
}

type MonitoringResponse struct {
	Ocs struct {
		Meta struct {
			Status     string `json:"status"`
			Statuscode int    `json:"statuscode"`
			Message    string `json:"message"`
		} `json:"meta"`
		Data Monitoring `json:"data"`
	} `json:"ocs"`
}
