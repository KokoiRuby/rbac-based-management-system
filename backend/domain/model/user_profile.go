package model

type UserProfile struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	RoleList []uint `json:"roleList"`
}
