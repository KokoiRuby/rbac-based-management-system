package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
)

type UserRDB interface {
	Create(c context.Context, user *model.User) error
	List(c context.Context) ([]model.User, error)
	GetByID(c context.Context, id string) (model.User, error)
	GetByEmail(c context.Context, email string) (model.User, error)
}
