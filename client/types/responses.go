package types

type ErrorResponse struct {
	Ocs struct {
		Meta struct {
			Status       string `json:"status"`
			Statuscode   int    `json:"statuscode"`
			Message      string `json:"message"`
			Totalitems   string `json:"totalitems"`
			Itemsperpage string `json:"itemsperpage"`
		} `json:"meta"`
		Data []interface{} `json:"data"`
	} `json:"ocs"`
}

type UserListResponse struct {
	Ocs struct {
		Meta struct {
			Status       string `json:"status"`
			Statuscode   int    `json:"statuscode"`
			Message      string `json:"message"`
			Totalitems   string `json:"totalitems"`
			Itemsperpage string `json:"itemsperpage"`
		} `json:"meta"`
		Data struct {
			Users []string `json:"users"`
		} `json:"data"`
	} `json:"ocs"`
}

type UserResponse struct {
	Ocs struct {
		Meta struct {
			Status       string `json:"status"`
			Statuscode   int    `json:"statuscode"`
			Message      string `json:"message"`
			Totalitems   string `json:"totalitems"`
			Itemsperpage string `json:"itemsperpage"`
		} `json:"meta"`
		Data User `json:"data"`
	} `json:"ocs"`
}

type BaseResponse struct {
	Ocs struct {
		Meta struct {
			Status       string `json:"status"`
			Statuscode   int    `json:"statuscode"`
			Message      string `json:"message"`
			Totalitems   string `json:"totalitems"`
			Itemsperpage string `json:"itemsperpage"`
		} `json:"meta"`
		Data []string `json:"data"`
	} `json:"ocs"`
}

type GroupListResponse struct {
	Ocs struct {
		Meta struct {
			Status       string `json:"status"`
			Statuscode   int    `json:"statuscode"`
			Message      string `json:"message"`
			Totalitems   string `json:"totalitems"`
			Itemsperpage string `json:"itemsperpage"`
		} `json:"meta"`
		Data struct {
			Groups []string `json:"groups"`
		} `json:"data"`
	} `json:"ocs"`
}

type CapabilitiesResponse struct {
	Ocs struct {
		Meta struct {
			Status       string `json:"status"`
			Statuscode   int    `json:"statuscode"`
			Message      string `json:"message"`
			Totalitems   string `json:"totalitems"`
			Itemsperpage string `json:"itemsperpage"`
		} `json:"meta"`
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
