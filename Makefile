build:
	@go build -o bin/golang-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/golang-api