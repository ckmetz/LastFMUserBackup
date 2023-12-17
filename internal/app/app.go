package app

import (
	"fmt"
	"time"
)

type app struct {
	ErrShutdown         error
	LastFMClientService LastFMClient
	CsvWriterService    Writer
}

func NewApp(
	client LastFMClient,
	writer Writer,
) *app {
	return &app{
		LastFMClientService: client,
		CsvWriterService:    writer,
		ErrShutdown:         fmt.Errorf("application was shutdown gracefully"),
	}
}

func (a app) Start(params AppConfig) error {
	pageNumber := 1
	albumResults, err := a.LastFMClientService.GetUsersTopAlbums(params.Username, 50, pageNumber)
	if err != nil {
		println("Error getting albums results: ", err.Error())
		return err
	}

	for len(albumResults.Albums) == 50 {
		for _, album := range albumResults.Albums {
			a.CsvWriterService.WriteLastFMRecords(album.Artist.Name, album.Name, album.PlayCount)

		}

		pageNumber++
		albumResults, err = a.LastFMClientService.GetUsersTopAlbums(params.Username, 50, pageNumber)
		if err != nil {
			println("Error getting albums results: ", err.Error())
			return err
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
