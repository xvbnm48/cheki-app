.PHONY: build run test clean deps docker-build docker-run

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
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/chekisvc?sslmode=disable" up

# Database migration down
migrate-down:
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/chekisvc?sslmode=disable" down

swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/api/main.go
