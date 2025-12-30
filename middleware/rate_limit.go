package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimitMiddleware struct {
	Handler      http.Handler
	RateLimitMap map[string]*RateLimitData
	Mutex        *sync.Mutex
}

type RateLimitData struct {
	HitCount int
	FirstHit time.Time
}
