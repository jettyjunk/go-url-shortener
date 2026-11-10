include .env
export

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable

.PHONY: up down migrate-up migrate-down psql run

build:
	go build ./cmd/... ./internal/...

test:
	go test ./internal/...

up:
	docker compose up -d

down: 
	docker compose down

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

psql:
	docker compose exec shortener-postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

run: 
	go run ./cmd/app
