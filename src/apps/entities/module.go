package entities

import "time"

type Modules struct {
	ModuleId   string    `json:"module_id"`
	AppId      string    `json:"app_id"`
	ModuleName string    `json:"module_name"`
	CreatedAt  time.Time `json:"-"`
	Created_by string    `json:"created_by,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	UpdatedBy  string    `json:"updated_by,omitempty"`
}

func (Modules) TableName() string {
	return "tbm_module"
}

type ModuleRequest struct {
	ModuleName string `json:"module_name"`
}
