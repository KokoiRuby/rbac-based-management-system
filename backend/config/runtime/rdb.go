package runtime

import (
	"fmt"
	"github.com/glebarez/sqlite" // https://gorm.io/docs/connecting_to_the_database.html#SQLite
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBMode string

const (
	MYSQL  DBMode = "mysql"
	PG     DBMode = "postgres"
	SQLITE DBMode = "sqlite"
)

type RDBConfig struct {
	Mode         DBMode `yaml:"mode"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user"`
	Pass         string `yaml:"pass"`
	DBName       string `yaml:"dbName"`
	PingInterval int    `yaml:"pingInterval"`
}

func (db RDBConfig) GetDSN() gorm.Dialector {
	switch db.Mode {
	case MYSQL:
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			db.User,
			db.Pass,
			db.Host,
			db.Port,
			db.DBName,
		)
		return mysql.Open(dsn)
	case PG:
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
			db.User,
			db.Pass,
			db.Host,
			db.Port,
			db.DBName,
		)
		return postgres.Open(dsn)
	case SQLITE:
		return sqlite.Open(db.DBName)
	case "":
		zap.S().Warn("Database mode not specified")
		return nil
	default:
		zap.S().Fatal("Database is not supported")
		return nil
	}
}
