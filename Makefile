.PHONY: test test-coverage clean build

# Default target
all: test

# Run tests
test: clean
	go test -v ./...
	@make clean

# Run tests with race detection
test-race: clean
	go test -race -v ./...
	@make clean

# Run tests with coverage
test-coverage: clean
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	@make clean

# Clean build artifacts
clean:
	rm -f coverage.txt
	rm -rf bin/
	rm -f *.test
	rm -f *.out

# Build the application
build:
	go build -v ./...

# Build for multiple platforms
build-all:
	mkdir -p bin/linux-amd64 bin/darwin-amd64 bin/windows-amd64
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/main ./main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/main ./main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/main.exe ./main.go
