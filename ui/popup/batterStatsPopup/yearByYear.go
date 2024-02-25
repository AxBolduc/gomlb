package batterStatsPopup

import (
	"log"
	"slices"
	"strconv"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func getYearByYearHittingStats(batterId int) *mlb.PlayerHittingStats {
	statsRepo := repositories.NewPlayerStatsRepository()
	batterStats, err := statsRepo.GetYearByYearHittingStatsByPlayerId(batterId)
	if err != nil {
		return nil
	}

	if len(batterStats.Stats) == 0 {
		log.Printf("No stats batting stats for player %d", batterId)
		return nil
	}

	return &batterStats.Stats[0]
}

func getTeamNameMaxLen(splits []mlb.HittingStatsSplit) int {
	teamNameMaxLen := 0
	for _, split := range splits {
		if teamNameLen := len(split.Team.Name); teamNameLen > teamNameMaxLen {
			teamNameMaxLen = teamNameLen
		}
	}

	return teamNameMaxLen
}

func buildYearByYearHittingStatsTable(hittingStats []mlb.HittingStatsSplit, width int) table.Model {
	alignCenter := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)

	teamNameWidth := getTeamNameMaxLen(hittingStats)
	tableColumns := []table.Column{
		table.NewColumn("season", "Year", 4),
		table.NewColumn("team", "Team", teamNameWidth).WithStyle(alignCenter),
		table.NewFlexColumn("ab", "AB", 1).WithStyle(alignCenter),
		table.NewFlexColumn("r", "R", 1).WithStyle(alignCenter),
		table.NewFlexColumn("h", "H", 1).WithStyle(alignCenter),
		table.NewFlexColumn("2b", "2B", 1).WithStyle(alignCenter),
		table.NewFlexColumn("3b", "3B", 1).WithStyle(alignCenter),
		table.NewFlexColumn("hr", "HR", 1).WithStyle(alignCenter),
		table.NewFlexColumn("rbi", "RBI", 1).WithStyle(alignCenter),
		table.NewFlexColumn("bb", "BB", 1).WithStyle(alignCenter),
		table.NewFlexColumn("k", "K", 1).WithStyle(alignCenter),
		table.NewFlexColumn("avg", "AVG", 1).WithStyle(alignCenter),
		table.NewFlexColumn("obp", "OBP", 1).WithStyle(alignCenter),
		table.NewFlexColumn("slg", "SLG", 1).WithStyle(alignCenter),
		table.NewFlexColumn("ops", "OPS", 1).WithStyle(alignCenter),
	}

	tableRows := yearByYearStatsSplitsToTableRows(hittingStats)

	table := table.New(tableColumns).WithRows(tableRows).WithTargetWidth(width)

	return table
}

func statSplitToYearByYearHittingRow(split mlb.HittingStatsSplit) table.Row {

	rowTeam := split.Team.Name
	if rowTeam == "" {
		rowTeam = "Total"
	}

	return table.NewRow(table.RowData{
		"season": split.Season,
		"team":   rowTeam,
		"ab":     split.Stat.AtBats,
		"r":      split.Stat.Runs,
		"h":      split.Stat.Hits,
		"2b":     split.Stat.Doubles,
		"3b":     split.Stat.Triples,
		"hr":     split.Stat.HomeRuns,
		"rbi":    split.Stat.Rbi,
		"bb":     split.Stat.BaseOnBalls,
		"k":      split.Stat.StrikeOuts,
		"avg":    split.Stat.Avg,
		"obp":    split.Stat.Obp,
		"slg":    split.Stat.Slg,
		"ops":    split.Stat.Ops,
	})
}

func yearByYearStatsSplitsToTableRows(splits []mlb.HittingStatsSplit) []table.Row {
	var rows []table.Row

	for _, split := range splits {
		rows = append(rows, statSplitToYearByYearHittingRow(split))
	}

	slices.SortStableFunc(rows, func(a, b table.Row) int {
		aConvertedInt, aErr := strconv.Atoi(a.Data["season"].(string))
		bConvertedInt, bErr := strconv.Atoi(b.Data["season"].(string))

		if aErr != nil || bErr != nil {
			log.Fatalf("Failed to sort the year by year stats, could not convert year to int, \n %s \n %s", aErr.Error(), bErr.Error())
		}

		return bConvertedInt - aConvertedInt
	})

	return rows
}
