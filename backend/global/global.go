package global

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	// Readiness TODO: Moving to core/lifecycle brings import cycle :(
	Readiness sync.Map

	// Redis TODO: Shall we include global instance right here or as part of app (?!)
	// TODO: What if we need them before the app, such as middlewares...

	Redis *redis.Client
)
