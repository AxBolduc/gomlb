package ui

import (
	"fmt"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	"github.com/axbolduc/gomlb/ui/components"
	"github.com/axbolduc/gomlb/ui/constants"
	"github.com/axbolduc/gomlb/ui/popup"
	"github.com/axbolduc/gomlb/ui/popup/batterStatsPopup"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var DIRECTIONS = map[string]string{
	"LEFT":  "LEFT",
	"RIGHT": "RIGHT",
	"UP":    "UP",
	"DOWN":  "DOWN",
}

var TABLE_TO_INDEX_MAP = map[string]int{
	"awayBatters":  0,
	"homeBatters":  1,
	"awayPitchers": 2,
	"homePitchers": 3,
}

var INDEX_TO_TABLE_MAP = map[int]string{
	0: "awayBatters",
	1: "homeBatters",
	2: "awayPitchers",
	3: "homePitchers",
}

type GameScreenModel struct {
	linescoreTable    table.Model
	playerTables      []table.Model
	focusedTableIndex int
	game              mlb.Game
	boxscore          mlb.Boxscore
	help              help.Model
	previousModel     Model
	popup             popup.IPopup
	width, height     int
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
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("left/h", "Table Previous Page"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("right/l", "Table Next Page"),
	),
	LeftTable: key.NewBinding(
		key.WithKeys("shift+left", "H"),
		key.WithHelp("shift+left/H", "Left Table"),
	),
	RightTable: key.NewBinding(
		key.WithKeys("shift+right", "L"),
		key.WithHelp("shift+right/L", "Right Table"),
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
		key.WithKeys("J", "shift+down"),
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
		constants.DocStyle = lipgloss.NewStyle().Width(msg.Width).Height(msg.Height).Padding(constants.VPadding, constants.HPadding)

		m.width = msg.Width - constants.DocStyle.GetHorizontalFrameSize()
		m.height = msg.Height - constants.DocStyle.GetVerticalFrameSize()

		splitColumnTargetWidth := m.width / 2

		for i, table := range m.playerTables {
			m.playerTables[i] = table.WithTargetWidth(splitColumnTargetWidth)
		}

		m.linescoreTable = m.linescoreTable.WithTargetWidth(m.width)

		if m.popup != nil {
			// Send resize message to the popup
			m.popup = m.popup.Resize(msg, m.renderMainScreen())
		}

	case tea.KeyMsg:
		if m.popup == nil {
			m, cmd = m.updateFocusedTable(msg)
		}

		switch {
		case key.Matches(msg, gameScreenKM.Quit):
			return m, tea.Quit
		case key.Matches(msg, gameScreenKM.Back):
			if m.popup == nil {
				return m.previousModel, tea.Batch()
			} else {
				gameScreenKM.SetEnabled(true)
				m.popup = nil
			}
		case key.Matches(msg, gameScreenKM.LeftTable):
			m = m.updateFocusedTableIndex(DIRECTIONS["LEFT"])
		case key.Matches(msg, gameScreenKM.RightTable):
			m = m.updateFocusedTableIndex(DIRECTIONS["RIGHT"])
		case key.Matches(msg, gameScreenKM.UpTable):
			m = m.updateFocusedTableIndex(DIRECTIONS["UP"])
		case key.Matches(msg, gameScreenKM.DownTable):
			m = m.updateFocusedTableIndex(DIRECTIONS["DOWN"])
		case key.Matches(msg, gameScreenKM.Enter):
			batterId := m.playerTables[m.focusedTableIndex].HighlightedRow().Data["id"].(int)
			m.popup = batterStatsPopup.New(m.View(), batterId, m.width-2*constants.PopupHPadding, m.height-2*constants.PopupVPadding)
			if m.popup != nil {
				gameScreenKM.SetEnabled(false)
				return m, m.popup.Init()
			}
		}
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m GameScreenModel) updateFocusedTable(msg tea.Msg) (GameScreenModel, tea.Cmd) {
	var cmd tea.Cmd
	m.playerTables[m.focusedTableIndex], cmd = m.playerTables[m.focusedTableIndex].Update(msg)

	return m, cmd
}

func (m GameScreenModel) updateFocusedTableIndex(direction string) GameScreenModel {
	m.playerTables[m.focusedTableIndex] = m.playerTables[m.focusedTableIndex].Focused(false)

	numTables := len(m.playerTables)

	switch direction {
	case DIRECTIONS["LEFT"]:
		m.focusedTableIndex = (m.focusedTableIndex - 1 + numTables) % numTables
	case DIRECTIONS["RIGHT"]:
		m.focusedTableIndex = (m.focusedTableIndex + 1) % numTables
	case DIRECTIONS["UP"]:
		m.focusedTableIndex = (m.focusedTableIndex - 2 + numTables) % numTables
	case DIRECTIONS["DOWN"]:
		m.focusedTableIndex = (m.focusedTableIndex + 2) % numTables
	}

	m.playerTables[m.focusedTableIndex] = m.playerTables[m.focusedTableIndex].Focused(true)

	return m
}

func (m GameScreenModel) View() string {

	if m.popup != nil {
		return m.popup.View()
	}

	return m.renderMainScreen()
}

func (m GameScreenModel) renderMainScreen() string {
	scoreBox := components.RenderScoreText(m.game.Linescore.Teams.Away.Runs, m.game.Linescore.Teams.Home.Runs, m.game.Teams.Away.Team.Name, m.game.Teams.Home.Team.Name)

	battersTables := lipgloss.JoinHorizontal(lipgloss.Top, m.playerTables[TABLE_TO_INDEX_MAP["awayBatters"]].View(), m.playerTables[TABLE_TO_INDEX_MAP["homeBatters"]].View())
	pitchersTables := lipgloss.JoinHorizontal(lipgloss.Top, m.playerTables[TABLE_TO_INDEX_MAP["awayPitchers"]].View(), m.playerTables[TABLE_TO_INDEX_MAP["homePitchers"]].View())

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

	initialAwayTableFocused := true
	initialBatterTableFocused := true

	awayBatters := positionListToPlayerList(boxscore.Teams.Away.BattingOrder, boxscore.Teams.Away.Players)
	awayPitchers := positionListToPlayerList(boxscore.Teams.Away.Pitchers, boxscore.Teams.Away.Players)

	homeBatters := positionListToPlayerList(boxscore.Teams.Home.BattingOrder, boxscore.Teams.Home.Players)
	homePitchers := positionListToPlayerList(boxscore.Teams.Home.Pitchers, boxscore.Teams.Home.Players)

	awayBattersTable := components.BuildBatterStatsTable(awayBatters, initialAwayTableFocused && initialBatterTableFocused)
	awayPitchersTable := components.BuildPitcherStatsTable(awayPitchers, initialAwayTableFocused && !initialBatterTableFocused).WithPageSize(5)

	homeBattersTable := components.BuildBatterStatsTable(homeBatters, !initialAwayTableFocused && initialBatterTableFocused)
	homePitchersTable := components.BuildPitcherStatsTable(homePitchers, initialAwayTableFocused && !initialBatterTableFocused).WithPageSize(5)

	playerTables := []table.Model{awayBattersTable, homeBattersTable, awayPitchersTable, homePitchersTable}

	linescoreTable := components.BuildLinescoreTable(game.Teams.Away.Team.Name, game.Teams.Home.Team.Name, game.Linescore)

	gameScreenModel := GameScreenModel{
		game:              game,
		previousModel:     previousModel,
		linescoreTable:    linescoreTable,
		playerTables:      playerTables,
		focusedTableIndex: TABLE_TO_INDEX_MAP["awayBatters"],
		boxscore:          *boxscore,
		help:              help.New(),
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
