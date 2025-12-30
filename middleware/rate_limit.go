package middleware

import (
	"encoding/json"
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

func (middleware *RateLimitMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userID := request.Header.Get("User-ID")
	if userID == "" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]interface{}{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"data":   "User-Id header requaired",
		})
		return
	}

}
