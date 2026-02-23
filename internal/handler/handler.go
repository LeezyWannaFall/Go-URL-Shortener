package handler

import (
	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UrlHandler struct {
	service Service
}

func NewHandler(service Service) *UrlHandler {
	return &UrlHandler{service: service}
}

func (h *UrlHandler) AddShortUrl(w http.ResponseWriter, r *http.Request) {
	var req model.URL

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	shrtUrl, err := h.service.AddShortUrl(r.Context(), req.Full)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shrtUrl)
}

func (h *UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short")

	fullUrl, err := h.service.GetFullUrl(r.Context(), short)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
}