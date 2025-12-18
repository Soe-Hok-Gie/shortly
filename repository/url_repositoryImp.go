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
	script := "INSERT INTO urls (code,long_url,hit_count) VALUES ?,?,?"
	result, err := repository.DB.ExecContext(ctx, script, url.Code, url.LongURL, url.HitCount)
	if err != nil {
		return url, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		url.Id = id
	}
	return url, nil
}
