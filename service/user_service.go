package service

import (
	"context"
	"shortly/model/dto"
)

type UserService interface {
	Save(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error)
}
