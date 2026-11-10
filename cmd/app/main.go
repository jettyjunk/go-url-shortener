package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jettyjunk/go-url-shortener/config"
	"github.com/jettyjunk/go-url-shortener/internal/handler"
	"github.com/jettyjunk/go-url-shortener/internal/repository"
	"github.com/jettyjunk/go-url-shortener/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("app: %v", err)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("created pool: %w", err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping postgres: %w", err)
	}

	repo := repository.NewLinksRepository(pool)
	svc := service.NewLinkService(repo)
	h := handler.NewLinkHTTPHandler(svc, cfg.BaseURL)

	server := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: h.Routes(),
	}

	log.Printf("listening on %s", cfg.HTTPAddr)

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}
