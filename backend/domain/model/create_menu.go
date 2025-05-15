package model

type CreateMenuRequest struct {
	Name         string `json:"name" binding:"required"`
	Path         string `json:"path" binding:"required"`
	Component    string `json:"component"`
	ParentMenuID *uint  `json:"parentMenuId"`
	Sort         int    `json:"sort"`
	Enable       bool   `json:"enable"`
	Icon         string `json:"icon"`
	Title        string `json:"title"`
}
