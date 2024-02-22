package popup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type BatterStatsPopup struct {
	popup    Popup
	batterId int
}

func (p BatterStatsPopup) Init() tea.Cmd {
	return nil
}

func (p BatterStatsPopup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "q":
			return p, tea.Quit
		}
	}

	return p, tea.Batch(cmds...)
}

func (p BatterStatsPopup) View() string {

	p.popup.SetFg("Hello my name is test")
	return p.popup.View()
}

func NewBatterStatsPopup(bg string, batterId int, width int, height int) *BatterStatsPopup {
	// Get batter info from batterID?

	title := fmt.Sprintf("Batting stats for player %d", batterId)

	popup := NewPopup(bg, title, width, height)

	return &BatterStatsPopup{popup: popup}
}
