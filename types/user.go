package types

//User
type User struct {
	Enabled bool   `json:"enabled"`
	ID      string `json:"id"`
	Quota   struct {
		Free     int64   `json:"free"`
		Used     int     `json:"used"`
		Total    int64   `json:"total"`
		Relative float64 `json:"relative"`
		Quota    int     `json:"quota"`
	} `json:"quota"`
	Email       string   `json:"email"`
	Displayname string   `json:"displayname"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	Website     string   `json:"website"`
	Twitter     string   `json:"twitter"`
	Groups      []string `json:"groups"`
	Language    string   `json:"language"`
}
