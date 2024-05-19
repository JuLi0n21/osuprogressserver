package api

import (
	"context"
	"fmt"
	"osuprogressserver/cmp"
	"osuprogressserver/types"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Userpage(c *fiber.Ctx) error {
	//todo
	player := c.Locals("User").(types.UserContext)
	fmt.Println(player.User.Username)

	stats := types.Stats{
		Time:   "9h",
		Status: "Afk",
		Count:  8234,
		Screen: "Mainmenu",
	}

	ctx := context.WithValue(context.Background(), "player", player)

	component := cmp.View_Userpage(stats)

	handler := adaptor.HTTPHandler(C(templ.Handler(component), ctx))

	return handler(c)
}
