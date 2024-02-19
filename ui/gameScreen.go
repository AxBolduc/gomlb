package ui

import (
	"fmt"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	"github.com/axbolduc/gomlb/ui/components"
	"github.com/axbolduc/gomlb/ui/constants"

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
	previousModel             Model
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

		splitColumnTargetWidth := (msg.Width - h) / 2

		m.homeBattersTable = m.homeBattersTable.WithTargetWidth(splitColumnTargetWidth)
		m.awayBattersTable = m.awayBattersTable.WithTargetWidth(splitColumnTargetWidth)
		m.homePitchersTable = m.homePitchersTable.WithTargetWidth(splitColumnTargetWidth)
		m.awayPitchersTable = m.awayPitchersTable.WithTargetWidth(splitColumnTargetWidth)

		m.linescoreTable = m.linescoreTable.WithTargetWidth(msg.Width - h)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Back):
			return m.previousModel, tea.Batch()
		case key.Matches(msg, constants.Keymap.Left):
			m = m.swapFocusedTable()
		case key.Matches(msg, constants.Keymap.Right):
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

	ui := lipgloss.JoinVertical(lipgloss.Center, scoreBox, m.linescoreTable.View(), battersTables, pitchersTables)

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
