// Credit to @TypicalAM https://github.com/TypicalAM/goread
package popup

import (
	"github.com/axbolduc/gomlb/ui/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type IPopup interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (IPopup, tea.Cmd)
	View() string
	Resize(msg tea.WindowSizeMsg, bgRaw string) IPopup
}

// Popup is the feed popup where a user can create/edit a feed.
type Popup struct {
	title   string
	fg      string
	style   popupStyle
	overlay Overlay
}

// NewPopup returns a new feed popup.
func NewPopup(bgRaw string, title string, width, height int) Popup {
	style := newPopupStyle(width, height)
	overlay := NewOverlay(bgRaw, width, height)

	return Popup{
		title:   title,
		overlay: overlay,
		style:   style,
	}
}

// Init initializes the popup.
func (p Popup) Init() tea.Cmd {
	return nil
}

// Update updates the popup.
func (p Popup) Update(msg tea.Msg) (Popup, tea.Cmd) {
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

func (p Popup) Resize(msg tea.WindowSizeMsg, bgRaw string) Popup {
	newWidth := msg.Width - constants.DocStyle.GetHorizontalFrameSize() - p.style.general.GetVerticalBorderSize() - 2*constants.PopupHPadding
	newHeight := msg.Height - constants.DocStyle.GetVerticalFrameSize() - p.style.general.GetHorizontalBorderSize() - 2*constants.PopupVPadding

	p.overlay = NewOverlay(bgRaw, newWidth, newHeight)

	p.style.general = p.style.general.Width(newWidth).Height(newHeight)
	p.style.heading = p.style.heading.Width(newWidth)
	p.style.item = p.style.item.Width(newWidth).MaxHeight(p.style.general.GetHeight() - p.style.heading.GetHeight())

	return p
}

// View renders the popup.
func (p Popup) View() string {
	title := p.style.heading.Render(p.title)
	fgRender := p.style.item.Render(p.fg)
	popup := lipgloss.JoinVertical(lipgloss.Left, title, fgRender)

	return p.overlay.WrapView(p.style.general.Render(popup))
}

func (p Popup) GetStyles() popupStyle {
	return p.style
}

func (p Popup) SetFg(fg string) Popup {
	p.fg = fg
	return p
}
