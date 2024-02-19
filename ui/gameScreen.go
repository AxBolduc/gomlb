package ui

import (
	"fmt"
	"gomlb/api/mlb"
	"gomlb/api/mlb/repositories"
	"gomlb/ui/components"
	scoretext "gomlb/ui/components"
	"gomlb/ui/constants"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type GameScreenModel struct {
	linescoreTable           table.Model
	awayPlayerTable          table.Model
	homePlayerTable          table.Model
	isAwayPlayerTableFocused bool
	game                     mlb.Game
	boxscore                 mlb.Boxscore
	previousModel            Model
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
		m.homePlayerTable = m.homePlayerTable.WithTargetWidth((msg.Width - h) / 2)
		m.awayPlayerTable = m.awayPlayerTable.WithTargetWidth((msg.Width - h) / 2)
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

	m.awayPlayerTable, cmd = m.awayPlayerTable.Update(msg)
	cmds = append(cmds, cmd)

	m.homePlayerTable, cmd = m.homePlayerTable.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m GameScreenModel) swapFocusedTable() GameScreenModel {
	m.isAwayPlayerTableFocused = !m.isAwayPlayerTableFocused
	m.awayPlayerTable = m.awayPlayerTable.Focused(m.isAwayPlayerTableFocused)
	m.homePlayerTable = m.homePlayerTable.Focused(!m.isAwayPlayerTableFocused)

	return m
}

func (m GameScreenModel) View() string {
	scoreBox := scoretext.RenderScoreText(m.game.Linescore.Teams.Away.Runs, m.game.Linescore.Teams.Home.Runs, m.game.Teams.Away.Team.Name, m.game.Teams.Home.Team.Name)

	lineups := lipgloss.JoinHorizontal(lipgloss.Center, m.awayPlayerTable.View(), m.homePlayerTable.View())

	return constants.DocStyle.Render(scoreBox + m.linescoreTable.View() + "\n" + lineups)
}

func InitGameScreenModel(game mlb.Game, previousModel Model) *GameScreenModel {
	boxscore, err := repositories.NewBoxscoreRepository().GetBoxscoreFromGamePk(game.GamePk)
	if err != nil {
		fmt.Printf("Failed to get boxscore")
		panic(err)
	}

	awayBatters := battersListToPlayerList(boxscore.Teams.Away.BattingOrder, boxscore.Teams.Away.Players)
	homeBatters := battersListToPlayerList(boxscore.Teams.Home.BattingOrder, boxscore.Teams.Home.Players)

	awayPlayerTable := components.BuildPlayerStatTable(awayBatters, true)
	homePlayerTable := components.BuildPlayerStatTable(homeBatters, false)

	linescoreTable := components.BuildLinescoreTable(game.Teams.Away.Team.Name, game.Teams.Home.Team.Name, game.Linescore)

	gameScreenModel := GameScreenModel{
		game:                     game,
		previousModel:            previousModel,
		awayPlayerTable:          awayPlayerTable,
		homePlayerTable:          homePlayerTable,
		linescoreTable:           linescoreTable,
		isAwayPlayerTableFocused: true,
		boxscore:                 *boxscore,
	}

	return &gameScreenModel
}

func battersListToPlayerList(batters []int, roster map[string]mlb.BoxscorePlayer) []mlb.BoxscorePlayer {
	var players []mlb.BoxscorePlayer
	for _, playerId := range batters {
		retrievedPlayer := roster[fmt.Sprintf("ID%d", playerId)]
		players = append(players, retrievedPlayer)
	}

	return players
}
