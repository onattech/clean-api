# Use the 'make' command without any arguments to list all the available commands

.PHONY: list # Default command. Won't get listed
list:
	@echo "📋 Available commands:"
	@awk -F':.*?## ' '/^[a-zA-Z0-9_-]+:/ && !/^[[:blank:]]*list:/ { if ($$2 == "") { printf "   • %s\n", $$1 } else { printf "   • %-20s %s\n", $$1, $$2 } }' $(MAKEFILE_LIST)

.PHONY: up
up: ## 🐘 Starts PostgreSQL container on port 5440
	docker compose up -d

.PHONY: build
build: ## 🏗️  Builds the Ceeyu application from main.go
	go build -o ./bin/main ./cmd/main.go

.PHONY: test
test: ## 🧪 Runs all tests in the project once without caching
	go test ./... --count 1 --short

