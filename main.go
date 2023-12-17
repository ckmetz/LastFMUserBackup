package main

import (
	"os"

	"lastFMUserHistoryBackup/internal/app"
	csvWriter "lastFMUserHistoryBackup/internal/pkg/csvWriter"
	lastFM "lastFMUserHistoryBackup/internal/pkg/lastFMClient"
)

// main runs the command-line parsing and validations. This function will also start the application logic execution.
func main() {
	appCfg, _ := app.LoadConfig("dev")
	client, err := lastFM.NewLastFMClient(appCfg.ApiKey, appCfg.ApiSecret, appCfg.Username, appCfg.Password)
	writer, err := csvWriter.NewWriter("scrobbleAlbums.csv")
	if err != nil {
		println("Shit fucked up")
	}
	app := app.NewApp(client, writer)

	// Run the App
	err = app.Start(*appCfg)
	if err != nil {
		// do stuff
		os.Exit(1)
	}
}
