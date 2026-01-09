package dto

import "time"

type UserResponse struct {
	Username   string    `json:"username"`
	Created_At time.Time `json:"created_at"`
}
