package main

import (
	"context"
	"net/http"
	"osuprogressserver/templates"
	"strconv"

	"github.com/a-h/templ"
)

func Home(w http.ResponseWriter, req *http.Request) {

	id := req.PathValue("id")

	c, err := strconv.Atoi(id)

	var component templ.Component

	if err != nil {
		component = templates.Hello(0)

	} else {
		Count = c
		component = templates.Hello(c)
	}
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

	r.HandleFunc("GET /{id...}", Home)

	r.HandleFunc("POST /api/count", Api)

	r.HandleFunc("DELETE /api/count", Reset)

	http.ListenAndServe(":8080", r)
}
