package types

import "time"

type AuthUser struct {
	Tokentype    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TimeStamp    time.Time
}
