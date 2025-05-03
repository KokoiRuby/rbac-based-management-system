package model

type Api struct {
	Model
	Name   string `gorm:"size:64;uniqueIndex" json:"name"` // API unique name/identifier
	Path   string `gorm:"size:256" json:"path"`            // API URL path
	Method string `gorm:"size:32" json:"method"`           // HTTP Method (GET, POST, etc.)
	Group  string `gorm:"size:128" json:"group"`           // API Group/Category
}
