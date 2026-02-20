package repository

import (
	"context"
	"database/sql"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewDataBase(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (p *PostgresStorage) AddShortUrl(ctx context.Context, fullUrl string) (string, error) {

}

func (p *PostgresStorage) GetFullUrl(ctx context.Context, shortUrl string) (string, error) {
	
}