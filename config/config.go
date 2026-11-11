package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	HTTPAddr        string
	BaseURL         string
	ShutdownTimeout time.Duration
	PostgresUser    string
	PostgresPass    string
	PostgresDB      string
	RedisAddr       string
	CacheTTL        time.Duration
}

func Load() (Config, error) {
	addr, err := requireEnv("HTTP_ADDR")
	if err != nil {
		return Config{}, err
	}

	url, err := requireEnv("BASE_URL")
	if err != nil {
		return Config{}, err
	}

	puser, err := requireEnv("POSTGRES_USER")
	if err != nil {
		return Config{}, err
	}

	ppass, err := requireEnv("POSTGRES_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	pdb, err := requireEnv("POSTGRES_DB")
	if err != nil {
		return Config{}, err
	}

	raw, err := requireEnv("HTTP_SHUTDOWN_TIMEOUT")
	if err != nil {
		return Config{}, err
	}

	timeout, err := time.ParseDuration(raw)
	if err != nil {
		return Config{}, fmt.Errorf("parse HTTP_SHUTDOWN_TIMEOUT: %w", err)
	}

	rds, err := requireEnv("REDIS_ADDR")
	if err != nil {
		return Config{}, err
	}

	rawTTL, err := requireEnv("CACHE_TTL")
	if err != nil {
		return Config{}, err
	}

	ttl, err := time.ParseDuration(rawTTL)
	if err != nil {
		return Config{}, fmt.Errorf("parse CACHE_TTL: %w", err)
	}

	if ttl <= 0 {
		return Config{}, fmt.Errorf("CACHE_TTL must be positive")
	}

	return Config{
		HTTPAddr:        addr,
		BaseURL:         url,
		ShutdownTimeout: timeout,
		PostgresUser:    puser,
		PostgresPass:    ppass,
		PostgresDB:      pdb,
		RedisAddr:       rds,
		CacheTTL:        ttl,
	}, nil
}

func requireEnv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", fmt.Errorf("env %s is required", name)
	}

	return v, nil
}

func (c Config) PostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		c.PostgresUser, c.PostgresPass, c.PostgresDB,
	)
}
