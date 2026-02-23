package repository

import (
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/config"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/repository/memory"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/repository/postgres"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/service"

	"database/sql"
	"fmt"
	"log"
)

func ChooseMemory(cfg *config.Config) service.Repository {
	switch cfg.Storage.Type {
	case "postgres":
		connStr := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.Storage.Postgres.User,
			cfg.Storage.Postgres.Password,
			cfg.Storage.Postgres.Host,
			cfg.Storage.Postgres.Port, 
			cfg.Storage.Postgres.Database,
		)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal("cant connect to database:", err)
		}

		log.Println("Storage: Postgres connected")
        return postgres.NewDataBase(db)

	case "memory":
		log.Println("Storage: In-Memory initialized")
        return memory.NewInMemoryStorage()

	default:
        log.Fatalf("Unknown storage type: %s", cfg.Storage.Type)
		return nil
    }
}