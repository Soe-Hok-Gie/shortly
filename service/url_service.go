package service

import (
	"context"
	"shortly/model/domain"
	"shortly/model/dto"
)

type UrlService interface {
	Save(ctx context.Context, longURL string) (domain.URL, error)
	// Redirect(ctx context.Context, code string) (domain.URL, error)
	RedirectAndIncrement(ctx context.Context, code string) (domain.URL, error)
	GetTopVisited(ctx context.Context) ([]*dto.TopLinkResponse, error)
}
