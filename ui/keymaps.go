package ui

import "github.com/charmbracelet/bubbles/key"

type GamesListKM struct {
	Up          key.Binding
	Down        key.Binding
	Previous    key.Binding
	Next        key.Binding
	Quit        key.Binding
	Enter       key.Binding
	FocusPicker key.Binding
}

func (k GamesListKM) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (k GamesListKM) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Previous, k.Next, k.Quit, k.Enter}
}

type GameScreenKM struct {
	Up        key.Binding
	Down      key.Binding
	UpTable   key.Binding
	DownTable key.Binding
	Back      key.Binding
	Quit      key.Binding
	Enter     key.Binding
	Left      key.Binding
	Right     key.Binding
}

func (k GameScreenKM) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (k GameScreenKM) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Left, k.Right, k.UpTable, k.DownTable, k.Enter, k.Back, k.Quit}
}

func (k *GameScreenKM) SetEnabled(enabled bool) {
	k.Up.SetEnabled(enabled)
	k.Down.SetEnabled(enabled)
	k.DownTable.SetEnabled(enabled)
	k.UpTable.SetEnabled(enabled)
	k.Enter.SetEnabled(enabled)
	k.Left.SetEnabled(enabled)
	k.Right.SetEnabled(enabled)
}
