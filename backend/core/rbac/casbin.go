package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewCasbin(db *gorm.DB) *casbin.CachedEnforcer {
	a, _ := gormadapter.NewAdapterByDB(db)
	m, err := model.NewModelFromFile("./core/rbac/rbac.pml")
	if err != nil {
		zap.S().Fatalf("Failed to create model from file: %v", err)
	}

	e, err := casbin.NewCachedEnforcer(m, a)
	e.SetExpireTime(60 * 60)
	_ = e.LoadPolicy()

	zap.S().Info("Initialized casbin enforcer successfully")
	return e
}
