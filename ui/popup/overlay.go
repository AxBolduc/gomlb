// Credit to @TypicalAM https://github.com/TypicalAM/goread
package popup

import (
	"strings"

	"github.com/muesli/ansi"
)

// Overlay allows you to overlay text on top of a background and achieve a popup.
type Overlay struct {
	bgRaw     string
	textAbove string
	textBelow string
	rowPrefix []string
	rowSuffix []string
	width     int
	height    int
}

// NewOverlay creates a new overlay and computes the necessary indices.
func NewOverlay(bgRaw string, popupWidth, popupHeight int) Overlay {
	bg := strings.Split(bgRaw, "\n")
	bgWidth := ansi.PrintableRuneWidth(bg[0])
	bgHeight := len(bg)

	if popupHeight > bgHeight {
		popupHeight = bgHeight
	}
	if popupWidth > bgWidth {
		popupWidth = bgWidth
	}

	startRow := (bgHeight - popupHeight) / 2
	startCol := (bgWidth - popupWidth) / 2

	rowPrefix := make([]string, popupHeight)
	rowSuffix := make([]string, popupHeight)

	for i, text := range bg[startRow : startRow+popupHeight] {
		popupStart := findPrintIndex(text, startCol)
		popupEnd := findPrintIndex(text, startCol+popupWidth)

		if popupStart != -1 {
			rowPrefix[i] = text[:popupStart]
		} else {
			rowPrintable := ansi.PrintableRuneWidth(text)
			rowPrefix[i] = text + strings.Repeat(" ", startCol-rowPrintable)
		}

		if popupEnd != -1 {
			rowSuffix[i] = text[popupEnd:]
		} else {
			rowSuffix[i] = ""
		}
	}

	prefix := strings.Join(bg[:startRow], "\n")
	suffix := strings.Join(bg[startRow+popupHeight:], "\n")

	return Overlay{
		bgRaw:     bgRaw,
		rowPrefix: rowPrefix,
		rowSuffix: rowSuffix,
		width:     popupWidth,
		height:    popupHeight,
		textAbove: prefix,
		textBelow: suffix,
	}
}

// WrapView overlays the given text on top of the background.
// TODO: Maybe handle the box here. It's a bit weird to have to do it in the view.
func (p Overlay) WrapView(view string) string {
	var b strings.Builder
	b.WriteString(p.textAbove)
	b.WriteRune('\n')

	lines := strings.Split(view, "\n")
	for i := 0; i < len(lines) && i < p.height; i++ {
		b.WriteString(p.rowPrefix[i])
		b.WriteString(lines[i])
		b.WriteString(p.rowSuffix[i])
		b.WriteRune('\n')
	}

	b.WriteString(p.textBelow)
	return b.String()
}

// Width returns the width of the popup window.
func (p Overlay) Width() int {
	return p.width
}

// Height returns the height of the popup window.
func (p Overlay) Height() int {
	return p.height
}

// findPrintIndex finds the print index, that is what string index corresponds to the given printable rune index.
func findPrintIndex(str string, index int) int {
	for i := len(str) - 1; i >= 0; i-- {
		if ansi.PrintableRuneWidth(str[:i]) == index {
			return i
		}
	}

	return -1
}
