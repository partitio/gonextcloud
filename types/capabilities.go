package types

//Capabilities
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
