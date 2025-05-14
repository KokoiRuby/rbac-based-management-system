package model

type UserListRequest struct {
	Pagination
	Username string `form:"username"` // Fuzzy matching
	Email    string `form:"email"`
	Role     uint   `form:"role"`
}

type UserListResponse struct {
	User
}
