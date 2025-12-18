package repository

import (
	"context"
	"shortly/model/domain"
)

type UrlRepository interface {
	Save(ctx context.Context, url domain.URL) (domain.URL, error)
}
