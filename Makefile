SHELL = /bin/bash

NAME = gotifications
SRC_FOLDER = .
MAIN = $(SRC_FOLDER)/main.go
SET_ENV = godotenv -f .env.local,.env

DOCKER_IMG = $(NAME)-img
DOCKER_CONTAINER = $(NAME)-container

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

.PHONY: docker.build
docker.build:
	docker build -t dlacreme/gotifications:latest .

.PHONY: docker.run
docker.run:
	$(SET_ENV) docker run -p 3333:1042 --name $(DOCKER_CONTAINER) $(DOCKER_IMG)

.PHONY: docker.stop
docker.stop:
	docker stop $(DOCKER_CONTAINER)