package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/jettyjunk/go-url-shortener/internal/repository"
)

func (h *LinkHTTPHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.PathValue("code")

	url, err := h.linkService.Resolve(ctx, code)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		log.Printf("redirect: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
