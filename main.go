package main

import (
	"context"
	"net/http"
	"osuprogressserver/templates"
)

func Home(w http.ResponseWriter, req *http.Request) {
	component := templates.Hello("Hello")
	component.Render(context.Background(), w)
}

func Api(w http.ResponseWriter, req *http.Request) {

	Count++
	component := templates.Api(Count)
	component.Render(context.Background(), w)
}

func Reset(w http.ResponseWriter, req *http.Request) {

	Count = 0
	component := templates.Api(Count)
	component.Render(context.Background(), w)
}

var Count = 0

func main() {

	r := http.NewServeMux()

	r.HandleFunc("GET /", Home)

	r.HandleFunc("POST /api/count", Api)

	r.HandleFunc("DELETE /api/count", Reset)

	http.ListenAndServe(":8080", r)
}
