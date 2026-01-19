package repository

import (
	"context"
	"shortly/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, username string) (*domain.User, error)
}
