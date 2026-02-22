package service

import (
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
	"context"
	"errors"
)

type UrlService struct {
	repo Repository
}

func NewUrlService(repo Repository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) AddShortUrl(ctx context.Context, full string) (string, error) {
	if !isValidUrl(full) {
		return "", errors.New("invalid URL")
	}

	maxAttempts := 5
	url, err := s.repo.GetByFullUrl(ctx, full)

	if err != nil {
		for i := 0; i < maxAttempts; i++ {
			generatedUrl := model.URL{Short: GenerateShortLink(), Full: full}
			errSave := s.repo.Save(ctx, &generatedUrl)

			if errSave == nil { 
				return generatedUrl.Short, nil
			}
		}
	} else {
		return url.Short, nil
	}

	return "", errors.New("failed to generate unique short link after several attempts")
}

func (s *UrlService) GetFullUrl(ctx context.Context, short string) (string, error) {
	url, err := s.repo.GetByShortUrl(ctx, short)
	if err != nil {
		return "", err
	}

	return url.Full, nil
}