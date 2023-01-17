package entities

import "time"

type Apps struct {
	AppId          string    `json:"app_id"`
	AppName        string    `json:"app_name"`
	AppUrl         string    `json:"app_url"`
	AppUrlCallback string    `json:"app_url_callback"`
	AppDesc        string    `json:"app_desc"`
	SecretCode     string    `json:"secret_code"`
	CreatedAt      time.Time `json:"-"`
	Created_by     string    `json:"created_by,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	UpdatedBy      string    `json:"updated_by,omitempty"`
}

type Tabler interface {
	TableName() string
}

func (Apps) TableName() string {
	return "tbm_apps"
}

type AppRequest struct {
	AppName        string          `json:"app_name"`
	AppUrl         string          `json:"app_url"`
	AppUrlCallback string          `json:"app_url_callback"`
	AppDesc        string          `json:"app_desc"`
	SecretCode     string          `json:"secret_code"`
	Modules        []ModuleRequest `json:"modules"`
}

type CreateApp struct {
	App    Apps
	Module []Modules
}
type UpdateApp struct {
	Apps
	Modules       []Modules
	DeleteModules []Modules
}

type SendApps struct {
	Apps
	Module []Modules `json:"modules"`
}
