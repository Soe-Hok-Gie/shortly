package service

import (
	"context"
	"fmt"
	"math/rand"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"
)

type urlServiceImp struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlServiceImp{UrlRepository: urlRepository}
}

// rand package digunakan untuk generate angka random
// Jika tidak set seed, Go akan selalu generate urutan angka yang sama setiap kali program dijalankan.
// rand.Seed() memberi nilai awal (seed) untuk generator random, supaya hasilnya berbeda setiap kali program dijalankan.
// Diletakkan di init() → otomatis dijalankan sebelum main()

func (service *urlServiceImp) Save(ctx context.Context, longURL string) (domain.URL, error) {
	url := domain.URL{
		Code:     generateShortCode(6),
		LongURL:  longURL,
		HitCount: 0,
	}

	url, err := service.UrlRepository.Save(ctx, url)
	if err != nil {
		return domain.URL{}, fmt.Errorf("failed to save URL: %w", err)
	}
	return url, nil

}

func generateShortCode(length int) string {
	const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = char[rand.Intn(len(char))] //pilih indeks random dari 0 sampai len(char)-1
	}

	return string(result)
}

func (service *urlServiceImp) RedirectAndIncrement(ctx context.Context, code string) (domain.URL, error) {

	url, err := service.UrlRepository.GetAndIncrementHits(ctx, code)

	if err != nil {
		return url, fmt.Errorf("failed url: %w", err)
	}

	return url, nil
}

func (service *urlServiceImp) GetTopVisited(ctx context.Context) ([]*dto.TopLinkResponse, error) {
	url, err := service.UrlRepository.GetTopVisited(ctx)
	if err != nil {
		return nil, err
	}

	result := []*dto.TopLinkResponse{}
	for _, u := range url {
		result = append(result, &dto.TopLinkResponse{
			Code:     u.Code,
			LongURL:  u.LongURL,
			HitCount: u.HitCount,
		})
	}
	return result, nil
}
