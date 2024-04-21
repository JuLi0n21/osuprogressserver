package api

import (
	"context"
	"fmt"
	"osuprogressserver/cmp"
	"osuprogressserver/types"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) ScorePage(c *fiber.Ctx) error {

	scoreid := c.Params("id")

	sid, err := strconv.Atoi(scoreid)
	if err != nil {
		return c.SendStatus(404)
	}

	scores, err := s.store.GetExtScoreById(sid)
	if err != nil {
		fmt.Printf(err.Error())
		return c.SendStatus(404)
	}

	fmt.Println(c.Locals("User").(types.UserContext).User.Username)
	player := c.Locals("User").(types.UserContext)

	if len(scores) == 0 {
		return c.SendStatus(404)
	}

	themes := types.Theme{
		Dark:         "score-backdrop--dark",
		Medium_dark:  "score-backdrop--medium--dark",
		Medium:       "score-backdrop--medium",
		Medium_light: "score-backdrop--medium--light",
		Light:        "score-backdrop--light",
	}

	ctx := context.WithValue(context.Background(), "theme", themes)
	ctx = context.WithValue(ctx, "player", player)

	component := cmp.View_ScoreSite(scores)

	handler := adaptor.HTTPHandler(C(templ.Handler(component), ctx))

	return handler(c)
}
