package global

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"gorm.io/gorm"
	"sync"
)

const VERSION = "0.0.1"

var (
	RuntimeConfig *config.RuntimeConfig

	RDB *gorm.DB

	// Readiness
	// map[string]bool is not inherently thread-safe for concurrent access from multiple goroutines.
	// Even on different keys. Replaced with sync.Map

	// Readiness TODO: Moving to core/lifecycle brings import cycle :(
	Readiness sync.Map
)
