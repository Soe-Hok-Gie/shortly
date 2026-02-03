package middleware

import (
	"net/http"
	"strings"
)

type contextKey string

const UserIdKey contextKey = "Id"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, "")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(writer, "nvalid authorization format", http.StatusUnauthorized)
			return
		}

		jwtToken := tokenParts[1]

	})
}
