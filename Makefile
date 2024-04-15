build:
	@go build -o bin/dbproject ./cmd
run: build
	@templ generate
	npx tailwindcss -i ./view/css/app.css -o ./view/assets/css/styles.css
	@./bin/dbproject
test:
	@go test -v ./...
