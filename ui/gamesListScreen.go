package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	"github.com/axbolduc/gomlb/ui/constants"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	date          time.Time
	gameList      list.Model
	help          help.Model
	width, height int
}

var gamesListKM = GamesListKM{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↲/enter", "select"),
	),
	Previous: key.NewBinding(
		key.WithKeys("<"),
		key.WithHelp("<", "previous day"),
	),
	Next: key.NewBinding(
		key.WithKeys(">"),
		key.WithHelp(">", "next day"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q", "esc"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		constants.DocStyle = lipgloss.NewStyle().Width(msg.Width).Height(msg.Height).Padding(constants.VPadding, constants.HPadding)

		m.width = msg.Width - constants.DocStyle.GetHorizontalFrameSize()
		m.height = msg.Height - constants.DocStyle.GetVerticalFrameSize()

		m.gameList.SetSize(m.width, m.height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, gamesListKM.Quit):
			return m, tea.Quit
		case key.Matches(msg, gamesListKM.Next):
			// Update the model for tomorrow
			m = m.UpdateWithDate(m.date.AddDate(0, 0, 1))
		case key.Matches(msg, gamesListKM.Previous):
			//Update the model for yesterday
			m = m.UpdateWithDate(m.date.AddDate(0, 0, -1))
		case key.Matches(msg, gamesListKM.Enter):
			activeGame := m.gameList.SelectedItem().(mlb.Game)
			gameScreenModel := InitGameScreenModel(activeGame, m)
			return gameScreenModel.Update(constants.WindowSize)
		}
	}

	m.gameList, cmd = m.gameList.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	return constants.DocStyle.Render(m.gameList.View())
}

func (m Model) UpdateWithDate(date time.Time) Model {
	schedule, err := repositories.NewScheduleRepository().GetScheduleForDate(date)
	if err != nil {
		panic(err)
	}

	m.date = date

	if len(schedule.Dates) == 0 {
		m.gameList.SetItems([]list.Item{})
		m.gameList.Title = fmt.Sprintf("No games on %s", date.Format(time.DateOnly))
	} else {
		newListItems := gamesToItems(schedule.Dates[0].Games)
		m.gameList.SetItems(newListItems)
		m.gameList.Title = fmt.Sprintf("Games for %s", date.Format(time.DateOnly))
	}

	return m
}

func InitModel(date time.Time) tea.Model {
	schedule, err := repositories.NewScheduleRepository().GetScheduleForDate(date)

	if err != nil {
		panic(err)
	}

	if len(schedule.Dates) == 0 {
		fmt.Println("No games on", date)
		os.Exit(0)
	}

	games := schedule.Dates[0].Games
	items := gamesToItems(games)

	// Custom styling
	customDelegate := list.NewDefaultDelegate()
	customDelegate.ShortHelpFunc = func() []key.Binding {
		return gamesListKM.ShortHelp()
	}

	customDelegate.FullHelpFunc = func() [][]key.Binding {
		return gamesListKM.FullHelp()
	}

	customDelegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(constants.PrimaryColor).
		Foreground(constants.PrimaryColor).
		Padding(0, 0, 0, 1)

	customDelegate.Styles.SelectedDesc = customDelegate.Styles.SelectedTitle.Copy()

	m := Model{gameList: list.New(items, customDelegate, 0, 0), date: date, help: help.New()}
	m.gameList.Title = fmt.Sprintf("Games on %s", date.Format(time.DateOnly))

	return m
}

func gamesToItems(games []mlb.Game) []list.Item {
	items := make([]list.Item, len(games))

	for i, game := range games {
		items[i] = list.Item(game)
	}

	return items
}
