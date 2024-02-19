package ui

import (
	"fmt"
	"gomlb/api/mlb"
	"gomlb/api/mlb/repositories"
	"gomlb/ui/constants"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	date     time.Time
	gameList list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		h, v := constants.DocStyle.GetFrameSize()
		m.gameList.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Tomorrow):
			// Update the model for tomorrow
			m = m.UpdateWithDate(m.date.AddDate(0, 0, 1))
		case key.Matches(msg, constants.Keymap.Yesterday):
			//Update the model for yesterday
			m = m.UpdateWithDate(m.date.AddDate(0, 0, -1))
		case key.Matches(msg, constants.Keymap.Enter):
			activeGame := m.gameList.SelectedItem().(mlb.Game)
			gameScreenModel := InitGameScreenModel(activeGame, m)
			return gameScreenModel.Update(constants.WindowSize)
		}
	}

	m.gameList, cmd = m.gameList.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return constants.DocStyle.Render(m.gameList.View())
}

func (m Model) UpdateWithDate(date time.Time) Model {
	schedule, err := repositories.NewScheduleRepository().GetScheduleForDate(date)
	if err != nil {
		panic(err)
	}

	if len(schedule.Dates) == 0 {
		return m
	}

	m.date = date

	newListItems := gamesToItems(schedule.Dates[0].Games)
	m.gameList.SetItems(newListItems)
	m.gameList.Title = fmt.Sprintf("Games for %s", date.Format(time.DateOnly))

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
	customDelegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#5875b4", Dark: "#5875b4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#5875b4", Dark: "#5875b4"}).
		Padding(0, 0, 0, 1)

	customDelegate.Styles.SelectedDesc = customDelegate.Styles.SelectedTitle.Copy()

	m := Model{gameList: list.New(items, customDelegate, 50, 25), date: date}
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
