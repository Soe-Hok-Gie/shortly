package service

import (
	"context"
	"shortly/model/dto"
)

type UserService interface {
	Register(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error)
	// Login(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error)
	Login(ctx context.Context, input dto.CreateUserInput) (dto.LoginResponse, error)
}
