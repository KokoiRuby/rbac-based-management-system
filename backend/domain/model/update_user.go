package model

// TODO: Shall we allow updating fields with unique index?!

type UserUpdate struct {
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
}

type UserUpdateConfirmRequest struct {
	UserID   uint
	Username string
	Nickname string
	Email    string
}
