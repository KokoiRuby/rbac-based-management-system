package model

type UserRoleBinding struct {
	Model
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`

	// One user can have many user role bindings
	// https://gorm.io/docs/has_many.html
	User User `gorm:"foreignKey:UserID" json:"-"`

	// One role can have many user role bindings
	// https://gorm.io/docs/has_many.html
	Role Role `gorm:"foreignKey:RoleID" json:"-"`
}
