package handler

import (
	"context"
	"net/http"
)

type LinkService interface {
	Shorten(ctx context.Context, rawURL string) (string, error)
	Resolve(ctx context.Context, code string) (string, error)
}

type LinkHTTPHandler struct {
	linkService LinkService
	baseURL     string
}

func NewLinkHTTPHandler(LinkService LinkService, baseURL string) *LinkHTTPHandler {
	return &LinkHTTPHandler{
		linkService: LinkService,
		baseURL:     baseURL,
	}
}

func (h *LinkHTTPHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", h.Shorten)
	mux.HandleFunc("GET /{code}", h.Redirect)

	return mux
}
