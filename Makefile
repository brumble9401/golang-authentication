build:
	@go build -o bin/golang-authentication cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/golang-authentication

migrate-up:
	@migrate -database postgres://postgres:cuong9401@0.0.0.0:5432/golang-authentication?sslmode=disable -path db/migration up