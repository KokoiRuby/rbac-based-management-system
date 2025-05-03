package lifecycle

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
)

func IsReady() bool {
	ready := true
	global.Readiness.Range(func(key, value interface{}) bool {
		isReady, ok := value.(bool) // Assert the type of the value
		if !ok || !isReady {
			ready = false
			return false // Stop the iteration early if any value is false
		}
		return true // Continue iterating
	})
	return ready
}
