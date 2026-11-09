package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *LinksRepository) GetCode(ctx context.Context, code string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	query := `
	SELECT original_url 
	FROM links
	WHERE code=$1
	`

	var url string

	err := r.pool.QueryRow(ctx, query, code).Scan(&url)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrNotFound
		}

		return "", fmt.Errorf("scan error: %w", err)
	}

	return url, nil
}
