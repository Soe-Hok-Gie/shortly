package service

import (
	"context"
	"shortly/model/domain"
	"shortly/model/dto"
)

type UserService interface {
	Save(ctx context.Context, input dto.CreateUserInput) (domain.User, error)
}
