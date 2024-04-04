package api

import (
	"osuprogressserver/storage"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	//app.Use(SessionChecker)

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	api := app.Group("/api")

	api.Get("/scoresearch/*", s.ScoreSearch)

	app.Static("assets", "./static")
	app.Get("/", s.Index)
	app.Get("/score/:id", s.ScorePage)

	return app.Listen(s.port)
}
