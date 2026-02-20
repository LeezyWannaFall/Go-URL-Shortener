package handler

import (
	"encoding/json"
	"net/http"
)

type UrlHandler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *UrlHandler {
	return &UrlHandler{service: service}
}

func (h *UrlHandler) AddShortUrl(w http.ResponseWriter, r *http.Request) {
	var url string

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, "invalid bodyy", http.StatusBadRequest)
	}

	shrtUrl, err := h.service.AddShortUrl(r.Context(), url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shrtUrl)
}

func (h *UrlHandler) GetFullUrl(w http.ResponseWriter, r * http.Request) {
	var url string

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, "invalid bodyy", http.StatusBadRequest)
	}

	fullUrl, err := h.service.GetFullUrl(r.Context(), url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fullUrl)
}