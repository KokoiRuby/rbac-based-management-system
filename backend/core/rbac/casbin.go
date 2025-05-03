package rbac

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

func InitCasbin() (*casbin.CachedEnforcer, error) {
	// Wait for RDB ready
	<-global.RDBReady
	a, _ := gormadapter.NewAdapterByDB(global.RDB)
	m, err := model.NewModelFromFile("./core/rbac/rbac.pml")
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewCachedEnforcer(m, a)
	e.SetExpireTime(60 * 60)
	_ = e.LoadPolicy()

	zap.S().Info("Initialized casbin enforcer successfully")
	return e, err
}
