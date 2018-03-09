default: build

build:
	go fmt
	go vet
	go build

test: build
	go test ./...

coverage-test:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out