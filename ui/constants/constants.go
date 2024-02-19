package constants

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	Enter     key.Binding
	Yesterday key.Binding
	Tomorrow  key.Binding
	Back      key.Binding
	Quit      key.Binding
	Left      key.Binding
	Right     key.Binding
}

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

var WindowSize tea.WindowSizeMsg

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↲/enter", "select"),
	),
	Yesterday: key.NewBinding(
		key.WithKeys("<"),
		key.WithHelp("<", "previous day"),
	),
	Tomorrow: key.NewBinding(
		key.WithKeys(">"),
		key.WithHelp(">", "next day"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("left/h", "Left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("right/l", "Right"),
	),
}
