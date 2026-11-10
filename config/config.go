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

	return Config{
		HTTPAddr:        addr,
		BaseURL:         url,
		ShutdownTimeout: timeout,
		PostgresUser:    puser,
		PostgresPass:    ppass,
		PostgresDB:      pdb,
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
