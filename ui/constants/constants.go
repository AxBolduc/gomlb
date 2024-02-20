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
