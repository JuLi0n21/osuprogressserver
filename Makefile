build:
	@./tailwind.exe -o style.css --watch & templ generate -watch & air

run:
	@go build .
	@go run .

setup:
	go mod download
