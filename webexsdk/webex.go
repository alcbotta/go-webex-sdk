package webexsdk

import (
	"errors"
	"net/url"
)

type Webex struct {
	Rooms *Room
}

var webexTokenID string
var baseURL *url.URL

func InitWebex(tokenID string) (*Webex, error) {
	if tokenID == "" {
		return nil, errors.New("Token ID shold not be empty")
	}
	u, err := url.Parse("https://api.ciscospark.com/v1")
	if err != nil {
		return nil, err
	}
	webexTokenID = tokenID
	baseURL = u
	webex := &Webex{}
	webex.Rooms = &Room{}
	return webex, nil
}
