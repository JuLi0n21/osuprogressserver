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

	player := types.User{
		Username:    "JuLi0n_",
		UserId:      14100399,
		Banner:      "https://assets.ppy.sh/user-profile-covers/14100399/cd1600936d7a56115cd147f47169addcf8f812133861b667c7dd3d177ca5068d.jpeg",
		Avatar:      "https://a.ppy.sh/14100399?1672009368.jpeg",
		GlobalRank:  "57497",
		LocalRank:   "2928",
		Country:     "Germany",
		Countrycode: "de",
	}

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

	component := cmp.View_ScoreSite(player, scores)

	component.Render(ctx, c.Context().Request.BodyWriter())

	handler := adaptor.HTTPHandler(C(templ.Handler(component), ctx))

	return handler(c)
}
