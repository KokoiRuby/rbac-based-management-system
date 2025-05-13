package model

// TODO: Shall we allow updating fields with unique index?!

type UserUpdate struct {
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
}
