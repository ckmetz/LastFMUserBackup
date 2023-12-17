package csvwriter

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

type Writer struct {
	file string
}

func NewWriter(fileToWrite string) (*Writer, error) {
	return &Writer{
		file: fileToWrite,
	}, nil
}

func (w *Writer) WriteLastFMRecords(artist string, album string, scrobbles string) error {
	f, err := os.OpenFile(w.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	record := []string{artist, album, scrobbles}
	err = writer.Write(record)
	if err != nil {
		log.Fatalln("error writing record to file", err)
		return errors.New(err.Error())
	}
	return nil
}
