package popup

import (
	"fmt"
	"log"

	"github.com/axbolduc/gomlb/api/mlb"
	"github.com/axbolduc/gomlb/api/mlb/repositories"
	tea "github.com/charmbracelet/bubbletea"
)

type IPopup interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (IPopup, tea.Cmd)
	View() string
	Resize(msg tea.WindowSizeMsg, bgRaw string) BatterStatsPopup
}

type BatterStatsPopup struct {
	popup        Popup
	hittingStats mlb.PlayerHittingStats
	batterId     int
}

func (p BatterStatsPopup) Init() tea.Cmd {
	return nil
}

func (p BatterStatsPopup) Update(msg tea.Msg) (IPopup, tea.Cmd) {
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

func (p BatterStatsPopup) View() string {

	p.popup.SetFg("Hello my name is test")
	return p.popup.View()
}

func (p BatterStatsPopup) Resize(msg tea.WindowSizeMsg, bgRaw string) BatterStatsPopup {
	p.popup = p.popup.Resize(msg, bgRaw)
	return p
}

func NewBatterStatsPopup(bg string, batterId int, width int, height int) *BatterStatsPopup {
	log.Printf("NEW BATTER STATS POPUP: %d x %d", width, height)
	statsRepo := repositories.NewPlayerStatsRepository()
	batterStats, err := statsRepo.GetCareerStatsByPlayerId(batterId)
	if err != nil {
		return nil
	}

	if len(batterStats.Stats) == 0 {
		log.Printf("No stats batting stats for player %d", batterId)
		return nil
	}

	hittingStats := batterStats.Stats[0]

	title := fmt.Sprintf("Batting stats for player %d", batterId)

	popup := NewPopup(bg, title, width, height)

	return &BatterStatsPopup{
		popup:        popup,
		batterId:     batterId,
		hittingStats: hittingStats,
	}
}
