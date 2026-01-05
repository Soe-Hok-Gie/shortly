package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimitMiddleware struct {
	RateLimitMap map[string]*RateLimitData
	Mutex        sync.Mutex
}

type RateLimitData struct {
	HitCount int
	FirstHit time.Time
	Action   string
}

func NewRateLimitMiddleware() *RateLimitMiddleware {
	return &RateLimitMiddleware{
		RateLimitMap: make(map[string]*RateLimitData),
		Mutex:        sync.Mutex{},
	}
}

func (middleware *RateLimitMiddleware) WithRateLimit() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			userID := request.Header.Get("User-ID")
			if userID == "" {
				userID = request.RemoteAddr
			}

			// Lock sebelum akses map
			// middleware.Mutex.Lock()
			// defer middleware.Mutex.Unlock() // 🔑 Pakai defer untuk safety

			// now := time.Now()
			// data, exists := middleware.RateLimitMap[userID]

			// if !exists || (now.Sub(data.FirstHit) > time.Minute && data.Action == request.URL.Path) {
			// 	//window baru
			// 	middleware.RateLimitMap[userID] = &RateLimitData{
			// 		HitCount: 1,
			// 		FirstHit: now,
			// 		Action:   request.URL.Path,
			// 	}
			// } else {
			// 	data.HitCount++
			// 	if data.HitCount > 10 && data.Action == request.URL.Path {
			// 		writer.Header().Set("Content-Type", "application/json")
			// 		writer.WriteHeader(http.StatusTooManyRequests)
			// 		json.NewEncoder(writer).Encode(map[string]interface{}{
			// 			"code":   http.StatusTooManyRequests,
			// 			"status": "Too Many Requests",
			// 			"data":   "Rate limit exceeded",
			// 		})
			// 		return

			// 	}
			// }

			h.ServeHTTP(writer, request)
		})
	}
}
