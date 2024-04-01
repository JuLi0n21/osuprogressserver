package api

import (
	"osuprogressserver/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) ScorePage(c *fiber.Ctx) error {

	component := views.ScoreSite()

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
