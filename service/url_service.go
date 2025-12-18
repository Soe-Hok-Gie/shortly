package service

import (
	"context"
	"shortly/model/domain"
)

type UrlService interface {
	Save(ctx context.Context, longURL string) (domain.URL, error)
}
