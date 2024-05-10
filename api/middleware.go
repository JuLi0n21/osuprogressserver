package api

import (
	"context"
	"log/slog"
	"math/rand"
	"net/http"
	"osuprogressserver/types"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Requests",
			Help: "Total amount of Requests",
		},
		[]string{"method", "route", "status"},
	)
)

var UserSessions = map[string]types.UserContext{}

func CookieClicker(c *fiber.Ctx) error {

	CookieID := c.Cookies("session")
	user, ok := UserSessions[CookieID]

	if !ok {
		slog.Log(c.Context(), slog.LevelInfo, "New Session Created")
		cookie := generateUserID(56)
		user = types.UserContext{
			User:     types.User{},
			ApiUser:  types.ApiUser{},
			Cookieid: cookie,
		}

		UserSessions[cookie] = user
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

func Metrics(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		return err
	}

	requestsTotal.WithLabelValues(c.Method(), c.Path(), string(c.Response().Header.StatusMessage()))

	return nil
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
