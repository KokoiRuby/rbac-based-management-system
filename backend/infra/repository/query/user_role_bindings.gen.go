// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUserRoleBinding(db *gorm.DB, opts ...gen.DOOption) userRoleBinding {
	_userRoleBinding := userRoleBinding{}

	_userRoleBinding.userRoleBindingDo.UseDB(db, opts...)
	_userRoleBinding.userRoleBindingDo.UseModel(&model.UserRoleBinding{})

	tableName := _userRoleBinding.userRoleBindingDo.TableName()
	_userRoleBinding.ALL = field.NewAsterisk(tableName)
	_userRoleBinding.ID = field.NewUint(tableName, "id")
	_userRoleBinding.CreatedAt = field.NewTime(tableName, "created_at")
	_userRoleBinding.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userRoleBinding.UserID = field.NewUint(tableName, "user_id")
	_userRoleBinding.RoleID = field.NewUint(tableName, "role_id")
	_userRoleBinding.User = userRoleBindingBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
		RoleList: struct {
			field.RelationField
			UserList struct {
				field.RelationField
			}
			MenuList struct {
				field.RelationField
				ParentMenu struct {
					field.RelationField
				}
				Children struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("User.RoleList", "model.Role"),
			UserList: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.RoleList.UserList", "model.User"),
			},
			MenuList: struct {
				field.RelationField
				ParentMenu struct {
					field.RelationField
				}
				Children struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("User.RoleList.MenuList", "model.Menu"),
				ParentMenu: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("User.RoleList.MenuList.ParentMenu", "model.Menu"),
				},
				Children: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("User.RoleList.MenuList.Children", "model.Menu"),
				},
			},
		},
	}

	_userRoleBinding.Role = userRoleBindingBelongsToRole{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Role", "model.Role"),
	}

	_userRoleBinding.fillFieldMap()

	return _userRoleBinding
}

type userRoleBinding struct {
	userRoleBindingDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	UserID    field.Uint
	RoleID    field.Uint
	User      userRoleBindingBelongsToUser

	Role userRoleBindingBelongsToRole

	fieldMap map[string]field.Expr
}

func (u userRoleBinding) Table(newTableName string) *userRoleBinding {
	u.userRoleBindingDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userRoleBinding) As(alias string) *userRoleBinding {
	u.userRoleBindingDo.DO = *(u.userRoleBindingDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userRoleBinding) updateTableName(table string) *userRoleBinding {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.UserID = field.NewUint(table, "user_id")
	u.RoleID = field.NewUint(table, "role_id")

	u.fillFieldMap()

	return u
}

func (u *userRoleBinding) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userRoleBinding) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 7)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["role_id"] = u.RoleID

}

func (u userRoleBinding) clone(db *gorm.DB) userRoleBinding {
	u.userRoleBindingDo.ReplaceConnPool(db.Statement.ConnPool)
	u.User.db = db.Session(&gorm.Session{Initialized: true})
	u.User.db.Statement.ConnPool = db.Statement.ConnPool
	u.Role.db = db.Session(&gorm.Session{Initialized: true})
	u.Role.db.Statement.ConnPool = db.Statement.ConnPool
	return u
}

func (u userRoleBinding) replaceDB(db *gorm.DB) userRoleBinding {
	u.userRoleBindingDo.ReplaceDB(db)
	u.User.db = db.Session(&gorm.Session{})
	u.Role.db = db.Session(&gorm.Session{})
	return u
}

type userRoleBindingBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	RoleList struct {
		field.RelationField
		UserList struct {
			field.RelationField
		}
		MenuList struct {
			field.RelationField
			ParentMenu struct {
				field.RelationField
			}
			Children struct {
				field.RelationField
			}
		}
	}
}

func (a userRoleBindingBelongsToUser) Where(conds ...field.Expr) *userRoleBindingBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userRoleBindingBelongsToUser) WithContext(ctx context.Context) *userRoleBindingBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userRoleBindingBelongsToUser) Session(session *gorm.Session) *userRoleBindingBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a userRoleBindingBelongsToUser) Model(m *model.UserRoleBinding) *userRoleBindingBelongsToUserTx {
	return &userRoleBindingBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a userRoleBindingBelongsToUser) Unscoped() *userRoleBindingBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type userRoleBindingBelongsToUserTx struct{ tx *gorm.Association }

func (a userRoleBindingBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a userRoleBindingBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userRoleBindingBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userRoleBindingBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userRoleBindingBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a userRoleBindingBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a userRoleBindingBelongsToUserTx) Unscoped() *userRoleBindingBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type userRoleBindingBelongsToRole struct {
	db *gorm.DB

	field.RelationField
}

func (a userRoleBindingBelongsToRole) Where(conds ...field.Expr) *userRoleBindingBelongsToRole {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userRoleBindingBelongsToRole) WithContext(ctx context.Context) *userRoleBindingBelongsToRole {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userRoleBindingBelongsToRole) Session(session *gorm.Session) *userRoleBindingBelongsToRole {
	a.db = a.db.Session(session)
	return &a
}

