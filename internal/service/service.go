package service

import (
	"Go-URL-Shortener/internal/model"
	"context"
	"errors"
)

type UrlService struct {
	repo RepositoryInterface
}

func NewUrlService(repo RepositoryInterface) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) AddShortUrl(ctx context.Context, full string) (string, error) {
	if !isValidUrl(full) {
		return "", errors.New("invalid URL")
	}

	url, err := s.repo.GetByFullUrl(ctx, full)
	if err == nil && url.Short != "" {
		return url.Short, nil
	}

	if err != nil {
		return "", errors.New("repository error")
	}

	for i := 0; i < 5; i++ {
		newUrl := model.URL{Short: ShortLinkGenerator(), Full: full}
		errSave := s.repo.Save(ctx, newUrl)
		if errSave == nil { 
			return newUrl.Short, nil
		}
	}

	return "", errors.New("failed to generate unique short link after several attempts")
}

func (s *UrlService) GetFullUrl(ctx context.Context, short string) (string, error) {
	url, err := s.repo.GetByShortUrl(ctx, short)
	if err != nil {
		return "", err
	}

	if url.Full == "" {
		return "", errors.New("Full link not found")
	}

	return url.Full, nil
}