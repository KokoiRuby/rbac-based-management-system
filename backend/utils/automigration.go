package utils

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func MigrateToRDB(db *gorm.DB) {
	err := db.AutoMigrate(
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
	zap.S().Info("Migrated to RDB table successfully")
}
