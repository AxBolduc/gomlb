package components

import (
	"gomlb/api/mlb"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func BuildPlayerStatTable(players []mlb.BoxscorePlayer, initialFocus bool) table.Model {
	playerNameMaxLen := 0
	for _, player := range players {
		if len(player.Person.FullName) > playerNameMaxLen {
			playerNameMaxLen = len(player.Person.FullName)
		}
	}
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

	// tableColumns := []table.Column{
	// 	table.NewColumn("playerName", "Player", playerNameMaxLen+1),
	// 	table.NewColumn("pos", "Pos", 3),
	// 	table.NewColumn("ab", "AB", 2),
	// 	table.NewColumn("r", "R", 1),
	// 	table.NewColumn("h", "H", 1),
	// 	table.NewColumn("rbi", "RBI", 3),
	// 	table.NewColumn("bb", "BB", 2),
	// 	table.NewColumn("k", "K", 1),
	// 	table.NewColumn("avg", "AVG", 5),
	// 	table.NewColumn("ops", "OPS", 5),
	// 	table.NewFlexColumn("pad", "", 1),
	// }
	lineup := []table.Row{}

	for _, player := range players {
		lineup = append(lineup, boxscorePlayerToTableRow(player))
	}

	return table.New(tableColumns).WithRows(lineup).Focused(initialFocus).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Left))

}

func boxscorePlayerToTableRow(player mlb.BoxscorePlayer) table.Row {
	row := table.NewRow(table.RowData{
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
