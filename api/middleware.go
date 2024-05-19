package api

import (
	"context"
	"log/slog"
	"math/rand"
	"net/http"
	"osuprogressserver/cmp"
	"osuprogressserver/types"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	UserSessions = types.NewUserSessionMap()
)

func CookieClicker(c *fiber.Ctx) error {

	CookieID := c.Cookies("session")

	user, err := UserSessions.Read(CookieID)

	if err != nil {
		slog.Log(c.Context(), slog.LevelInfo, "New Session Created")
		cookie := generateUserID(56)
		user = types.UserContext{
			User:     cmp.DefaultUser().User,
			ApiUser:  cmp.DefaultUser().ApiUser,
			Cookieid: cookie,
		}

		UserSessions.Write(CookieID, user)

		c.Cookie(&fiber.Cookie{
			Name:        "session",
			Value:       user.Cookieid,
			HTTPOnly:    true,
			Expires:     time.Now().AddDate(1, 0, 0),
			Path:        "/",
			Secure:      true,
			SessionOnly: false,
			SameSite:    "None",
		})
	}
	c.Locals("User", user)

	return c.Next()
}

func Authorization(c *fiber.Ctx) error {

	if c.Locals("User").(types.UserContext).User.UserId == 0 {
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
