package memory

import (
	"context"
	"testing"

	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
)

func TestInMemoryStorage(t *testing.T) {
	tests := []struct {
		name        string
		initialData *model.URL
		saveData	*model.URL
		expectError   bool	
	}{
		{
			name: "save success",
			initialData: nil,
			saveData: &model.URL{Full: "https://finance.ozon.ru", Short: "lol123"},
			expectError: false,
		},
		{
			name: "short already exists",
			initialData: &model.URL{Full: "https://finance.ozon.ru", Short: "lol123"},
			saveData: &model.URL{Full: "https://github.com", Short: "lol123"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewInMemoryStorage()

			if tt.initialData != nil {
				err := storage.Save(context.Background(), tt.initialData)
				if err != nil {
					t.Fatalf("Failed to set up initial data: %v", err)
				}
			}

			err := storage.Save(context.Background(), tt.saveData)
			if (err != nil) != tt.expectError {
				t.Errorf("Save() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestInMemoryStorage_GetByShortUrl(t *testing.T) {
	tests := []struct {
		name        string
		initialData *model.URL
		short       string
		expectError bool
	}{
		{
			name:        "found",
			initialData: &model.URL{Short: "abc", Full: "https://github.com"},
			short:       "abc",
			expectError: false,
		},
		{
			name:        "not found",
			initialData: nil,
			short:       "abc",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewInMemoryStorage()
			ctx := context.Background()

			if tt.initialData != nil {
				_ = storage.Save(ctx, tt.initialData)
			}

			url, err := storage.GetByShortUrl(ctx, tt.short)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectError && url.Full != tt.initialData.Full {
				t.Errorf("expected full %s, got %s", tt.initialData.Full, url.Full)
			}
		})
	}
}

func TestInMemoryStorage_GetByFullUrl(t *testing.T) {
	tests := []struct {
		name        string
		initialData *model.URL
		full        string
		expectError bool
	}{
		{
			name:        "found",
			initialData: &model.URL{Short: "abc", Full: "https://google.com"},
			full:        "https://google.com",
			expectError: false,
		},
		{
			name:        "not found",
			initialData: nil,
			full:        "https://google.com",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewInMemoryStorage()
			ctx := context.Background()

			if tt.initialData != nil {
				_ = storage.Save(ctx, tt.initialData)
			}

			url, err := storage.GetByFullUrl(ctx, tt.full)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectError && url.Short != tt.initialData.Short {
				t.Errorf("expected short %s, got %s", tt.initialData.Short, url.Short)
			}
		})
	}
}