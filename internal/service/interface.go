package service

import (
	"Go-URL-Shortener/internal/model"
	"context"
)

type RepositoryInterface interface {
	Save(ctx context.Context, url *model.URL) error
	GetByShortUrl(ctx context.Context, short string) (*model.URL, error)
	GetByFullUrl(ctx context.Context, full string) (*model.URL, error)
}