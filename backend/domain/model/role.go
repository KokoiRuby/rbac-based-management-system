package model

type Role struct {
	Model
	Name string `gorm:"size:64;uniqueIndex" json:"name"`

	// One role can have many users
	// https://gorm.io/docs/many_to_many.html
	UserList []User `gorm:"many2many:user_role_bindings;joinForeignKey:RoleID;joinReferences:UserID" json:"userList"`

	// One role can have many menus
	MenuList []Menu `gorm:"many2many:role_menu_bindings;joinForeignKey:RoleID;joinReferences:MenuID" json:"menuList"`
}
