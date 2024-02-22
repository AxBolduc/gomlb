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

// newPopupStyle creates a new popup style.
func newPopupStyle(width, height int) popupStyle {
	general := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(width - 2).
		Height(height - 2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(constants.PrimaryColor)

	heading := lipgloss.NewStyle().
		Margin(1, 0, 1, 0).
		Width(width - 2).
		Underline(true).
		Align(lipgloss.Center)

	item := lipgloss.NewStyle().
		Margin(0, 4).
		PaddingLeft(1).
		MaxHeight(general.GetHeight() - heading.GetHeight())

	return popupStyle{
		general: general,
		heading: heading,
		item:    item,
	}
}
