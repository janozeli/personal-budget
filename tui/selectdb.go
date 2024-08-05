package tui

import (
	"fmt"
	"main/db"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectDB struct {
	choices []string
	cursor  int
}

func SelectDBInitialModel() SelectDB {
	return SelectDB{choices: db.VerifyExistingDBs()}
}

func (m SelectDB) Init() tea.Cmd {
	return nil
}

func (m SelectDB) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", "down":
			m.cursor++
		case "k", "up":
			m.cursor--
		}
	}

	if m.cursor < 0 {
		m.cursor = len(m.choices) - 1
	} else if m.cursor > len(m.choices)-1 {
		m.cursor = 0
	}

	return m, nil
}

func (m SelectDB) View() string {
	s := "Selecione uma conta:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPressione 'q' para sair"
	return s
}
