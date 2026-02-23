package main

import (
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/handler"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/service"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/config"
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/repository"

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

	repo := repository.ChooseMemory(cfg)
	serv := service.NewUrlService(repo)
	h := handler.NewHandler(serv)

	// init router: chi
	r := chi.NewRouter()
	r.Post("/shorten", h.AddShortUrl)
	r.Get("/{short}", h.Redirect)

	// start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	
	log.Fatal(http.ListenAndServe(addr, r))
}