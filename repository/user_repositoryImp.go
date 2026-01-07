package repository

import (
	"context"
	"database/sql"
	"shortly/model/domain"
)

type userRepositoryImp struct {
	DB *sql.DB
}

func (repository *userRepositoryImp) Save(ctx context.Context, user domain.User) (error, domain.User)
