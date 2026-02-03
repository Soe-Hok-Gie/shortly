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

		Id, err := validateToken(jwtToken)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(request.Context(), UserIdKey, Id)
		next.ServeHTTP(writer, request.WithContext(ctx))

	})
}

func validateToken(tokenStr string) (int64, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return 0, errors.New("JWT_SECRET is empty")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	Id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, errors.New("invalid subject in token")
	}
	return Id, nil
}
