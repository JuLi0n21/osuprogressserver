package api

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var userSessions = map[string]string{}

func SessionChecker(c *fiber.Ctx) error {

	if c.Path() == "/" {
		return c.Next()
	}

	sessionID := c.Cookies("session")
	userid, ok := userSessions[sessionID]

	if !ok {
		userID := generateUserID()

		userSessions[sessionID] = userID
		fmt.Println(userID)
		c.Cookie(&fiber.Cookie{
			Name:  "session",
			Value: userID,
		})
	}

	_ = userid

	c.Locals("userid", 14100399)

	return c.Next()
}

func generateUserID() string {
	return "user_" + randString(64)
}

func randString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
