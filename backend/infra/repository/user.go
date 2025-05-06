package repository

import (
	"context"
	"errors"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository/query"
	"gorm.io/gorm"
)

type userRDB struct {
	rdb *gorm.DB
}

func NewUserRepository(rdb *gorm.DB) service.UserRDB {
	return &userRDB{
		rdb: rdb,
	}
}

func (u userRDB) Create(c context.Context, user *model.User) error {
	query.SetDefault(u.rdb)
	err := query.User.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u userRDB) List(c context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRDB) GetByID(c context.Context, id string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRDB) GetByEmail(c context.Context, email string) (model.User, error) {
	query.SetDefault(u.rdb)
	user, err := query.User.Where(query.User.Email.Eq(email)).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, err
		}
		return model.User{}, err
	}
	return *user, nil
}
