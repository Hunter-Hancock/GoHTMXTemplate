build:
	@go build -o bin/dbproject ./cmd
run: build
	@./bin/dbproject
test:
	@go test -v ./...
