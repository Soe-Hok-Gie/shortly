package repository

import (
	"context"
	"shortly/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (error, domain.User)
}
