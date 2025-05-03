package global

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
)

const VERSION = "0.0.1"

type Ready bool

var (
	RuntimeConfig *config.RuntimeConfig

	RDB    *gorm.DB
	Redis  *redis.Client
	Mongo  *mongo.Client
	Casbin *casbin.CachedEnforcer

	// Readiness
	// map is not inherently thread-safe for concurrent access from multiple goroutines.
	// Even on different keys. Replaced with sync.Map
	//Readiness = map[string]Ready{
	//	"redis": false,
	//	"mongo": false,
	//	"rdb":   false,
	//}
	Readiness sync.Map

	// RDBReady
	// Channel to signal once RDB is done
	RDBReady = make(chan struct{})

	// MigrationDone
	// Channel to signal once RDB is ready AND flag is enabled
	MigrationDone = make(chan struct{})
)
