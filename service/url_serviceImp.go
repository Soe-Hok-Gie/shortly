package service

import (
	"context"
	"shortly/model/domain"
	"shortly/repository"
)

type urlServiceImp struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlServiceImp{UrlRepository: urlRepository}
}

func (service *urlServiceImp) Save(ctx context.Context, longURL string) (domain.URL, error) {
	url := domain.URL{
		Code:     "abc123", // nanti bisa diganti random generator
		LongURL:  longURL,
		HitCount: 0,
	}

	return service.UrlRepository.Save(ctx, url)

}
