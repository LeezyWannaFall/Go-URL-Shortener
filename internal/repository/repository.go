package repository

import (
	"Go-URL-Shortener/internal/model"
	"context"
	"database/sql"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewDataBase(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (r *PostgresStorage) Save(ctx context.Context,url model.URL) error {

}

func (r *PostgresStorage) GetByShortUrl(ctx context.Context, short string) (model.URL, error) {

}

func (r *PostgresStorage) GetByFullUrl(ctx context.Context, full string) (model.URL, error) {
	
}