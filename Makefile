.PHONY: run build migrate-up migrate-down migrate-create compose-up compose-down

# Путь до миграций и название БД
MIGRATIONS_PATH=./migrations
DB_URL=postgres://postgres:postgres@localhost:5436/pvz_db?sslmode=disable

run:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down


compose-up:
	docker-compose up --build

compose-down:
	docker-compose down -v

test:
	go test -v -coverprofile=coverage.out ./internal/usecase/...
	go tool cover -func=coverage.out

test-integration:
	go test -v ./integration_test/...
