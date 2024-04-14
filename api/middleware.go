package api

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"osuprogressserver/types"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserContext struct {
	User     types.User
	cookieid string
}

var UserSessions = map[string]UserContext{}

func SessionChecker(c *fiber.Ctx) error {

	CookieID := c.Cookies("session")
	user, ok := UserSessions[CookieID]

	if !ok {
		user = UserContext{
			User:     types.User{},
			cookieid: generateUserID(),
		}

		UserSessions[CookieID] = user
		c.Cookie(&fiber.Cookie{
			Name:     "session",
			Value:    user.cookieid,
			HTTPOnly: true,
			Expires:  time.Now().AddDate(1, 0, 0),
			SameSite: "/",
		})
	}

	fmt.Println(user.User.UserId)
	if user.User.UserId == 0 && c.Path() != "/login" {

		return c.Redirect("/login")
	}

	c.Locals("userid", user.User.UserId)

	return c.Next()
}

func generateUserID() string {
	return "session_" + randString(64)
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
