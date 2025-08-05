include .env
export

.PHONY: build run test clean deps docker-build docker-run migrate-up migrate-down migrate-rollback migrate-reset migrate-status swagger

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=cheki-service
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/api/main.go

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/api/main.go
	./$(BINARY_NAME)

# Test the application
test:
	$(GOTEST) -v ./...

# Test with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Test repository
test-repository:
	$(GOTEST) -v ./test/...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Download dependencies
deps:
	$(GOGET) -d -v ./...
	$(GOCMD) mod tidy

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v cmd/api/main.go

# Docker build
docker-build:
	docker build -t $(BINARY_NAME) .

# Docker run
docker-run:
	docker-compose up --build

# Docker stop
docker-stop:
	docker-compose down

# Database migration up
migrate-up:
	goose -dir migrations postgres "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

# Database migration down (rolls back all migrations)
migrate-down:
	goose -dir migrations postgres "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down-to 0

# Database migration rollback (rolls back the last migration)
migrate-rollback:
	goose -dir migrations postgres "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down

# Database migration reset (rolls back all and applies all)
migrate-reset: migrate-down migrate-up

# Database migration status
migrate-status:
	goose -dir migrations postgres "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" status

swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/api/main.go
