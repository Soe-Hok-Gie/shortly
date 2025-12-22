package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	if err != nil {
		fmt.Println("err", err)
		return url, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err", err)

		url.Id = id
	}
	return url, nil
}

func (repository *urlRepositoryImp) FindByShortCode(ctx context.Context, code string) (domain.URL, error) {
	script := "SELECT id, code,long_url from urls WHERE code =?"
	row := repository.DB.QueryRowContext(ctx, script, code)

	var url domain.URL
	err := row.Scan(
		url.Id,
		url.Code,
		url.LongURL,
		url.HitCount,
	)
	if err != nil {
		fmt.Println("err", err)
		return url, err
	}
	return url, nil
}
