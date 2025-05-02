package global

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

const VERSION = "0.0.1"

type Ready bool

var (
	RuntimeConfig *config.RuntimeConfig
	RDB           *gorm.DB
	Redis         *redis.Client
	Mongo         *mongo.Client
	Readiness     = map[string]Ready{
		"redis": false,
		"mongo": false,
		"rdb":   false,
	}
)
