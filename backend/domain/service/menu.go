package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"gorm.io/gen"
)

type MenuRDB interface {
	Create(c context.Context, menu *model.Menu) error
	GetByCond(c context.Context, conds ...gen.Condition) (*model.Menu, error)
	GetByID(c context.Context, id uint) (*model.Menu, error)
	GetByName(c context.Context, name string) (*model.Menu, error)
	ListByCond(c context.Context, opt model.QueryOptions) ([]*model.Menu, int64, error)
}
