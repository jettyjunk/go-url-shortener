package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type LinkStore interface {
	Create(ctx context.Context, code string, originalURL string) error
	GetCode(ctx context.Context, code string) (string, error)
}

type CachedLinks struct {
	next LinkStore
	rdb  *redis.Client
	ttl  time.Duration
}

func NewCachedLinks(next LinkStore, rdb *redis.Client, ttl time.Duration) *CachedLinks {
	return &CachedLinks{
		next: next,
		rdb:  rdb,
		ttl:  ttl,
	}
}

func (c *CachedLinks) Create(ctx context.Context, code string, originalURL string) error {
	return c.next.Create(ctx, code, originalURL)
}

func (c *CachedLinks) GetCode(ctx context.Context, code string) (string, error) {
	key := "link:" + code

	res, err := c.rdb.Get(ctx, key).Result()
	if err == nil {
		return res, nil
	}
	if !errors.Is(err, redis.Nil) {
		log.Printf("cache get %s: %v", key, err)
	}

	url, err := c.next.GetCode(ctx, code)
	if err != nil {
		return "", err
	}

	if err := c.rdb.Set(ctx, key, url, c.ttl).Err(); err != nil {
		log.Printf("cache set %s: %v", key, err)
	}

	return url, nil
}
