package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIdKey contextKey = "Id"

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			authHeader := request.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(writer, "authorization header missing", http.StatusUnauthorized)
				return
			}

			// tokenParts := strings.Split(authHeader, " ")
			// if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			// 	http.Error(writer, "nvalid authorization format", http.StatusUnauthorized)
			// 	return
			// }

			// jwtToken := tokenParts[1]

			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(writer, "invalid authorization format", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == "" {
				http.Error(writer, "token missing", http.StatusUnauthorized)
				return
			}

			Id, err := validateToken(tokenStr)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(request.Context(), UserIdKey, *Id)
			next.ServeHTTP(writer, request.WithContext(ctx))

		})
	}
}

func validateToken(tokenStr string) (*int64, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is empty")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	Id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, errors.New("invalid subject in token")
	}
	return &Id, nil
}
