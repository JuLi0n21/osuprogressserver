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

	return http.ListenAndServe(s.port, nil)
}
