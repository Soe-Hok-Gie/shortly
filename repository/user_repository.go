package repository

import (
	"context"
	"shortly/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
}
