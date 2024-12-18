# Makefile for Persons Project

# Variables
BINARY_NAME=persons
SRC_DIR=.

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) $(SRC_DIR)/main.go $(SRC_DIR)/persons.go

# Run the project
run: build
	@echo "Running the project..."
	./$(BINARY_NAME) # run the binary
	rm -f $(BINARY_NAME) # Clean up build files

# Run tests
test:
	@echo "Running tests..."
	go test ./...

