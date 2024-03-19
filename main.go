package main

import (
	"context"
	"net/http"
	"osuprogressserver/templates"
)

func Home(w http.ResponseWriter, req *http.Request) {

	headers := req.Header

	component := templates.Home(Count)

	if ok := headers[http.CanonicalHeaderKey("Hx-request")]; len(ok) > 0 {
		w.Header().Add("HX-Push-Url", "/")
		w.Header().Add("HX-Replace-Url", "/")
		component.Render(context.Background(), w)
	} else {
		layout := templates.Layout(component)
		layout.Render(context.Background(), w)
	}

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

func About(w http.ResponseWriter, req *http.Request) {

	headers := req.Header

	component := templates.About()

	if ok := headers[http.CanonicalHeaderKey("Hx-request")]; len(ok) > 0 {

		w.Header().Set("HX-Push-Url", "/about")
		w.Header().Set("HX-Replace-Url", "/about")
		component.Render(context.Background(), w)
	} else {
		layout := templates.Layout(component)
		layout.Render(context.Background(), w)
	}
}

func Contact(w http.ResponseWriter, req *http.Request) {

	headers := req.Header

	component := templates.Contact()

	if ok := headers[http.CanonicalHeaderKey("Hx-request")]; len(ok) > 0 {
		//w.Header().Set("HX-Push-Url", "/contact")
		//w.Header().Set("HX-Replace-Url", "/contact")
		component.Render(context.Background(), w)
	} else {
		layout := templates.Layout(component)
		layout.Render(context.Background(), w)
	}

}

var Count = 0

func main() {

	r := http.NewServeMux()

	r.HandleFunc("GET /", Home)
	r.HandleFunc("GET /about", About)
	r.HandleFunc("GET /contact", Contact)

	r.HandleFunc("POST /api/count", Api)

	r.HandleFunc("DELETE /api/count", Reset)

	http.ListenAndServe(":8080", r)
}
