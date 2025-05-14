package repository

import (
	"context"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository/query"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// TODO: GORM Gen query instead of gorm.DB

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

func (u userRDB) ListByCond(c context.Context, opt model.QueryOptions) ([]*model.User, int64, error) {

	// Pagination
	page := opt.Page
	limit := opt.Limit
	offset := (page - 1) * limit

	// Order
	var order field.Expr
	if opt.Order {
		order = query.User.CreatedAt.Desc()
	} else {
		order = query.User.CreatedAt.Asc()
	}

	queryBuilder := query.User.Where(
		// Fuzzy AND
		query.User.Username.Like(fmt.Sprintf("%%%s%%", opt.Likes["username"].(string))),
		query.User.Email.Like(fmt.Sprintf("%%%s%%", opt.Likes["email"].(string))),
	)

	query.SetDefault(u.rdb)
	users, err := queryBuilder.
		Debug().
		Order(order).
		Offset(offset).
		Limit(limit).
		Find()
	if err != nil {
		return nil, 0, err
	}

	cnt, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}

	return users, cnt, nil
}

func (u userRDB) GetByCond(c context.Context, conds ...gen.Condition) (*model.User, error) {
	query.SetDefault(u.rdb)
	user, err := query.User.Debug().WithContext(c).Preload(query.User.RoleList).Where(conds...).Take()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRDB) GetByID(c context.Context, id uint) (*model.User, error) {
	query.SetDefault(u.rdb)
	cond := query.User.ID.Eq(id)
	return u.GetByCond(c, cond)
}

func (u userRDB) GetByEmail(c context.Context, email string) (*model.User, error) {
	query.SetDefault(u.rdb)
	cond := query.User.Email.Eq(email)
	return u.GetByCond(c, cond)
}

func (u userRDB) Update(c context.Context, user *model.User) error {
	query.SetDefault(u.rdb)
	_, err := query.User.WithContext(c).Where(query.User.Email.Eq(user.Email)).Updates(user)
	if err != nil {
		return err
	}
	return nil
}

func (u userRDB) DeleteByCond(c context.Context, conds ...gen.Condition) error {
	query.SetDefault(u.rdb)
	_, err := query.User.Where(conds...).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u userRDB) DeleteByID(c context.Context, id uint) error {
	cond := query.User.ID.Eq(id)
	return u.DeleteByCond(c, cond)
}
