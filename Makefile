DOCKER_COMPOSE=docker compose

DOCKER_MAC_BIN := /Applications/Docker.app/Contents/Resources/bin/
PATH := $(PATH):$(DOCKER_MAC_BIN)

.PHONY: start stop

build:
	@echo "Building ..."
	$(DOCKER_COMPOSE) build

start:
	@echo "Starting services ..."
	$(DOCKER_COMPOSE) up -d
	@echo "Starting application ..."

stop:
	@echo "Stopping services ..."
	$(DOCKER_COMPOSE) down

test:
	@echo "Run test ..."
	go test ./...
