package repository

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository/query"
	"gorm.io/gen"
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

func (u userRDB) GetByID(c context.Context, id uint) (model.User, error) {
	query.SetDefault(u.rdb)
	user, err := query.User.WithContext(c).Where(query.User.ID.Eq(id)).Take()
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}

func (u userRDB) GetByEmail(c context.Context, email string) (model.User, error) {
	query.SetDefault(u.rdb)
	user, err := query.User.WithContext(c).Where(query.User.Email.Eq(email)).Take()
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}

func (u userRDB) Update(c context.Context, user *model.User) error {
	query.SetDefault(u.rdb)
	_, err := query.User.WithContext(c).Where(query.User.Email.Eq(user.Email)).Updates(user)
	if err != nil {
		return err
	}
	return nil
}

func (u userRDB) GetByCond(c context.Context, conds ...gen.Condition) (model.User, bool, error) {
	query.SetDefault(u.rdb)
	user, err := query.User.WithContext(c).Where(conds...).Take()
	if err != nil {
		return model.User{}, false, err
	}

	return *user, true, nil
}
