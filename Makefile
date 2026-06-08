.PHONY: help build run test vet fmt tidy clean exercises syntax dsa patterns

# Default target: list available commands.
help:
	@echo "learn-golang — make targets:"
	@echo "  make build      build all packages"
	@echo "  make test       run all tests"
	@echo "  make vet        run go vet"
	@echo "  make fmt        format all code"
	@echo "  make tidy       tidy go.mod"
	@echo "  make clean      remove build artifacts"
	@echo ""
	@echo "  make run        run the root index"
	@echo "  make exercises  run leetcode + codewars"
	@echo "  make syntax     run syntax drills"
	@echo "  make dsa        run dsa practice"
	@echo "  make patterns   run concurrency patterns"

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

clean:
	go clean ./...
	rm -rf bin

run:
	go run .

exercises:
	go run ./cmd/exercises

syntax:
	go run ./cmd/syntax

dsa:
	go run ./cmd/dsa

patterns:
	go run ./cmd/patterns
