package api

import (
	"fmt"
	"osuprogressserver/types"
	"osuprogressserver/views"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) ScorePage(c *fiber.Ctx) error {

	scoreid := c.Params("id")

	fmt.Println(scoreid)

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

	component := views.ScoreSite(player, scores)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
