.PHONY: run build migrate-up migrate-down migrate-create compose-up compose-down

# Путь до миграций и название БД
MIGRATIONS_PATH=./migrations
DB_URL=postgres://postgres:postgres@localhost:5436/pvz_db?sslmode=disable

# запуск приложения
run:
	go run cmd/main.go

# билд приложения
build:
	go build -o bin/app cmd/main.go

# миграции вверх
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

# миграции вниз
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

# создать новую миграцию (пример: make migrate-create name=create_users_table)
migrate-create:
ifndef name
	$(error name is required. Example: make migrate-create name=create_users_table)
endif
	migrate create -ext go -dir $(MIGRATIONS_PATH) -seq $(name)

# запустить docker-compose
compose-up:
	docker-compose up --build

# остановить docker-compose и удалить volume
compose-down:
	docker-compose down -v