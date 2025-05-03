package model

import (
	"time"
)

// GORM model
//type Model struct {
//	ID        uint `gorm:"primarykey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt DeletedAt `gorm:"index"`
//}

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// PageInfo TODO: query param > form
type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Order string `form:"order"`
}
