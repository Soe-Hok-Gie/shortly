package repository

import (
	"context"
	"shortly/model/domain"
)

type UrlRepository interface {
	save(ctx context.Context, url domain.URL) (domain.URL, error)
}
