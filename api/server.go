package api

import (
	"net/http"
	"osuprogressserver/storage"
)

type Server struct {
	port  string
	store storage.Storage
}

func NewServer(port string, store storage.Storage) *Server {
	return &Server{
		port:  port,
		store: store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("GET /", s.Index)
	http.HandleFunc("GET /style.css", s.style)
	http.HandleFunc("GET /favicon.ico", s.icon)

	return http.ListenAndServe(s.port, nil)
}

func (s *Server) style(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./style.css")
}

func (s *Server) icon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}
