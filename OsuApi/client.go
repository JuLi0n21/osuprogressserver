package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"osuprogressserver/types"
	"time"
)

const (
	OsuApiUrl = "https://osu.ppy.sh/api/v2"
)

var clientid = os.Getenv("CLIENT_ID")
var clientsecret = os.Getenv("CLIENT_SECRET")
var redirect_uri = os.Getenv("REDIRECT_URI")
var scopes = []string{
	"public",
	"identify",
}

type OsuApiClient struct {
	User       types.User
	BaseURL    string
	HTTPclient *http.Client
}

func NewOsuApiClient(User types.User) (*OsuApiClient, error) {

	if User.Auth == (types.AuthUser{}) {
		return nil, errors.New("invalid client instaciaction")
	}

	if time.Now().After(User.Auth.TimeStamp.Add(time.Duration(User.Auth.ExpiresIn))) {
		//refresh token
	}

	return &OsuApiClient{
		User:    User,
		BaseURL: OsuApiUrl,
		HTTPclient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func (c *OsuApiClient) RefreshToken() error {

	return nil
}

func (c *OsuApiClient) Me() (*types.ApiUser, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/me/%s", OsuApiUrl, c.User.Mode), nil)
	if err != nil {
		return nil, err
	}

	res := types.ApiUser{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *OsuApiClient) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.User.Auth.AccessToken))

	res, err := c.HTTPclient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
