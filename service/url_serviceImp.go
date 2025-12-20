package service

import (
	"context"
	"math/rand"
	"shortly/model/domain"
	"shortly/repository"
	"time"
)

type urlServiceImp struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlServiceImp{UrlRepository: urlRepository}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (service *urlServiceImp) Save(ctx context.Context, longURL string) (domain.URL, error) {
	url := domain.URL{
		Code:     generateShortCode(6),
		LongURL:  longURL,
		HitCount: 0,
	}

	return service.UrlRepository.Save(ctx, url)

}

func generateShortCode(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}
