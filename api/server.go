package api

import (
	"osuprogressserver/storage"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

type Server struct {
	port   string
	store  storage.Storage
	secure bool
}

func NewServer(port string, store storage.Storage, env string) *Server {
	secure := false
	if env == "prod" {
		secure = true
	}
	return &Server{
		port:   port,
		store:  store,
		secure: secure,
	}
}

func (s *Server) Start() error {
	app := fiber.New(fiber.Config{
		CaseSensitive: false,
	})

	app.Get("/metrics", monitor.New(monitor.Config{
		Title: "Osu!Progress",
	}))
	prometheus := fiberprometheus.New("Osu!Progress")
	prometheus.RegisterAt(app, "/prometheus")
	app.Use(prometheus.Middleware)

	app.Use(logger.New(), pprof.New())

	if s.secure {
		app.Use(redirectToHTTPS)
	}

	app.Use(CookieClicker, Styler)
	app.Static("assets", "./static")

	api := app.Group("/api")
	api.Get("/scoresearch/*", s.ScoreSearch)
	api.Get("/score", s.Score)
	api.Post("/score", s.Score)
	api.Get("/player", s.PlayerIcon)

	app.Get("/login", s.Login)

	oauth := app.Group("/oauth")
	oauth.Get("/code", s.Oauth)
	oauth.Get("/token", s.OauthAccess)

	app.Get("/me", s.Userpage)
	app.Use(cache.New())
	app.Get("/", s.Index)
	app.Get("/users/:id", s.Userpage)
	app.Get("/score/:id", s.ScorePage)

	return app.Listen(s.port)
}
