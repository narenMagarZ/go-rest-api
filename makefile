# Project settings
APP_NAME := app
BIN_DIR := ./bin
DOCKER_IMAGE := go-app
DOCKER_CONTAINER := go-rest-api-app
COMPOSE_FILE := docker-compose.yml
SRC := ./cmd/api/main.go
BINARY := $(BIN_DIR)/$(APP_NAME)

# Go commands
.PHONY: build run clean test format lint docker-up docker-down logs db-migrate

## Build the Go binary
build:
	mkdir -p ./bin 
	go build -o $(BINARY) $(SRC)

## Run the Go binary locally
run:
	$(BINARY)

## Remove the built binary
clean:
	rm -rf $(BIN_DIR)

## Format the code
format:
	go fmt ./...

## Start Docker containers

docker-build:
	docker compose build 

docker-up:
	docker compose -f $(COMPOSE_FILE) up --build -d

## Stop Docker containers
docker-down:
	docker compose -f $(COMPOSE_FILE) down

## View app logs
logs:
	docker logs -f $(DOCKER_CONTAINER)

## Help: show available commands
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  build         Build the Go binary"
	@echo "  run           Run the app locally"
	@echo "  clean         Remove built binary"
	@echo "  format        Format code"
	@echo "  docker-build  Build docker image"
	@echo "  docker-up     Build and start containers"
	@echo "  docker-down   Stop containers"
	@echo "  logs          Tail app container logs"
