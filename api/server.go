package api

import (
	"osuprogressserver/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(
		logger.New(),
	)

	api := app.Group("/api")
	api.Get("/scoresearch/*", s.ScoreSearch)

	oauth := app.Group("/oauth")
	oauth.Get("/code", s.Oauth)
	oauth.Get("/token", s.OauthAccess)

	app.Get("/", s.Index)
	app.Static("assets", "./static")

	app.Use(SessionChecker)

	app.Get("/score/:id", s.ScorePage)
	app.Get("/login", s.Login)

	return app.Listen(s.port)
}
