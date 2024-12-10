APP_NAME := bookstore
GOOSE_MIGRATIONS_DIR := ./sql/migrations
GOOSE_DBSTRING := postgres://postgres:postgres@127.0.0.1:5432/postgres
GOOSE_DRIVER := postgres

run-local:
	go build -o ./build/$(APP_NAME) ./cmd/bookstore/main.go
	CONFIG_PATH=./config/local.yml ./build/$(APP_NAME)

run-prod:
	go build -o ./build/$(APP_NAME) ./cmd/bookstore/main.go
	CONFIG_PATH=./config/prod.yml ./build/$(APP_NAME)

run:
	go build -o ./build/$(APP_NAME) ./cmd/bookstore/main.go
	./build/$(APP_NAME)

migrations:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir $(GOOSE_MIGRATIONS_DIR) up
