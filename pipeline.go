package main

import (
	"cmp"
	"fmt"
	"math"
	"path/filepath"
	"slices"
)

type DataFile struct {
	Filename string
	Count    int
	Minimun  float64
	Maximun  float64
	Mean     float64
	StdDev   float64
}

func CalculateStats(records []Record) (DataFile, error) {
	if len(records) == 0 {
		return DataFile{}, fmt.Errorf("cannot calculate statistics for empty data")
	}

	minimum := records[0].Salary
	maximun := records[0].Salary

	var mean float64
	var m2 float64

	for i, record := range records {
		value := record.Salary
		if value < minimum {
			minimum = value
		}

		if value > maximun {
			maximun = value
		}

		// Welford's algorithm
		count := float64(i + 1)
		delta := value - mean
		mean += delta / count

		delta2 := value - mean
		m2 += delta * delta2
	}

	variance := m2 / float64(len(records))
	stdDev := math.Sqrt(variance)

	return DataFile{
		Count:   len(records),
		Minimun: minimum,
		Maximun: maximun,
		Mean:    mean,
		StdDev:  stdDev,
	}, nil

}

func Run(args []string) error {
	if len(args) < 2 {
		program := "stats"

		if len(args) > 0 {
			program = filepath.Base(args[0])
		}

		return fmt.Errorf(
			"usage: %s <input.csv> <output.csv>",
			program,
		)
	}

	files := make([]DataFile, 0, len(args)-1)

	for _, filename := range args[1:] {
		records, err := ReadCSVFile(filename)
		if err != nil {
			return fmt.Errorf("read %q: %w", filename, err)
		}

		stats, err := CalculateStats(records)
		if err != nil {
			return fmt.Errorf("process %q: %w", filename, err)
		}

		stats.Filename = filename
		files = append(files, stats)
	}

	slices.SortFunc(files, func(a, b DataFile) int {
		return cmp.Compare(a.Mean, b.Mean)
	})

	for _, file := range files {
		fmt.Printf(
			"%s: count=%d mean=%.4f stdev=%.4f min=%.4f max=%.4f\n",
			file.Filename,
			file.Count,
			file.Mean,
			file.StdDev,
			file.Minimun,
			file.Maximun,
		)
	}

	return nil

}
