package model

type User struct {
	Model
	Username string `gorm:"size:64;uniqueIndex" json:"username"` // Username often needs to be unique
	Password string `gorm:"size:128" json:"-"`                   // Store password hash
	Nickname string `gorm:"size:256" json:"nickname,omitempty"`
	Avatar   string `gorm:"size:256" json:"avatar,omitempty"`
	Email    string `gorm:"size:128" json:"email,omitempty"` // Assuming email is optional
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`

	// One user can have many roles
	// https://gorm.io/docs/many_to_many.html
	RoleList []Role `gorm:"many2many:user_role_bindings;joinForeignKey:UserID;joinReferences:RoleID" json:"role_list"`
}
