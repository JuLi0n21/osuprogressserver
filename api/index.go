package api

import (
	"context"
	"fmt"
	"net/http"
	"osuprogressserver/views"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	//todo

	component := views.Index()

	component.Render(context.Background(), w)
	fmt.Fprintf(w, "Hello world")
}
