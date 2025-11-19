.PHONY: dev build run test clean

# Variabel
APP_NAME=api
MAIN_FILE=main.go
BUILD_DIR=build
AIR_CONFIG=.air.toml

# Development dengan hot reload menggunakan Air
dev:
	@echo "Menjalankan aplikasi dalam mode development dengan Air..."
	@if ! command -v air > /dev/null; then \
		echo "Air tidak ditemukan. Menginstal Air..."; \
		go install github.com/cosmtrek/air@latest; \
	fi
	air -c $(AIR_CONFIG)

# Build aplikasi
build:
	@echo "Membangun aplikasi..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Menjalankan aplikasi yang sudah di-build
run: build
	@echo "Menjalankan aplikasi..."
	./$(BUILD_DIR)/$(APP_NAME)

# Menjalankan test
test:
	@echo "Menjalankan test..."
	go test -v ./...

# Membersihkan file build
clean:
	@echo "Membersihkan file build..."
	@rm -rf $(BUILD_DIR)
	@go clean

# Menjalankan linter
lint:
	@echo "Menjalankan linter..."
	golangci-lint run

# Menjalankan formatter
fmt:
	@echo "Menjalankan formatter..."
	go fmt ./...

# Menjalankan semua tools (lint, fmt, test)
tools: fmt lint test

# Menampilkan bantuan
help:
	@echo "Perintah yang tersedia:"
	@echo "  make dev     - Menjalankan aplikasi dalam mode development dengan Air"
	@echo "  make build   - Membangun aplikasi"
	@echo "  make run     - Menjalankan aplikasi yang sudah di-build"
	@echo "  make test    - Menjalankan test"
	@echo "  make clean   - Membersihkan file build"
	@echo "  make lint    - Menjalankan linter"
	@echo "  make fmt     - Menjalankan formatter"
	@echo "  make tools   - Menjalankan semua tools (fmt, lint, test)"
	@echo "  make help    - Menampilkan bantuan ini" 