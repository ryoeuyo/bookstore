APP_NAME := bookstore

run-local:
	go build -o ./build/$(APP_NAME) ./cmd/main.go
	CONFIG_PATH=./config/local.yml ./build/$(APP_NAME)

run-prod:
	go build -o ./build/$(APP_NAME) ./cmd/main.go
	CONFIG_PATH=./config/prod.yml ./build/$(APP_NAME)

run:
	go build -o ./build/$(APP_NAME) ./cmd/main.go
	./build/$(APP_NAME)
