package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
)

type InMemoryStorage struct {
	shortToFull map[string]string
	fullToShort map[string]string
	mu   sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		shortToFull: make(map[string]string),
		fullToShort: make(map[string]string),
	}
}

func (m *InMemoryStorage) Save(ctx context.Context, url *model.URL) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.shortToFull[url.Short]
	if ok {
		return errors.New("short url already exists")
	}

	m.shortToFull[url.Short] = url.Full
	m.fullToShort[url.Full] = url.Short
	return nil
}

func (m *InMemoryStorage) GetByShortUrl(ctx context.Context, short string) (*model.URL, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	full, ok := m.shortToFull[short]
	if !ok {
		return &model.URL{}, errors.New("short url not found")
	}

	return &model.URL{Full: full, Short: short}, nil
}

func (m *InMemoryStorage) GetByFullUrl(ctx context.Context, full string) (*model.URL, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	short, ok := m.fullToShort[full]
	if !ok {
		return &model.URL{}, errors.New("full url not found")
	}

	return &model.URL{Full: full, Short: short}, nil
}