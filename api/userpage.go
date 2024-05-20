package api

import (
	"osuprogressserver/cmp"
	"osuprogressserver/types"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Userpage(c *fiber.Ctx) error {
	//todo

	id := c.Params("id")

	player := c.UserContext().Value("player").(types.UserContext)
	user := types.UserContext{}
	if c.Path() == "/me" {
		if player.User.UserId == 0 {
			return c.Redirect("/login")
		}
		user = player

	} else {

		uid, err := strconv.Atoi(id)
		if err != nil {
			return c.SendStatus(404)
		}

		apiu, err := s.store.GetApiUser(uid)
		if err != nil {
			return c.SendStatus(404)
		}

		user.ApiUser = apiu
		user.User = types.User{
			Username: apiu.Username,
			UserId:   apiu.ID,
		}

	}

	stats := types.Stats{
		Time:   "9h",
		Status: "Afk",
		Count:  8234,
		Screen: "Mainmenu",
	}

	component := cmp.View_Userpage(user.ApiUser, stats)

	handler := adaptor.HTTPHandler(C(templ.Handler(component), c.UserContext()))

	return handler(c)
}
