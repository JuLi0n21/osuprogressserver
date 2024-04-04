package api

import (
	"fmt"
	"osuprogressserver/types"
	"osuprogressserver/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Index(c *fiber.Ctx) error {
	//todo
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

	stats := types.Stats{
		Time:   "9h",
		Status: "Afk",
		Count:  8234,
		Screen: "Mainmenu",
	}

	scores, err := s.store.GetExtScore("", 14100399, 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	component := views.Index(player, stats, scores)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
