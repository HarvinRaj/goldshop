BIN_NAME=app

# Default target when `make` is run
all: build

# Build the Go binary
build:
	go build -o $(BIN_NAME) ./cmd/app

# Run the Go program from the root directory where the binary was built
run: build
	./$(BIN_NAME)

# Clean the build artifacts
clean:
	rm -f $(BIN_NAME)

# Test the Go program
test:
	go test ./...
