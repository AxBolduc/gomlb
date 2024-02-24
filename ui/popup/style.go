// Credit to @TypicalAM https://github.com/TypicalAM/goread/
package popup

import (
	"github.com/axbolduc/gomlb/ui/constants"
	"github.com/charmbracelet/lipgloss"
)

// popupStyle is the style of the popup window.
type popupStyle struct {
	general lipgloss.Style
	heading lipgloss.Style
	item    lipgloss.Style
}

func (p popupStyle) GetGeneral() lipgloss.Style {
	return p.general
}
func (p popupStyle) GetHeading() lipgloss.Style {
	return p.heading
}
func (p popupStyle) GetItem() lipgloss.Style {
	return p.item
}

// newPopupStyle creates a new popup style.
// (-2) on width and height is for border characters
func newPopupStyle(width, height int) popupStyle {
	general := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(width - 2).
		Height(height - 2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(constants.PrimaryColor)

	heading := lipgloss.NewStyle().
		Margin(1, 0, 1, 0).
		Width(width).
		Underline(true).
		Align(lipgloss.Center)

	item := lipgloss.NewStyle().
		MaxHeight(general.GetHeight() - heading.GetHeight())

	return popupStyle{
		general: general,
		heading: heading,
		item:    item,
	}
}
