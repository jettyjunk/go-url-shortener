package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

const queryTimeout = 5 * time.Second

func (r *LinksRepository) Create(ctx context.Context, code string, originalURL string) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	query := `
	INSERT INTO links (code, original_url)
	VALUES ($1, $2)
	`

	_, err := r.pool.Exec(ctx, query, code, originalURL)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrCodeExists
		}
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
