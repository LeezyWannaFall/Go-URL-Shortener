package main

import (
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/handler"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/repository/postgres"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/repository/memory"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/service"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/config"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	// init config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var repo service.Repository

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
        repo = postgres.NewDataBase(db)

	case "memory":
		log.Println("Storage: In-Memory initialized")
        repo = memory.NewInMemoryStorage()

	default:
        log.Fatalf("Unknown storage type: %s", cfg.Storage.Type)
    }

	serv := service.NewUrlService(repo)
	h := handler.NewHandler(serv)

	// init router: chi
	r := chi.NewRouter()
	r.Post("/shorten", h.AddShortUrl)
	r.Get("/shorten/{short}", h.GetFullUrl)
	r.Get("/{short}", h.Redirect)

	// start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	
	log.Printf("Server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}