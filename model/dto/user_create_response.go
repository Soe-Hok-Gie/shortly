package dto

import "time"

type UserResponse struct {
	Username   string    `json:"username"`
	Created_At time.Time `json:"created_at"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // opsional, biasanya "JWT"
}
