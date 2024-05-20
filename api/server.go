package api

import (
	"osuprogressserver/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	//	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app := fiber.New(fiber.Config{
		CaseSensitive: false,
	})

	//prometheus := fiberprometheus.New("Osu!Progress")
	//prometheus.RegisterAt(app, "/metrics")
	//app.Use(prometheus.Middleware)

	app.Use(logger.New(), pprof.New())

	app.Get("/metrics", monitor.New())

	app.Use(CookieClicker)
	app.Static("assets", "./static")

	api := app.Group("/api")
	api.Get("/scoresearch/*", s.ScoreSearch)
	api.Get("/score", s.Score)
	api.Post("/score", s.Score)

	app.Get("/", s.Index)

	app.Get("/login", s.Login)

	oauth := app.Group("/oauth")
	oauth.Get("/code", s.Oauth)
	oauth.Get("/token", s.OauthAccess)

	//app.Use(Authorization)
	app.Get("/me", s.Userpage)
	app.Get("/users/:id", s.Userpage)
	app.Get("/score/:id", s.ScorePage)

	return app.Listen(s.port)
}
