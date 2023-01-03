package apps

type Apps struct {
	AppId          string `json:"app_id"`
	AppName        string `json:"app_name"`
	AppUrl         string `json:"app_url"`
	AppUrlCallback string `json:"app_url_callback"`
	AppDesc        string `json:"app_desc"`
	SecretCode     string `json:"secret_code"`
	CreatedAt      string `json:"created_at"`
	Created_by     string `json:"created_by"`
	UpdatedAt      string `json:"updated_at"`
	UpdatedBy      string `json:"updated_by"`
}

type Get struct {
	AppId          string `json:"app_id"`
	AppName        string `json:"app_name"`
	AppUrl         string `json:"app_url"`
	AppUrlCallback string `json:"app_url_callback"`
	AppDesc        string `json:"app_desc"`
	SecretCode     string `json:"secret_code"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Apps) TableName() string {
	return "tbm_apps"
}
