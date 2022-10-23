SHELL = /bin/bash

SRC_FOLDER = .
MAIN = $(SRC_FOLDER)/main.go
SET_ENV = godotenv -f .env.local,.env

.PHONY: all
all: run

.PHONY: run # Starts golive
run:
	$(SET_ENV) go run $(MAIN)

.PHONE: build
build:
	GOOS=linux GOARCH=amd64 go build $(MAIN) -tags netgo -o gauth

.PHONY: test ## Run tests
test:
	$(SET_ENV) go test ./...
