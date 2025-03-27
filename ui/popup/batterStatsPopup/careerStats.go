package batterStatsPopup

import (
	"log"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func getCareerHittingStats(batterId int) *mlb.HittingStatsSplit {
	statsRepo := repositories.NewPlayerStatsRepository()
	batterStats, err := statsRepo.GetCareerStatsByPlayerId(batterId)
	if err != nil {
		return nil
	}

	if len(batterStats.Stats) == 0 {
		log.Printf("No stats batting stats for player %d", batterId)
		return nil
	}

	if len(batterStats.Stats[0].Splits) == 0 {
		log.Printf("No splits for batting stast for player %d", batterId)
	}

	return &batterStats.Stats[0].Splits[0]
}

func buildBattingStatsTable(battingStats mlb.HittingStats, width int) table.Model {
	tableColumns := []table.Column{
		table.NewFlexColumn("ab", "AB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("r", "R", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("h", "H", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("2b", "2B", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("3b", "3B", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("hr", "HR", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("rbi", "RBI", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("bb", "BB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("k", "K", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("avg", "AVG", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("obp", "OBP", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("slg", "SLG", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("ops", "OPS", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
	}

	tableRow := table.NewRow(table.RowData{
		"ab":  battingStats.AtBats,
		"r":   battingStats.Runs,
		"h":   battingStats.Hits,
		"2b":  battingStats.Doubles,
		"3b":  battingStats.Triples,
		"hr":  battingStats.HomeRuns,
		"rbi": battingStats.Rbi,
		"bb":  battingStats.BaseOnBalls,
		"k":   battingStats.StrikeOuts,
		"avg": battingStats.Avg,
		"obp": battingStats.Obp,
		"slg": battingStats.Slg,
		"ops": battingStats.Ops,
	})

	table := table.New(tableColumns).WithRows([]table.Row{tableRow}).WithTargetWidth(width)

	return table
}
