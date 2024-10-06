# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOFMT = $(GOCMD) fmt
GOINSTALL = $(GOCMD) install
GOCLEAN = $(GOCMD) clean
GOMOD = $(GOCMD) mod tidy
BINARY_NAME = tracesync
BINARY_UNIX = $(BINARY_NAME)_unix
BUILD_DIR = build

# Directories
CMD_DIR = ./cmd
TEST_DIR = ./tests

.PHONY: all test build fmt clean mod install run help

all: test build

# Build the binary
build:
	@echo "Building the binary..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Binary built at $(BUILD_DIR)/$(BINARY_NAME)"

# Run tests
test:
	@echo "Running unit tests..."
	$(GOTEST) -v $(TEST_DIR)/unit

# Format the code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Install dependencies
mod:
	@echo "Tidying Go modules..."
	$(GOMOD)

# Install the binary globally
install:
	@echo "Installing the binary globally..."
	$(GOINSTALL) main.go

# Run the binary
run:
	@echo "Running TraceSync..."
	$(GOCMD) run main.go

# Display help
help:
	@echo "Usage:"
	@echo "  make build     - Build the binary"
	@echo "  make test      - Run unit tests"
	@echo "  make fmt       - Format the code"
	@echo "  make clean     - Clean up the build artifacts"
	@echo "  make mod       - Tidy Go modules"
	@echo "  make install   - Install the binary globally"
	@echo "  make run       - Run the TraceSync CLI"

