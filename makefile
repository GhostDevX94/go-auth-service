include .env

BINARY_NAME := myapp
BINARY_PATH := bin/$(BINARY_NAME)

build:
	go build -o $(BINARY_PATH) cmd/api/main.go

clean:
	rm -f $(BINARY_PATH)

run: build
	@echo "Start $(BINARY_PATH)..."
	@$(BINARY_PATH) &

stop:
	@echo "Stop $(BINARY_NAME)..."
	@pkill -f "$(BINARY_PATH)" || true

restart: stop build run