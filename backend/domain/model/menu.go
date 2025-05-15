package model

type Meta struct {
	Icon  string `gorm:"size:256" json:"icon"`
	Title string `gorm:"size:32" json:"title"`
}

type Menu struct {
	Model
	Name         string `gorm:"size:32,unique" json:"name"`
	Path         string `gorm:"size:128" json:"path"`
	Component    string `gorm:"size:128" json:"component"`
	ParentMenuID *uint  `json:"parentMenuId"` // pointer = NULL-able
	ParentMenu   *Menu  `gorm:"foreignKey:ParentMenuID" json:"-"`
	Children     []Menu `gorm:"foreignKey:ParentMenuID" json:"children"`
	Sort         int    `json:"sort"`
	Enable       bool   `json:"enable"`

	Meta `gorm:"embedded" json:"meta"`

	// One menu can have many roles
	//RoleList []Menu `gorm:"many2many:role_menu_bindings;joinForeignKey:MenuID;joinReferences:RoleID" json:"roleList"`
}
