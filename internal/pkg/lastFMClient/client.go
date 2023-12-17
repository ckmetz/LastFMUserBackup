package lastFM

import (
	"errors"
	"fmt"

	"github.com/shkh/lastfm-go/lastfm"
)

type LastFMClient struct {
	api lastfm.Api
}

func NewLastFMClient(apiKey string, apiSecret string, username string, password string) (*LastFMClient, error) {
	api := lastfm.New(apiKey, apiSecret)

	err := api.Login(username, password)
	if err != nil {
		fmt.Println(err)
		return &LastFMClient{}, errors.New("LastFM login error")
	}

	return &LastFMClient{
		api: *api,
	}, nil
}

func (l LastFMClient) GetUsersTopAlbums(user string, limit int, page int) (lastfm.UserGetTopAlbums, error) {
	result, err := l.api.User.GetTopAlbums(lastfm.P{"user": user, "limit": limit, "page": page})
	if err != nil {
		println(err.Error())
	}

	return result, nil
}
