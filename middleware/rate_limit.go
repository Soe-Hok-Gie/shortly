package middleware

import (
	"net/http"
	"sync"
)

type RateLimitMiddleware struct {
	Handler      http.Handler
	RateLimitMap map[string]*RateLimitData
	Mutex        *sync.Mutex
}
