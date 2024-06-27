package dateJumpPopup

import (
	"github.com/axbolduc/gomlb/ui/popup"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	popup                     popup.Popup
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

	fg := lipgloss.JoinVertical(lipgloss.Left, "Jump To Date", "\n", "JUMP")

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

func New(bg string, width int, height int) Model {
	title := "Jump To Date"
	popup := popup.NewPopup(bg, title, width, height)

	return Model{
		popup:        popup,
		windowWidth:  width - 2,
		windowHeight: height - 2,
	}
}
