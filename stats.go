package main

import (
	"log"
	"os"
)

func main() {
	if err := Run(os.Args); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
