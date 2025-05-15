package model

type MenuListRequest struct {
	Pagination
	Name  string `form:"name"` // Fuzzy matching
	Title string `form:"title"`
}

type MenuListResponse struct {
	Menu
}
