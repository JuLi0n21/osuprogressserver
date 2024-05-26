package api

import (
	"osuprogressserver/cmp"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Index(c *fiber.Ctx) error {

	component := cmp.View_Index()

	handler := adaptor.HTTPHandler(C(templ.Handler(component), c.UserContext()))

	return handler(c)
}