func (a userRoleBindingBelongsToRole) Model(m *model.UserRoleBinding) *userRoleBindingBelongsToRoleTx {
	return &userRoleBindingBelongsToRoleTx{a.db.Model(m).Association(a.Name())}
}

func (a userRoleBindingBelongsToRole) Unscoped() *userRoleBindingBelongsToRole {
	a.db = a.db.Unscoped()
	return &a
}

type userRoleBindingBelongsToRoleTx struct{ tx *gorm.Association }

func (a userRoleBindingBelongsToRoleTx) Find() (result *model.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a userRoleBindingBelongsToRoleTx) Append(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userRoleBindingBelongsToRoleTx) Replace(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userRoleBindingBelongsToRoleTx) Delete(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userRoleBindingBelongsToRoleTx) Clear() error {
	return a.tx.Clear()
}

func (a userRoleBindingBelongsToRoleTx) Count() int64 {
	return a.tx.Count()
}

func (a userRoleBindingBelongsToRoleTx) Unscoped() *userRoleBindingBelongsToRoleTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type userRoleBindingDo struct{ gen.DO }

func (u userRoleBindingDo) Debug() *userRoleBindingDo {
	return u.withDO(u.DO.Debug())
}

func (u userRoleBindingDo) WithContext(ctx context.Context) *userRoleBindingDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userRoleBindingDo) ReadDB() *userRoleBindingDo {
	return u.Clauses(dbresolver.Read)
}

func (u userRoleBindingDo) WriteDB() *userRoleBindingDo {
	return u.Clauses(dbresolver.Write)
}

func (u userRoleBindingDo) Session(config *gorm.Session) *userRoleBindingDo {
	return u.withDO(u.DO.Session(config))
}

func (u userRoleBindingDo) Clauses(conds ...clause.Expression) *userRoleBindingDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userRoleBindingDo) Returning(value interface{}, columns ...string) *userRoleBindingDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userRoleBindingDo) Not(conds ...gen.Condition) *userRoleBindingDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userRoleBindingDo) Or(conds ...gen.Condition) *userRoleBindingDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userRoleBindingDo) Select(conds ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userRoleBindingDo) Where(conds ...gen.Condition) *userRoleBindingDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userRoleBindingDo) Order(conds ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userRoleBindingDo) Distinct(cols ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userRoleBindingDo) Omit(cols ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userRoleBindingDo) Join(table schema.Tabler, on ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userRoleBindingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userRoleBindingDo) RightJoin(table schema.Tabler, on ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userRoleBindingDo) Group(cols ...field.Expr) *userRoleBindingDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userRoleBindingDo) Having(conds ...gen.Condition) *userRoleBindingDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userRoleBindingDo) Limit(limit int) *userRoleBindingDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userRoleBindingDo) Offset(offset int) *userRoleBindingDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userRoleBindingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userRoleBindingDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userRoleBindingDo) Unscoped() *userRoleBindingDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userRoleBindingDo) Create(values ...*model.UserRoleBinding) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userRoleBindingDo) CreateInBatches(values []*model.UserRoleBinding, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userRoleBindingDo) Save(values ...*model.UserRoleBinding) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userRoleBindingDo) First() (*model.UserRoleBinding, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleBinding), nil
	}
}

func (u userRoleBindingDo) Take() (*model.UserRoleBinding, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleBinding), nil
	}
}

func (u userRoleBindingDo) Last() (*model.UserRoleBinding, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleBinding), nil
	}
}

func (u userRoleBindingDo) Find() ([]*model.UserRoleBinding, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserRoleBinding), err
}

func (u userRoleBindingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRoleBinding, err error) {
	buf := make([]*model.UserRoleBinding, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userRoleBindingDo) FindInBatches(result *[]*model.UserRoleBinding, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userRoleBindingDo) Attrs(attrs ...field.AssignExpr) *userRoleBindingDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userRoleBindingDo) Assign(attrs ...field.AssignExpr) *userRoleBindingDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userRoleBindingDo) Joins(fields ...field.RelationField) *userRoleBindingDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userRoleBindingDo) Preload(fields ...field.RelationField) *userRoleBindingDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userRoleBindingDo) FirstOrInit() (*model.UserRoleBinding, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleBinding), nil
	}
}

func (u userRoleBindingDo) FirstOrCreate() (*model.UserRoleBinding, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleBinding), nil
	}
}

func (u userRoleBindingDo) FindByPage(offset int, limit int) (result []*model.UserRoleBinding, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userRoleBindingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userRoleBindingDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userRoleBindingDo) Delete(models ...*model.UserRoleBinding) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userRoleBindingDo) withDO(do gen.Dao) *userRoleBindingDo {
	u.DO = *do.(*gen.DO)
	return u
}
