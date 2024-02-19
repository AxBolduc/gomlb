package ui

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTea(date time.Time) {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	// Initialize model
	m := InitModel(date)

	// Run ui with model
	t := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := t.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
