package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Println("csvData input output!")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := ReadCSVFile(input)
	var data = []Record{}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}

		data = append(data, temp)
		log.Println(temp)
	}

	err = SaveFile(output, data)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
