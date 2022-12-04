package main

import "math/rand"

var asciiRepresentation = map[string]string{
	"rain": "/",
	"snow": "*"}

var density = map[string]float64{
	"rain": rainDensity,
	"snow": snowDensity}

func drawChar(weather string) string {
	if rand.Float64() < density[weather] {
		return asciiRepresentation[weather]
	}
	return " "
}
