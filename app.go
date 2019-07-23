package gonextcloud

// App is a nextcloud application (plugin)
type App struct {
	ID            string   `json:"id"`
	Ocsid         string   `json:"ocsid"`
	Name          string   `json:"name"`
	Summary       string   `json:"summary"`
	Description   string   `json:"description"`
	Licence       string   `json:"licence"`
	Author        string   `json:"author"`
	Version       string   `json:"version"`
	Namespace     string   `json:"namespace"`
	Types         []string `json:"types"`
	Documentation struct {
		Admin     string `json:"admin"`
		Developer string `json:"developer"`
		User      string `json:"user"`
	} `json:"documentation"`
	Category   []string `json:"category"`
	Website    string   `json:"website"`
	Bugs       string   `json:"bugs"`
	Repository struct {
		Attributes struct {
			Type string `json:"type"`
		} `json:"@attributes"`
		Value string `json:"@value"`
	} `json:"repository"`
	Screenshot   []interface{} `json:"screenshot"`
	Dependencies struct {
		Owncloud struct {
			Attributes struct {
				MinVersion string `json:"min-version"`
				MaxVersion string `json:"max-version"`
			} `json:"@attributes"`
		} `json:"owncloud"`
		Nextcloud struct {
			Attributes struct {
				MinVersion string `json:"min-version"`
				MaxVersion string `json:"max-version"`
			} `json:"@attributes"`
		} `json:"nextcloud"`
	} `json:"dependencies"`
	Settings struct {
		Admin           []string      `json:"admin"`
		AdminSection    []string      `json:"admin-section"`
		Personal        []interface{} `json:"personal"`
		PersonalSection []interface{} `json:"personal-section"`
	} `json:"settings"`
	Info        []interface{} `json:"info"`
	Remote      []interface{} `json:"remote"`
	Public      []interface{} `json:"public"`
	RepairSteps struct {
		Install       []interface{} `json:"install"`
		PreMigration  []interface{} `json:"pre-migration"`
		PostMigration []interface{} `json:"post-migration"`
		LiveMigration []interface{} `json:"live-migration"`
		Uninstall     []interface{} `json:"uninstall"`
	} `json:"repair-steps"`
	BackgroundJobs     []interface{} `json:"background-jobs"`
	TwoFactorProviders []interface{} `json:"two-factor-providers"`
	Commands           []interface{} `json:"commands"`
	Activity           struct {
		Filters   []interface{} `json:"filters"`
		Settings  []interface{} `json:"settings"`
		Providers []interface{} `json:"providers"`
	} `json:"activity"`
}
