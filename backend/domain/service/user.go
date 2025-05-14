package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"gorm.io/gen"
)

type UserRDB interface {
	Create(c context.Context, user *model.User) error
	ListByCond(c context.Context, opt model.QueryOptions) ([]*model.User, int64, error)
	GetByID(c context.Context, id uint) (*model.User, error)
	GetByEmail(c context.Context, email string) (*model.User, error)
	GetByCond(c context.Context, conds ...gen.Condition) (*model.User, error)
	Update(c context.Context, user *model.User) error
	DeleteByCond(c context.Context, conds ...gen.Condition) error
	DeleteByID(c context.Context, id uint) error
}
