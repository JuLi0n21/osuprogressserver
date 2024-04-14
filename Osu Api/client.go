package client

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"osuprogressserver/types"
	"strings"
	"time"
)

const (
	OsuApiUrl = "https://osu.ppy.sh/api/2/"
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

func NewOsuApiClient(User types.User) *OsuApiClient {

	if User.Auth == (types.AuthUser{}) {
		return nil
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
	}
}

func requestAuthCode(scopes []string, state string) error {

	baseURL := "https://osu.ppy.sh/api/2/"
	scope := strings.Join(scopes, " ")

	req, err := http.NewRequest("Get", fmt.Sprintf("%s/oauth/authorize?client_id=%s&redirect_id=%s&response_type=code&scopes=%s&state=%s", baseURL, clientid, redirect_uri, scope, state), nil)
	if err != nil {
		return err
	}

	httpclient := &http.Client{
		Timeout: time.Minute,
	}

	res, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.New("Authorization Request Failed!")
	}

	return nil
}

func (c *OsuApiClient) accesstoken(code int, redirect_uri string) (types.AuthUser, error) {
	grant_type := ""

	_ = grant_type

	return types.AuthUser{}, nil
}
