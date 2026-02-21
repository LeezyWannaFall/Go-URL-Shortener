package main

import (
	"Go-URL-Shortener/internal/handler"
	"Go-URL-Shortener/internal/repository"
	"Go-URL-Shortener/internal/service"
	"Go-URL-Shortener/internal/config"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	// todo init config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// todo init db
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

	repo := repository.NewDataBase(db)
	serv := service.NewUrlService(repo)
	h := handler.NewHandler(serv)

	// todo init router: chi
	r := chi.NewRouter()
	r.Post("/shorten", h.AddShortUrl)
	r.Get("/shorten/{short}", h.GetFullUrl)
	r.Get("/{short}", h.Redirect)

	// todo start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	
	log.Printf("Server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}