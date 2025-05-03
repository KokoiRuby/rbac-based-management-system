package utils

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

func MigrateToRDB() {
	err := global.RDB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.UserRoleBinding{},
		&model.Menu{},
		&model.Api{},
		&model.RoleMenuBinding{},
		// Casbin
		&gormadapter.CasbinRule{},
	)
	if err != nil {
		zap.S().Fatalf("Failed to migrate RDB table: %v", err)
		return
	}
	close(global.MigrationDone)
	zap.S().Info("Migrated to RDB table successfully")
}
