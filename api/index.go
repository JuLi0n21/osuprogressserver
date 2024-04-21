package api

import (
	"context"
	"osuprogressserver/cmp"
	"osuprogressserver/types"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Index(c *fiber.Ctx) error {

	scores, err := s.store.GetRandomScores(10)
	if err != nil {
		return nil
	}

	component := cmp.View_Index(scores)

	ctx := context.WithValue(context.Background(), "player", c.Locals("User").(types.UserContext))

	handler := adaptor.HTTPHandler(C(templ.Handler(component), ctx))

	return handler(c)
}
