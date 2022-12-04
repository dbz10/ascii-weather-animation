package main

import "github.com/buger/goterm"

func clearTerminal() {
	goterm.MoveCursor(1, 2)
	goterm.Flush()

}
