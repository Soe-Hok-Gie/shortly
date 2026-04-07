package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(Id int64) (string, error) {
	fmt.Println("user.Id:", Id)

	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return "", errors.New("JWT_SECRET is empty")
	}

	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "service login",
		Subject:   strconv.FormatInt(Id, 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func GenerateRefreshToken(Id int64) (string, error) {
	secret := []byte(os.Getenv("JWT_REFRESH"))
	if len(secret) == 0 {
		return "", errors.New("JWT_REFRESH is EMPTY")
	}
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "service refresh token",
		Subject:   strconv.FormatInt(Id, 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// tambahkan Validate Token
func ValidateRefreshToken(tokenString string) (*int64, error) {
	secret := os.Getenv("REFRESH_SECRET")
	if secret == "" {
		return nil, errors.New("REFRESH_SECRET is empty")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(os.Getenv("REFRESH_SECRET")), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	Id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, errors.New("invalid subject in refresh token")
	}

	return &Id, nil

}

func Hashrefresh(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])

}
