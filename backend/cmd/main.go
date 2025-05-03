package main

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/database"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/rbac"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"go.uber.org/zap"
	"time"
)

func main() {

	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}

	err = config.ParseLoadRun()
	if err != nil {
		panic(err)
	}

	go func() {
		global.RDB, err = database.NewGormDB(global.RuntimeConfig.RDB)
		if err != nil {
			panic(err)
		}
		close(global.RDBReady) // Signal when instance is set
	}()

	go func() {
		global.Redis, err = database.NewRedisClient(global.RuntimeConfig.Redis)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		global.Mongo, err = database.NewMongoClient(global.RuntimeConfig.Mongo)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		global.Casbin, err = rbac.InitCasbin()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			if utils.IsReady() {
				zap.S().Debug("Ready")
			} else {
				zap.S().Debug("Not Ready")
			}
			time.Sleep(3 * time.Second)
		}
	}()

	select {}

}
