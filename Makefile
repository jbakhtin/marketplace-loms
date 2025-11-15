APP_NAME ?= marketplace-loms
IMAGE ?= $(APP_NAME)
TAG ?= latest
CONTAINER ?= $(APP_NAME)
PORT ?= 8080

.PHONY: build docker-build docker-run docker-stop docker-logs

build:
	@echo "Building Go binaries..."
	@go build ./...

docker-build:
	@echo "Building Docker image $(IMAGE):$(TAG)..."
	@docker build -t $(IMAGE):$(TAG) .

docker-run: docker-build
	@echo "Starting container $(CONTAINER)..."
	@docker run --rm -d --name $(CONTAINER) -p $(PORT):$(PORT) $(IMAGE):$(TAG)

docker-stop:
	@echo "Stopping container $(CONTAINER)..."
	@docker stop $(CONTAINER) || true

docker-logs:
	@docker logs -f $(CONTAINER)

