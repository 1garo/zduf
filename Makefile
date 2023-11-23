up:
	@go run cmd/api/main.go

test:
	@go test -v ./...

build:
	@go build cmd/api/main.go
