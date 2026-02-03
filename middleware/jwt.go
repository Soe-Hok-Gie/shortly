package middleware

import (
	"errors"
	"net/http"
	"os"
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

	})
}

func validateToken(tokenStr string) (int64, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return 0, errors.New("JWT_SECRET is empty")
	}
	token, err := JWTParsWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.claims(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

}
