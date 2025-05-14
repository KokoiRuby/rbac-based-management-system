package model

import "gorm.io/gen"

type Pagination struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"` // Fuzzing
	Order bool   `form:"order"`
}

type QueryOptions struct {
	Pagination
	Likes    map[string]any
	Conds    []gen.Condition
	Preloads []string
}
