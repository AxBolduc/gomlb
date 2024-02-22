package components

import (
	"github.com/axbolduc/gomlb/api/mlb"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func BuildBatterStatsTable(players []mlb.BoxscorePlayer, initialFocus bool) table.Model {
	playerNameMaxLen := getPlayerNameMaxLen(players)
	tableColumns := []table.Column{
		table.NewColumn("playerName", "Player", playerNameMaxLen+1),
		table.NewColumn("pos", "Pos", 3).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("ab", "AB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("r", "R", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("h", "H", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("rbi", "RBI", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("bb", "BB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("k", "K", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("avg", "AVG", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("ops", "OPS", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
	}

	lineup := []table.Row{}

	for _, player := range players {
		lineup = append(lineup, boxscorePlayerToBatterTableRow(player))
	}

	return table.New(tableColumns).WithRows(lineup).Focused(initialFocus).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Left))

}

func getPlayerNameMaxLen(players []mlb.BoxscorePlayer) int {
	playerNameMaxLen := 0
	for _, player := range players {
		if len(player.Person.FullName) > playerNameMaxLen {
			playerNameMaxLen = len(player.Person.FullName)
		}
	}

	return playerNameMaxLen
}

func BuildPitcherStatsTable(players []mlb.BoxscorePlayer, initialFocus bool) table.Model {
	playerNameMaxLen := getPlayerNameMaxLen(players)

	tableColumns := []table.Column{
		table.NewColumn("playerName", "Player", playerNameMaxLen+1),
		table.NewColumn("pos", "Pos", 3).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("ip", "IP", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("h", "H", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("r", "R", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("er", "ER", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("bb", "BB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("k", "K", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("hr", "HR", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("era", "ERA", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
	}

	lineup := []table.Row{}

	for _, player := range players {
		lineup = append(lineup, boxscorePlayerToPitcherTableRow(player))
	}

	return table.New(tableColumns).WithRows(lineup).Focused(initialFocus).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Left))

}

func boxscorePlayerToBatterTableRow(player mlb.BoxscorePlayer) table.Row {
	row := table.NewRow(table.RowData{
		"id":         player.Person.Id,
		"playerName": player.Person.FullName,
		"pos":        player.Position.Abbreviation,
		"ab":         player.Stats.Batting.AtBats,
		"r":          player.Stats.Batting.Runs,
		"h":          player.Stats.Batting.Hits,
		"rbi":        player.Stats.Batting.Rbi,
		"bb":         player.Stats.Batting.BaseOnBalls,
		"k":          player.Stats.Batting.StrikeOuts,
		"avg":        player.SeasonStats.Batting.Avg,
		"ops":        player.SeasonStats.Batting.Ops,
	})

	return row
}

func boxscorePlayerToPitcherTableRow(player mlb.BoxscorePlayer) table.Row {
	row := table.NewRow(table.RowData{
		"playerName": player.Person.FullName,
		"pos":        player.Position.Abbreviation,
		"ip":         player.Stats.Pitching.InningsPitched,
		"h":          player.Stats.Pitching.Hits,
		"r":          player.Stats.Pitching.Runs,
		"er":         player.Stats.Pitching.EarnedRuns,
		"bb":         player.Stats.Pitching.BaseOnBalls,
		"k":          player.Stats.Pitching.StrikeOuts,
		"hr":         player.Stats.Pitching.HomeRuns,
		"era":        player.SeasonStats.Pitching.Era,
	})

	return row
}
