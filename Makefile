build:
	@go build -o bin/dbproject
run: build
	@./bin/dbproject
test:
	@go test -v ./...
