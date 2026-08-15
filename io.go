package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func ReadCSVFile(path string) ([][]string, error) {
	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	fi, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fi.Close()

	lines, err := csv.NewReader(fi).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
