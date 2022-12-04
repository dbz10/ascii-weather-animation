package main

import (
	"fmt"
	"time"
)

const rainDensity = 0.05
const sleepMillis = 250

type rainPlane [][]string

func (p rainPlane) draw() {
	for _, row := range p {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}

func (p rainPlane) step() rainPlane {
	// everything down one and to the left one.
	// drop the bottom row and
	// instantiate a new random row at the top

	xShape := len(p[0])
	yShape := len(p)
	var newRow = make([]string, xShape)
	for i := 0; i < xShape; i++ {
		newRow[i] = drawChar("rain")
	}

	newPlane := rainPlane{}
	newPlane = append(newPlane, newRow)
	for i := 0; i < yShape-1; i++ {
		newPlane = append(newPlane, append(p[i][1:], drawChar("rain")))
	}

	return newPlane

}

func rain(width, height int) {
	// Initialize the screen
	screen := rainPlane{}
	for j := 0; j < height; j++ {
		row := make([]string, width)
		for i := 0; i < width; i++ {
			row[i] = drawChar("rain")
		}
		screen = append(screen, row)
	}

	for {
		clearTerminal()
		screen.draw()
		screen = screen.step()
		time.Sleep(sleepMillis * time.Millisecond)
	}

}
