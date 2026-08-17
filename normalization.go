package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Normalize(data []float64, mean, stdev float64) ([]float64, error) {
	if math.IsNaN(mean) || math.IsInf(mean, 0) {
		return nil, fmt.Errorf("mean must be a finite number")
	}

	if stdev <= 0 || math.IsNaN(stdev) || math.IsInf(stdev, 0) {
		return nil, fmt.Errorf(
			"standard deviation must be a finite number greater than zero")
	}

	normalized := make([]float64, len(data))

	for i, value := range data {
		normalized[i] = (value - mean) / stdev
	}

	return normalized, nil

}

func RandomFloat(minValue, maxValue float64) (float64, error) {
	if math.IsNaN(minValue) ||
		math.IsNaN(maxValue) ||
		math.IsInf(minValue, 0) ||
		math.IsInf(maxValue, 0) {
		return 0, fmt.Errorf("min and max must be finite numbers")
	}

	if minValue > maxValue {
		return 0, fmt.Errorf(
			"min value %.4f cannot be greater than max value %.4f",
			minValue, maxValue,
		)
	}

	if minValue == maxValue {
		return minValue, nil
	}

	return minValue + rand.Float64()*(maxValue-minValue), nil

}
