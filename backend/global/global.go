package global

import (
	"sync"
)

var (
	// Readiness TODO: Moving to core/lifecycle brings import cycle :(
	Readiness sync.Map
)
