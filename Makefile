# Variables
GO          := go
BINARY      := pluto
SRC_DIR     := cmd
SRC         := $(wildcard $(SRC_DIR)/*.go)
LDFLAGS     := -ldflags="-s -w"
TEST_DIR    := tests/unit
COVERAGE    := coverage.out

# Build the application
build:
	$(GO) build $(LDFLAGS) -o $(BINARY) $(SRC)

# Clean build artifacts
clean:
	rm -f $(BINARY)

# Run the application
run: clean build
	./$(BINARY)

# Run tests
test:
	$(GO) test -v ./$(TEST_DIR)/...

# Format code
format:
	$(GO) fmt ./...

# Lint code
lint:
	golint ./...

# Vet code
vet:
	$(GO) vet ./...

# Generate code coverage report
coverage:
	$(GO) test -coverprofile=$(COVERAGE) ./$(TEST_DIR)/...
	$(GO) tool cover -func=$(COVERAGE)

.PHONY: build clean run test format lint vet coverage
