package repository

import (
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
	"context"
	"database/sql"
	"errors"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewDataBase(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (r *PostgresStorage) Save(ctx context.Context, url *model.URL) error {
	query := `INSERT INTO urls (full_url, short_url) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, url.Full, url.Short)
	if err != nil {
		return errors.New("not found")
	}

	return nil 
}

func (r *PostgresStorage) GetByShortUrl(ctx context.Context, short string) (*model.URL, error) {
	var url model.URL
	query := `SELECT full_url FROM urls WHERE short_url = $1`

	err := r.db.QueryRowContext(ctx, query, short).Scan(&url.Full)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *PostgresStorage) GetByFullUrl(ctx context.Context, full string) (*model.URL, error) {
	var url model.URL
	query := `SELECT short_url FROM urls WHERE full_url = $1`

	err := r.db.QueryRowContext(ctx, query, full).Scan(&url.Short)
	if err != nil {
		return nil, err
	}

	return &url, nil
}