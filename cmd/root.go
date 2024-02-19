package cmd

import (
	"os"
	"time"

	"github.com/axbolduc/gomlb/ui"

	"github.com/spf13/cobra"
)

var inputDate string

var rootCmd = &cobra.Command{
	Use:   "gomlb",
	Short: "TUI for MLB",
	Long:  "TUI application for viewing live and past MLB games and statistics",
	Run: func(cmd *cobra.Command, args []string) {

		var date time.Time
		var err error
		if inputDate == "" {
			date = time.Now()
		} else {
			date, err = time.Parse(time.DateOnly, inputDate)
			if err != nil {
				panic("Invalid date")
			}
		}

		ui.StartTea(date)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputDate, "date", "d", "", "Date to get the schedule for (YYYY-MM-DD)")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
