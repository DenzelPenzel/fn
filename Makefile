.PHONY: all test cover vet lint bench clean

all: lint test

test:
	go test -race -count=1 ./...

cover:
	go test -race -coverprofile=coverage.out -count=1 ./...
	go tool cover -html=coverage.out -o coverage.html

vet:
	go vet ./...

lint:
	golangci-lint run

bench:
	go test -bench=. -benchmem ./...

clean:
	rm -f coverage.out coverage.html
