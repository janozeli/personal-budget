package main

import (
	"main/tui"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	tui.ShowDataBaseOptions(app)
}
