package repository

import (
	"context"
	"shortly/model/domain"
)

type UrlRepository interface {
	Save(ctx context.Context, url domain.URL) (domain.URL, error)
	// Redirect(ctx context.Context, shortCode string) (domain.URL, error)
	GetAndIncrementHits(ctx context.Context, code string) (domain.URL, error)
}
