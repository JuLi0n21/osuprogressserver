package api

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserContext struct {
	userid     int
	sessoionid string
}

var userSessions = map[string]UserContext{}

func SessionChecker(c *fiber.Ctx) error {

	if c.Path() == "/" {
		return c.Next()
	}

	sessionID := c.Cookies("session")
	user, ok := userSessions[sessionID]

	if !ok {
		user = UserContext{
			userid:     -1,
			sessoionid: generateUserID(),
		}

		userSessions[sessionID] = user
		fmt.Println(user)
		c.Cookie(&fiber.Cookie{
			Name:  "session",
			Value: user.sessoionid,
		})
	}

	_ = user.userid

	c.Locals("userid", user.userid)

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

func C(next http.Handler, context context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
