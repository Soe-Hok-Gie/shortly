package middleware

import "net/http"

type contextKey string

const UserIdKey contextKey = "Id"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "authorization header missing", http.StatusUnauthorized)
			return
		}

	})
}
