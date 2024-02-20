package ui

import (
	"fmt"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	"github.com/axbolduc/gomlb/ui/components"
	"github.com/axbolduc/gomlb/ui/constants"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type GameScreenModel struct {
	linescoreTable            table.Model
	awayBattersTable          table.Model
	awayPitchersTable         table.Model
	homeBattersTable          table.Model
	homePitchersTable         table.Model
	isAwayBattersTableFocused bool
	game                      mlb.Game
	boxscore                  mlb.Boxscore
	help                      help.Model
	width                     int
	previousModel             Model
}

var gameScreenKM = GameScreenKM{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↲/enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q", "esc"),
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
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("up/k", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("down/j", "Down"),
	),
	UpTable: key.NewBinding(
		key.WithKeys("K", "shift+up"),
		key.WithHelp("shift+up/K", "Up a Table"),
	),
	DownTable: key.NewBinding(
		key.WithKeys("K", "shift+down"),
		key.WithHelp("shift+down/J", "Down a Table"),
	),
}

func (m GameScreenModel) Init() tea.Cmd {
	return nil
}

func (m GameScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, _ := constants.DocStyle.GetFrameSize()

		m.width = msg.Width - h

		splitColumnTargetWidth := m.width / 2

		m.homeBattersTable = m.homeBattersTable.WithTargetWidth(splitColumnTargetWidth)
		m.awayBattersTable = m.awayBattersTable.WithTargetWidth(splitColumnTargetWidth)
		m.homePitchersTable = m.homePitchersTable.WithTargetWidth(splitColumnTargetWidth)
		m.awayPitchersTable = m.awayPitchersTable.WithTargetWidth(splitColumnTargetWidth)

		m.linescoreTable = m.linescoreTable.WithTargetWidth(m.width)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, gameScreenKM.Quit):
			return m, tea.Quit
		case key.Matches(msg, gameScreenKM.Back):
			return m.previousModel, tea.Batch()
		case key.Matches(msg, gameScreenKM.Left):
			m = m.swapFocusedTable()
		case key.Matches(msg, gameScreenKM.Right):
			m = m.swapFocusedTable()
		}
	}

	m.awayBattersTable, cmd = m.awayBattersTable.Update(msg)
	cmds = append(cmds, cmd)

	m.homeBattersTable, cmd = m.homeBattersTable.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m GameScreenModel) swapFocusedTable() GameScreenModel {
	m.isAwayBattersTableFocused = !m.isAwayBattersTableFocused
	m.awayBattersTable = m.awayBattersTable.Focused(m.isAwayBattersTableFocused)
	m.homeBattersTable = m.homeBattersTable.Focused(!m.isAwayBattersTableFocused)

	return m
}

func (m GameScreenModel) View() string {
	scoreBox := components.RenderScoreText(m.game.Linescore.Teams.Away.Runs, m.game.Linescore.Teams.Home.Runs, m.game.Teams.Away.Team.Name, m.game.Teams.Home.Team.Name)

	battersTables := lipgloss.JoinHorizontal(lipgloss.Top, m.awayBattersTable.View(), m.homeBattersTable.View())
	pitchersTables := lipgloss.JoinHorizontal(lipgloss.Top, m.awayPitchersTable.View(), m.homePitchersTable.View())

	helpContainer := lipgloss.NewStyle().
		SetString(m.help.View(gameScreenKM)).
		Width(m.width).
		Align(lipgloss.Left).
		PaddingTop(1).
		String()

	ui := lipgloss.JoinVertical(lipgloss.Center, scoreBox, m.linescoreTable.View(), battersTables, pitchersTables, helpContainer)

	return constants.DocStyle.Render(ui)
}

func InitGameScreenModel(game mlb.Game, previousModel Model) *GameScreenModel {
	boxscore, err := repositories.NewBoxscoreRepository().GetBoxscoreFromGamePk(game.GamePk)
	if err != nil {
		fmt.Printf("Failed to get boxscore")
		panic(err)
	}

	awayBatters := positionListToPlayerList(boxscore.Teams.Away.BattingOrder, boxscore.Teams.Away.Players)
	awayPitchers := positionListToPlayerList(boxscore.Teams.Away.Pitchers, boxscore.Teams.Away.Players)

	homeBatters := positionListToPlayerList(boxscore.Teams.Home.BattingOrder, boxscore.Teams.Home.Players)
	homePitchers := positionListToPlayerList(boxscore.Teams.Home.Pitchers, boxscore.Teams.Home.Players)

	awayBattersTable := components.BuildBatterStatsTable(awayBatters, true)
	awayPitchersTable := components.BuildPitcherStatsTable(awayPitchers)

	homePlayerTable := components.BuildBatterStatsTable(homeBatters, false)
	homePitchersTable := components.BuildPitcherStatsTable(homePitchers)

	linescoreTable := components.BuildLinescoreTable(game.Teams.Away.Team.Name, game.Teams.Home.Team.Name, game.Linescore)

	gameScreenModel := GameScreenModel{
		game:                      game,
		previousModel:             previousModel,
		awayBattersTable:          awayBattersTable,
		awayPitchersTable:         awayPitchersTable,
		homeBattersTable:          homePlayerTable,
		homePitchersTable:         homePitchersTable,
		linescoreTable:            linescoreTable,
		isAwayBattersTableFocused: true,
		boxscore:                  *boxscore,
		help:                      help.New(),
	}

	return &gameScreenModel
}

func positionListToPlayerList(batters []int, roster map[string]mlb.BoxscorePlayer) []mlb.BoxscorePlayer {
	var players []mlb.BoxscorePlayer
	for _, playerId := range batters {
		retrievedPlayer := roster[fmt.Sprintf("ID%d", playerId)]
		players = append(players, retrievedPlayer)
	}

	return players
}
