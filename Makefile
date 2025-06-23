include .env

BINARY_NAME := myapp

BINARY_PATH := bin/$(BINARY_NAME)

MIGRATE := /opt/homebrew/bin/migrate

DB_URL := $(DATABASE_URL)

MIGRATIONS_DIR := migrations


build:
	@echo "Putting the application together..."
	go build -o $(BINARY_PATH) cmd/api/main.go

clean:
	@echo "Removing binary file..."
	rm -f $(BINARY_PATH)

run: build
	@echo "Запускаем $(BINARY_PATH)..."
	@$(BINARY_PATH) &

stop:
	@echo "Stopping $(BINARY_NAME)..."
	@pkill -f "$(BINARY_PATH)" || true

restart: stop build run
	@echo "Restart completed."

create-migration:
	@read -p "Enter a migration name: " NAME; \
	if [ -z "$$NAME" ]; then echo "Migration name must not be empty!"; exit 1; fi; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq "$$NAME"

migrate-up:
	@echo "Applying new migrations..."
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) up

migrate-down:
	@echo "Rolling back the last migration..."
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) down 1

migrate-steps:
	@read -p "Enter the number of steps:" STEPS; \
	if [ -z "$$STEPS" ]; then echo "Шаги не указаны!"; exit 1; fi; \
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) up $$STEPS

migrate-rollback:
	@read -p "Enter the number of steps to roll back: " STEPS; \
	if [ -z "$$STEPS" ]; then echo "Шаги не указаны!"; exit 1; fi; \
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) down $$STEPS
