.PHONY: generate build run test clean

generate:
	@echo "Generating code..."
	go generate ./...

build:
	@echo "Building..."
	go build -o bin/api cmd/api/main.go

run:
	@echo "Running..."
	go run cmd/api/main.go

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	rm -rf bin

migrate-up:
	@echo "Running migrations..."
	migrate -path migrations -database "postgresql://orderuser:orderpass@localhost:5432/orderdb?sslmode=disable" up

migrate-down:
	@echo "Reverting migrations..."
	migrate -path migrations -database "postgresql://orderuser:orderpass@localhost:5432/orderdb?sslmode=disable" down

generate-sqlc:
	@echo "Generating sqlc code..."
	sqlc generate
	
.DEFAULT_GOAL := build