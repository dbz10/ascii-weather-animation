package main

import "github.com/buger/goterm"

func clearTerminal() {
	goterm.Flush()
	goterm.Clear()
	goterm.MoveCursor(1, 1)

}
