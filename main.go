package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Need 3 arguments: weather type, width, and height")
	}

	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	weatherType := os.Args[1]

	if weatherType == "rain" {
		rain(width, height)
	}
	if weatherType == "snow" {
		snow(width, height)
	}
}
