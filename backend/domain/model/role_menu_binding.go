package model

type RoleMenuBinding struct {
	Model
	RoleID uint `json:"role_id"`
	MenuID uint `json:"menu_id"`

	Role Role `gorm:"foreignKey:RoleID" json:"-"`
	Menu Menu `gorm:"foreignKey:MenuID" json:"-"`
}
