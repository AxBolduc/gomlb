package cmd

import (
	"gomlb/ui"
	"time"

	"github.com/spf13/cobra"
)

var inputDate = ""

var mlbCmd = &cobra.Command{
	Use:   "mlb",
	Short: "Get MLB games",
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
	rootCmd.AddCommand(mlbCmd)
	rootCmd.PersistentFlags().StringVarP(&inputDate, "date", "d", "", "Date to get the schedule for (YYYY-MM-DD)")
}
