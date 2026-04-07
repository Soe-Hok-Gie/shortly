package service

import (
	"context"
	"shortly/model/dto"
)

type UserService interface {
	Register(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error)
	// Login(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error)
	Login(ctx context.Context, input dto.CreateUserInput) (dto.LoginResponse, error)
	Refresh(ctx context.Context, RefreshToken string) (dto.LoginResponse, error)
	DeleteRefresh(ctx context.Context, hash string) error
}
