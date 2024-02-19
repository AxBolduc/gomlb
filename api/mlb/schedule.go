package mlb

type Schedule struct {
	TotalItems int
	TotalGames int
	Dates      []ScheduleDate
}

type ScheduleDate struct {
	Date       string
	TotalItems int
	TotalGames int
	Games      []Game
}
