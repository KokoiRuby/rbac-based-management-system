package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"gorm.io/gen"
)

type UserRDB interface {
	Create(c context.Context, user *model.User) error
	List(c context.Context) ([]model.User, error)
	GetByID(c context.Context, id uint) (model.User, error)
	GetByEmail(c context.Context, email string) (model.User, error)
	Update(c context.Context, user *model.User) error
	GetByCond(c context.Context, conds ...gen.Condition) (model.User, bool, error)
}
