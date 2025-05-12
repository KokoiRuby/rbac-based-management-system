package middleware

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex

func LimitMiddleware(limit int) gin.HandlerFunc {
	return NewRateLimiter(limit, 5*time.Second).Middleware
}

type RateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	go rl.resetVisitorCount() // Reset window
	return rl
}

func (rl *RateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		rl.visitors = make(map[string]int) // Clear
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Middleware(c *gin.Context) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	ip := c.ClientIP()
	rl.visitors[ip]++
	if rl.visitors[ip] >= rl.limit {
		utils.FailWithMsg(c, http.StatusTooManyRequests, "Too many requests.")
		c.Abort()
		return
	}

	c.Next()
}
