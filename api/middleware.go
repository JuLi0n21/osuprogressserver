package api

import "github.com/gofiber/fiber/v2"

var userSessions = map[string]string{}

func sessionChecker(c *fiber.Ctx) error {

	if c.Path() == "/" {
		c.Next()
		return nil
	}

	sessionID := c.Cookies("session")
	userid, ok := userSessions[sessionID]
	if !ok {
		c.Redirect("/", fiber.StatusFound)
	}

	c.Locals("userid", userid)
	c.Next()
	return nil
}
