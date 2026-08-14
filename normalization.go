package main

import (
	"math"
	"math/rand"
)

func normalize(data []float64, mean, stdev float64) []float64 {
	if stdev == 0 {
		return data
	}

	normalized := make([]float64, len(data))

	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdev*10000) / 10000
	}

	return normalized

}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
