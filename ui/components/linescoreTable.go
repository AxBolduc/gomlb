package components

import (
	"fmt"
	"sports-cli/api/mlb"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func generateInningsColumns(num int) []table.Column {
	inningsColumns := []table.Column{}

	for i := range num {
		inningsColumns = append(inningsColumns, table.NewFlexColumn(fmt.Sprintf("%d", i+1), fmt.Sprintf("%d", i+1), 1))
	}

	return inningsColumns
}

func createLinescoreRows(awayTeamName string, homeTeamName string, linescore mlb.ScheduleLinescore) (table.Row, table.Row) {
	awayRowData := table.RowData{}
	homeRowData := table.RowData{}

	awayRowData["teamName"] = awayTeamName
	homeRowData["teamName"] = homeTeamName

	for i, inning := range linescore.Innings {
		awayRowData[fmt.Sprintf("%d", i+1)] = inning.Away.Runs
		homeRowData[fmt.Sprintf("%d", i+1)] = inning.Home.Runs
	}

	awayRowData["runs"] = linescore.Teams.Away.Runs
	awayRowData["hits"] = linescore.Teams.Away.Hits
	awayRowData["errors"] = linescore.Teams.Away.Errors
	awayRowData["lob"] = linescore.Teams.Away.LeftOnBase

	homeRowData["runs"] = linescore.Teams.Home.Runs
	homeRowData["hits"] = linescore.Teams.Home.Hits
	homeRowData["errors"] = linescore.Teams.Home.Errors
	homeRowData["lob"] = linescore.Teams.Home.LeftOnBase

	awayRow := table.NewRow(awayRowData)
	homeRow := table.NewRow(homeRowData)

	return awayRow, homeRow
}

func BuildLinescoreTable(awayTeamName string, homeTeamName string, linescore mlb.ScheduleLinescore) table.Model {
	linescoreColumns := []table.Column{
		table.NewFlexColumn("teamName", "Team", 3),
	}

	linescoreColumns = append(linescoreColumns, generateInningsColumns(len(linescore.Innings))...)

	summaryColumns := []table.Column{
		table.NewFlexColumn("runs", "R", 1),
		table.NewFlexColumn("hits", "H", 1),
		table.NewFlexColumn("errors", "E", 1),
		table.NewFlexColumn("lob", "LOB", 1),
	}

	linescoreColumns = append(linescoreColumns, summaryColumns...)

	linescoreRows := []table.Row{}

	awayLinescore, homeLinescore := createLinescoreRows(awayTeamName, homeTeamName, linescore)
	linescoreRows = append(linescoreRows, awayLinescore, homeLinescore)

	return table.New(linescoreColumns).WithRows(linescoreRows).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center))
}
