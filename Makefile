BINARY_NAME=bot-keren
MAIN_PATH=cmd/main/main.go

.PHONY: all build run clean tidy

all: build

build:
	@echo "Membangun binary..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

run:
	@echo "Menjalankan bot..."
	go run $(MAIN_PATH)

clean:
	@echo "Membersihkan..."
	rm -rf bin/
	go clean

tidy:
	@echo "Merapikan dependencies..."
	go mod tidy
