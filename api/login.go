package api

import (
	"fmt"
	"os"
	"osuprogressserver/cmp"
	"osuprogressserver/types"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var clientsecret = os.Getenv("CLIENT_SECRET")
var scopes = []string{
	"public",
	"identify",
}

func (s *Server) Login(c *fiber.Ctx) error {

	fmt.Println(c.Locals("User").(types.UserContext).User.Username)

	if c.Locals("User").(types.UserContext).User.UserId != 0 {
		return c.Redirect("/me")
	}

	scope := strings.Join(scopes, " ")
	var redirect_uri = os.Getenv("REDIRECT_URI")
	var clientid = os.Getenv("CLIENT_ID")

	CookieID := c.Cookies("session")

	user, err := UserSessions.Read(CookieID)

	if err != nil {
		return c.Redirect("/")
	}

	component := cmp.Login(clientid, redirect_uri, scope, user.Cookieid)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
