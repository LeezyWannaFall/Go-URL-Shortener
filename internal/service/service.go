package service

import (
	"context"
	"Go-URL-Shortener/internal/model"
)

type UrlService struct {
	repo RepositoryInterface
}

func NewUrlService(repo RepositoryInterface) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) AddShortUrl(ctx context.Context, fullUrl string) (string, error) {

}

func (s *UrlService) GetFullUrl(ctx context.Context, shortUrl string) (string, error) {

}