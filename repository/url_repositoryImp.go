package repository

import (
	"context"
	"database/sql"
	"shortly/model/domain"
)

type urlRepositoryImp struct {
	DB *sql.DB
}

func (repository *urlRepositoryImp) save(ctx context.Context, url domain.URL) (domain.URL, error) {

}
