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

// func (repository *urlRepositoryImp) Redirect(ctx context.Context, code string) (domain.URL, error) {
// 	script := "SELECT id, code,long_url from urls WHERE code =?"
// 	row := repository.DB.QueryRowContext(ctx, script, code)

// 	var url domain.URL
// 	err := row.Scan(
// 		&url.Id,
// 		&url.Code,
// 		&url.LongURL,
// 	)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		return url, err
// 	}
// 	return url, nil
// }

func (repository *urlRepositoryImp) GetAndIncrementHits(ctx context.Context, code string) (domain.URL, error) {
	hitscript := "UPDATE urls set hit_count=hit_count + 1 WHERE code = ?"
	if _, err := repository.DB.ExecContext(ctx, hitscript, code); err != nil {
		return domain.URL{}, err
	}

	script := "SELECT id, code, long_url, hit_count FROM urls WHERE code = ?"
	row := repository.DB.QueryRowContext(ctx, script, code)

	var url domain.URL
	err := row.Scan(&url.Id, &url.Code, &url.LongURL, &url.HitCount)
	if err != nil {
		return url, err
	}

	return url, nil

}

func (repository *urlRepositoryImp) GetTopVisited(ctx context.Context) ([]*domain.URL, error) {
	script := "SELECT code, long_url, hit_count FROM urls ORDER BY hit_count DESC LIMIT 10"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.URL
	for rows.Next() {
		//buat var baru dalam loop untuk mencegah Bug Pointer Reuse / Loop Variable Capture
		url := new(domain.URL)
		if err := rows.Scan(&url.Code, &url.LongURL, &url.HitCount); err != nil {
			return nil, err
		}
		result = append(result, url)
	}
	return result, nil
}
