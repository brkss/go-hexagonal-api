# Go parameters
BINARY_NAME=dextrace-server
MAIN_FILE=cmd/main.go
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
GOMOD=$(GO) mod

# Database 
DB_CONTAINER="postgres-dex"
DB_NAME="dextrace"
DB_USER="root"
DB_PASS="secret"

# Build flags
LDFLAGS=-ldflags "-s -w"

.PHONY: all build clean test run deps tidy

all: clean build

build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET) -v ./...

tidy:
	$(GOMOD) tidy

# Development tools
.PHONY: lint fmt

lint:
	golangci-lint run

fmt:
	$(GO) fmt ./...

# Database migrations
.PHONY: migrate-up migrate-down

createdb:
	docker exec -it $(DB_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

migrate-up:
	$(GO) run github.com/golang-migrate/migrate/v4/cmd/migrate -database "postgresql://postgres:postgres@localhost:5432/dextrace?sslmode=disable" -path internal/infrastructure/database/migrations up

migrate-down:
	$(GO) run github.com/golang-migrate/migrate/v4/cmd/migrate -database "postgresql://postgres:postgres@localhost:5432/dextrace?sslmode=disable" -path internal/infrastructure/database/migrations down
