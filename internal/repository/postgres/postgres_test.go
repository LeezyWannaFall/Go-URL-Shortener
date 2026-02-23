package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
)

func TestPostgresStorage_Save(t *testing.T) {
	tests := []struct {
		name        string
		url         *model.URL
		mockError   error
		expectError bool
	}{
		{
			name:        "success",
			url:         &model.URL{Full: "https://github.com/LeezyWannaFall", Short: "abc"},
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "insert error",
			url:         &model.URL{Full: "https://github.com/LeezyWannaFall", Short: "abc"},
			mockError:   errors.New("db error"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, _ := sqlmock.New()
			defer db.Close()

			repo := NewDataBase(db)

			query := `INSERT INTO urls \(full_url, short_url\) VALUES \(\$1, \$2\)`

			if tt.mockError == nil {
				mock.ExpectExec(query).
					WithArgs(tt.url.Full, tt.url.Short).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec(query).
					WithArgs(tt.url.Full, tt.url.Short).
					WillReturnError(tt.mockError)
			}

			err := repo.Save(context.Background(), tt.url)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestPostgresStorage_GetByShortUrl(t *testing.T) {
	tests := []struct {
		name        string
		short       string
		mockError   error
		expectError bool
	}{
		{
			name:        "found",
			short:       "abc",
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "not found",
			short:       "abc",
			mockError:   errors.New("not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, _ := sqlmock.New()
			defer db.Close()

			repo := NewDataBase(db)

			query := `SELECT full_url FROM urls WHERE short_url = \$1`

			if tt.mockError == nil {
				rows := sqlmock.NewRows([]string{"full_url"}).
					AddRow("https://google.com")

				mock.ExpectQuery(query).
					WithArgs(tt.short).
					WillReturnRows(rows)
			} else {
				mock.ExpectQuery(query).
					WithArgs(tt.short).
					WillReturnError(tt.mockError)
			}

			url, err := repo.GetByShortUrl(context.Background(), tt.short)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && url.Full != "https://google.com" {
				t.Errorf("unexpected full url")
			}
		})
	}
}

func TestPostgresStorage_GetByFullUrl(t *testing.T) {
	tests := []struct {
		name        string
		full        string
		mockError   error
		expectError bool
	}{
		{
			name:        "found",
			full:        "https://google.com",
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "not found",
			full:        "https://google.com",
			mockError:   errors.New("not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, _ := sqlmock.New()
			defer db.Close()

			repo := NewDataBase(db)

			query := `SELECT short_url FROM urls WHERE full_url = \$1`

			if tt.mockError == nil {
				rows := sqlmock.NewRows([]string{"short_url"}).
					AddRow("abc")

				mock.ExpectQuery(query).
					WithArgs(tt.full).
					WillReturnRows(rows)
			} else {
				mock.ExpectQuery(query).
					WithArgs(tt.full).
					WillReturnError(tt.mockError)
			}

			url, err := repo.GetByFullUrl(context.Background(), tt.full)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && url.Short != "abc" {
				t.Errorf("unexpected short url")
			}
		})
	}
}

