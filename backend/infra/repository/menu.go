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

type menuRDB struct {
	rdb *gorm.DB
}

func NewMenuRepository(rdb *gorm.DB) service.MenuRDB {
	return &menuRDB{
		rdb: rdb,
	}
}

func (m menuRDB) Create(c context.Context, menu *model.Menu) error {
	query.SetDefault(m.rdb)
	err := query.Menu.Create(menu)
	if err != nil {
		return err
	}
	return nil
}

func (m menuRDB) GetByCond(c context.Context, conds ...gen.Condition) (*model.Menu, error) {
	query.SetDefault(m.rdb)
	user, err := query.Menu.Debug().WithContext(c).Preload(query.Menu.Children).Where(conds...).Take()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m menuRDB) GetByID(c context.Context, id uint) (*model.Menu, error) {
	query.SetDefault(m.rdb)
	cond := query.Menu.ID.Eq(id)
	return m.GetByCond(c, cond)
}

func (m menuRDB) GetByName(c context.Context, name string) (*model.Menu, error) {
	query.SetDefault(m.rdb)
	cond := query.Menu.Name.Eq(name)
	return m.GetByCond(c, cond)
}

func (m menuRDB) ListByCond(c context.Context, opt model.QueryOptions) ([]*model.Menu, int64, error) {
	// Pagination
	page := opt.Page
	limit := opt.Limit
	offset := (page - 1) * limit

	// Order
	var order field.Expr
	if opt.Order {
		order = query.Menu.CreatedAt.Desc()
	} else {
		order = query.Menu.CreatedAt.Asc()
	}

	queryBuilder := query.Menu.Where(
		// Fuzzy AND
		query.Menu.Name.Like(fmt.Sprintf("%%%s%%", opt.Likes["name"].(string))),
		query.Menu.MetaTitle.Like(fmt.Sprintf("%%%s%%", opt.Likes["title"].(string))),
	)

	query.SetDefault(m.rdb)
	users, err := queryBuilder.
		Debug().
		WithContext(c).
		Preload(query.Menu.Children).
		Where(query.Menu.ParentMenuID.IsNull()).
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
