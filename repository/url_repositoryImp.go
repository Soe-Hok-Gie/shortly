package repository

import (
	"context"
	"database/sql"
	"shortly/model/domain"
)

type urlRepositoryImp struct {
	DB *sql.DB
}

func NewUrlRepository(DB *sql.DB) UrlRepository {
	return &urlRepositoryImp{DB: DB}
}

func (repository *urlRepositoryImp) Save(ctx context.Context, url domain.URL) (domain.URL, error) {
	script := "INSERT INTO urls (code,long_url,hit_count) VALUES (?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, url.Code, url.LongURL, url.HitCount)
	if err == nil {
		return url, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		url.Id = id
	}
	return url, nil
}
