// Credit to @TypicalAM https://github.com/TypicalAM/goread
package popup

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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
func (p Popup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "q":
			return p, tea.Quit
		}
	}

	return p, tea.Batch(cmds...)
}

// View renders the popup.
func (p Popup) View() string {
	title := p.style.heading.Render(p.title)
	fgRender := p.style.item.Render(p.fg)
	popup := lipgloss.JoinVertical(lipgloss.Left, title, fgRender)
	return p.overlay.WrapView(p.style.general.Render(popup))
}

func (p *Popup) SetFg(fg string) {
	p.fg = fg
}
