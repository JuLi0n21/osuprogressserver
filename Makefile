build:
	@air & templ generate -watch

style:
	@npx tailwindcss -i input.css -o style.css --watch

run:
	@go build .
	@go run .

setup:
	go mod download
