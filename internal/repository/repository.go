package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinksRepository struct {
	pool *pgxpool.Pool
}

func NewLinksRepository(pool *pgxpool.Pool) *LinksRepository {
	return &LinksRepository{
		pool: pool,
	}
}
