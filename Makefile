build:
	@air 

style:
	@npx tailwindcss -i ./static/input.css -o ./static/style.css --watch

run:
	@go build .
	@go run .

setup:
	go mod download

prod:
	npx tailwindcss -i ./static/input.css -o ./static/style.css && templ generate && go build .
