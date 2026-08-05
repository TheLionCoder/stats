package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	var initialized = 0
	nValues := 0
	var sum float64

	for i:= 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[1], 64)
		if err != nil {
			continue
		}
	}
}
