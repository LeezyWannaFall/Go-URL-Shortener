package handler

import (
	"context"
)

type ServiceInterface interface {
	AddShortUrl(ctx context.Context, fullUrl string) (string, error)
	GetFullUrl(ctx context.Context, shortUrl string) (string, error)
}