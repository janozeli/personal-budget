package main

import (
	"main/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.SelectDBInitialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
