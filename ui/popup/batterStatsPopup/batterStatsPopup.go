package batterStatsPopup

import (
	"fmt"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/ui/popup"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	playerName                string
	popup                     popup.Popup
	careerHittingStats        *mlb.HittingStats
	yearByYearHittingStats    *mlb.PlayerHittingStats
	batterId                  int
	windowWidth, windowHeight int
}

func (p Model) Init() tea.Cmd {
	return nil
}

func (p Model) Update(msg tea.Msg) (popup.IPopup, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return p, tea.Quit
		}
	}

	return p, tea.Batch(cmds...)
}

func (p Model) View() string {
	// Career hitting stats table
	tableHPadding := 4
	tableWidth := p.windowWidth - 2*tableHPadding
	tableContainer := lipgloss.NewStyle().Width(tableWidth).Margin(0, tableHPadding)

	if p.careerHittingStats == nil || p.yearByYearHittingStats == nil {
		p.popup = p.popup.SetFg("No Batting Stats for player")
		return p.popup.View()
	}

	careerHittingStatsTable := lipgloss.JoinVertical(lipgloss.Left, "Career", buildBattingStatsTable(*p.careerHittingStats, tableWidth).View())
	yearByYearHittingStatsTable := lipgloss.JoinVertical(lipgloss.Left, "Year-By-Year", buildYearByYearHittingStatsTable(p.yearByYearHittingStats.Splits, tableWidth).View())

	fg := lipgloss.JoinVertical(lipgloss.Left, careerHittingStatsTable, "\n", yearByYearHittingStatsTable)

	p.popup = p.popup.SetFg(tableContainer.Render(fg))
	return p.popup.View()
}

func (p Model) Resize(msg tea.WindowSizeMsg, bgRaw string) popup.IPopup {
	p.popup = p.popup.Resize(msg, bgRaw)

	generalStyle := p.popup.GetStyles().GetGeneral()

	p.windowWidth = generalStyle.GetWidth() - generalStyle.GetVerticalBorderSize()
	p.windowHeight = generalStyle.GetHeight() - generalStyle.GetHorizontalBorderSize()

	return p
}

func New(bg string, batterId int, width int, height int) Model {
	hittingStatsResponse := getCareerHittingStats(batterId)

	playerName := hittingStatsResponse.Player.FullName
	hittingStats := &hittingStatsResponse.Stat

	title := fmt.Sprintf("Batting stats for %s", playerName)

	yearByYearHittingStats := getYearByYearHittingStats(batterId)

	popup := popup.NewPopup(bg, title, width, height)

	return Model{
		popup:                  popup,
		batterId:               batterId,
		playerName:             playerName,
		careerHittingStats:     hittingStats,
		yearByYearHittingStats: yearByYearHittingStats,
		windowWidth:            width - 2,
		windowHeight:           height - 2,
	}
}
