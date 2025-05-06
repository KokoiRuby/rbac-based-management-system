package main

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/bootstrap"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//gormGen()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}

	app := bootstrap.NewApp(ctx)

	signal.Notify(bootstrap.SigChan, syscall.SIGINT, syscall.SIGTERM)
	go app.Shutdown(cancel)

	go app.Start(ctx)
	zap.S().Infof("Server running on :%v...", app.RuntimeConfig.Gin.Port)

	select {
	case <-bootstrap.ShutDownChan:
		zap.S().Info("Graceful shutdown completed.")
		os.Exit(0)
	case <-time.After(5 * time.Minute):
		zap.S().Info("Program timeout, exiting in 3s")
		time.Sleep(3 * time.Second)
	}
}

//func gormGen() {
//	g := gen.NewGenerator(gen.Config{
//		OutPath:       "./repository/query",
//		ModelPkgPath:  "models",
//		Mode:          gen.WithDefaultQuery | gen.WithoutContext,
//		FieldNullable: true, // Null to pointer field
//	})
//
//	dia := mysql.Open("root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
//	session, _ := gorm.Open(dia, &gorm.Config{})
//	g.UseDB(session)
//	g.ApplyBasic(
//		&model.User{},
//		&model.Role{},
//		&model.Menu{},
//		&model.UserRoleBinding{},
//		&model.RoleMenuBinding{},
//		&model.Api{},
//	)
//	g.Execute()
//}
