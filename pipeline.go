package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func Run(args []string) error {
	if len(args) != 3 {
		program := "stats"

		if len(args) > 0 {
			program = filepath.Base(args[0])
		}

		return fmt.Errorf(
			"usage: %s <input.csv> <output.csv>",
			program,
		)
	}

	inputPath := args[1]
	outputPath := args[2]

	records, err := ReadCSVFile(inputPath)
	if err != nil {
		return fmt.Errorf("save output :%w", err)
	}

	if err := SaveFile(outputPath, records); err != nil {
		return fmt.Errorf("save output: %w", err)
	}

	log.Printf(
		"successfully processed %d records",
		len(records),
	)

	return nil
}
