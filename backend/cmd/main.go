package main

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/database"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
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

	err = config.ParseAndLoad()
	if err != nil {
		panic(err)
	}

	global.RDB, err = database.NewGormDB(global.RuntimeConfig.RDB)
	if err != nil {
		panic(err)
	}

	global.Redis, err = database.NewRedisClient(global.RuntimeConfig.Redis)
	if err != nil {
		panic(err)
	}

	global.Mongo, err = database.NewMongoClient(global.RuntimeConfig.Mongo)
	if err != nil {
		panic(err)
	}

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
