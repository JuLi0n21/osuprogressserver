package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	client "osuprogressserver/OsuApi"
	"osuprogressserver/types"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) Oauth(c *fiber.Ctx) error {

	code := c.Query("code")
	state := c.Query("state")

	clientsecret := os.Getenv("CLIENT_SECRET")
	clientid := os.Getenv("CLIENT_ID")

	CookieID := c.Cookies("session")
	user, ok := UserSessions[CookieID]

	if !ok {
		return errors.New("Session not found")
	}

	if user.cookieid == state {
		fmt.Println("Sign in Success, Requesting Accesstoken")

		body := url.Values{
			"client_id":     {clientid},
			"client_secret": {clientsecret},
			"code":          {code},
			"grant_type":    {"authorization_code"},
		}

		req, err := http.NewRequest("POST", "https://osu.ppy.sh/oauth/token", bytes.NewBufferString(body.Encode()))
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		httpclient := &http.Client{
			Timeout: time.Minute,
		}

		res, err := httpclient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
			return errors.New("Authorization Request Failed!")
		}

		fmt.Println("Token Rechived")

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		authUser := types.AuthUser{
			TimeStamp: time.Now(),
		}

		err = json.Unmarshal(data, &authUser)
		if err != nil {
			return err
		}

		user := types.User{
			Auth: authUser,
		}

		fmt.Println("Create new Api client")

		apiclient, err := client.NewOsuApiClient(user)
		if err != nil {
			return err
		}

		fmt.Println("Requesting Userdata")

		apidata, err := apiclient.Me()
		if err != nil {
			return err
		}

		newuser := types.User{
			UserId:      apidata.ID,
			Username:    apidata.Username,
			Banner:      apidata.CoverURL,
			Avatar:      apidata.AvatarURL,
			GlobalRank:  strconv.Itoa(apidata.Statistics.Rank.Global),
			LocalRank:   strconv.Itoa(apidata.Statistics.Rank.Country),
			Country:     apidata.Country.Name,
			Countrycode: apidata.Country.Code,
			Mode:        apidata.Playmode,
			Auth:        user.Auth,
		}

		s.store.SaveUser(newuser)
		if err != nil {
			return err
		}

		UserSessions[CookieID] = UserContext{
			cookieid: CookieID,
			User:     newuser,
		}
	}

	c.Redirect("/")

	return nil
}

func (s *Server) OauthAccess(c *fiber.Ctx) error {
	fmt.Print(c)

	return nil
}
