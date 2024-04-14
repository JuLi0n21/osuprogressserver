package api

import (
	"os"
	"osuprogressserver/cmp"
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

	scope := strings.Join(scopes, " ")
	var redirect_uri = os.Getenv("REDIRECT_URI")
	var clientid = os.Getenv("CLIENT_ID")

	CookieID := c.Cookies("session")
	user, ok := UserSessions[CookieID]

	if !ok {
		return nil
	}

	component := cmp.Login(clientid, redirect_uri, scope, user.cookieid)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
