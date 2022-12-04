package main

import (
	"fmt"
	"math/rand"
	"time"
)

const snowDensity = 0.05
const sleepMillisLow = 400
const sleepMillisHigh = 850

type snowPlane [][]string

func (p snowPlane) draw() {
	for _, row := range p {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}

func (p snowPlane) step() snowPlane {
	// everything down one and to the left one.
	// drop the bottom row and
	// instantiate a new random row at the top

	xShape := len(p[0])
	yShape := len(p)
	var newRow = make([]string, xShape)
	for i := 0; i < xShape; i++ {
		newRow[i] = drawChar("snow")
	}

	newPlane := snowPlane{}
	newPlane = append(newPlane, newRow)
	rnd := rand.Float64()
	if rnd < 0.3333 {
		// step down and to the left
		for i := 0; i < yShape-1; i++ {
			newPlane = append(newPlane, append(p[i][1:], drawChar("snow")))
		}
	} else if rnd < 0.6666 {
		// step down
		for i := 0; i < yShape-1; i++ {
			newPlane = append(newPlane, p[i])
		}
	} else {
		// step down and to the right
		for i := 0; i < yShape-1; i++ {
			newPlane = append(newPlane, append([]string{drawChar("snow")}, p[i][:xShape-1]...))
		}
	}

	return newPlane

}

func snow(width, height int) {
	// Initialize the screen
	screen := snowPlane{}
	for j := 0; j < height; j++ {
		row := make([]string, width)
		for i := 0; i < width; i++ {
			row[i] = drawChar("snow")
		}
		screen = append(screen, row)
	}

	for {
		clearTerminal()
		screen.draw()
		screen = screen.step()

		rndSleep := sleepMillisLow + rand.Float64()*(sleepMillisHigh-sleepMillisLow)
		time.Sleep(time.Duration(rndSleep) * time.Millisecond)
	}

}
