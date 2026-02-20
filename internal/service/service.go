package service

import (
	// "Go-URL-Shortener/internal/model"
	"context"
)

type UrlService struct {
	repo RepositoryInterface
}

func NewUrlService(repo RepositoryInterface) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) AddShortUrl(ctx context.Context, full string) (string, error) {
	// todo generate short link
}

func (s *UrlService) GetFullUrl(ctx context.Context, short string) (string, error) {
	// todo get full url
}