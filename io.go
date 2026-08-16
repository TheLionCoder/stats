package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	csvDelimiter      = '|'
	expectedCSVFields = 4
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

func ReadCSVFile(path string) ([]Record, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open CSV file %q: %w", path, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = expectedCSVFields
	reader.ReuseRecord = true

	records := make([]Record, 0)
	rowNumber := 0

	for {
		row, err := reader.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		rowNumber++

		if err != nil {
			return nil, fmt.Errorf(
				"read CSV row from %q: %w",
				rowNumber,
				path,
				err,
			)
		}

		records = append(records, Record{
			Name:       row[0],
			Surname:    row[1],
			Number:     row[2],
			LastAccess: row[3],
		})
	}

	return records, nil
}

func SaveFile(path string, data []Record) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create CSV file %q: %w", path, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf(
				"close CSV file %q: %w",
				path,
				closeErr,
			)
		}
	}()

	writer := csv.NewWriter(file)
	writer.Comma = csvDelimiter

	for i, record := range data {
		row := []string{
			record.Name,
			record.Surname,
			record.Number,
			record.LastAccess,
		}

		if writerErr := writer.Write(row); writerErr != nil {
			return fmt.Errorf(
				"write record %d to %q: %w",
				i+1,
				path,
				writerErr,
			)
		}
	}

	writer.Flush()

	if writerErr := writer.Error(); writerErr != nil {
		return fmt.Errorf(
			"flush CSV file %q: %w",
			path,
			writerErr,
		)
	}

	return nil
}
