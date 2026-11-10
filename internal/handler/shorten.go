package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jettyjunk/go-url-shortener/internal/service"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *LinkHTTPHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req shortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	code, err := h.linkService.Shorten(ctx, req.URL)
	if err != nil {
		if errors.Is(err, service.ErrInvalidURL) {
			http.Error(w, "invalid url", http.StatusBadRequest)
			return
		}
		log.Printf("shorten: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shortenResponse{ShortURL: h.baseURL + "/" + code})
}
