package app

import "github.com/shkh/lastfm-go/lastfm"

type Writer interface {
	WriteLastFMRecords(artist string, album string, scrobbles string) error
}

type LastFMClient interface {
	GetUsersTopAlbums(user string, limit int, page int) (lastfm.UserGetTopAlbums, error)
}
