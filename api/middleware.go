package api

import (
	"context"
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

func CookieClicker(c *fiber.Ctx) error {

	CookieID := c.Cookies("session")
	user, ok := UserSessions[CookieID]

	if !ok {

		cookie := generateUserID(56)
		user = UserContext{
			User:     types.User{},
			cookieid: cookie,
		}

		UserSessions[cookie] = user
		c.Cookie(&fiber.Cookie{
			Name:     "session",
			Value:    user.cookieid,
			HTTPOnly: true,
			Expires:  time.Now().AddDate(1, 0, 0),
			SameSite: "/",
			Secure:   true,
		})
	}

	c.Locals("userid", user.User.UserId)

	return c.Next()
}

func Authorization(c *fiber.Ctx) error {

	uid := c.Locals("userid").(int)

	if uid == 0 {

		return c.Redirect("/login")
	}

	return c.Next()

}

func generateUserID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return "session_" + string(b)
}

func C(next http.Handler, context context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
