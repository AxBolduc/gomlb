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

var VPadding = 1
var HPadding = 2

var PopupVPadding = 3
var PopupHPadding = 6

var PrimaryColor = lipgloss.AdaptiveColor{Light: "#5875b4", Dark: "#5875b4"}
var DocStyle lipgloss.Style

var WindowSize tea.WindowSizeMsg
