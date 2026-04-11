package repository

import (
	"context"
	"shortly/model/domain"
)

type UrlRepository interface {
	Save(ctx context.Context, url domain.URL) (domain.URL, error)
	// Redirect(ctx context.Context, shortCode string) (domain.URL, error)
	GetAndIncrementHits(ctx context.Context, code string) (domain.URL, error)
	GetTopVisited(ctx context.Context) ([]*domain.URL, error)
	FindURLs(ctx context.Context, Params domain.FindURLParams) ([]*domain.URL, error)
	CountUrl(ctx context.Context) (int, error)
}
