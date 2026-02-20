package handler

import (
	"context"
)

type ServiceInterface interface {
	AddShortUrl(ctx context.Context, FullUrl string) (string, error)
	GetFullUrl(ctx context.Context, ShortUrl string) (string, error)
}