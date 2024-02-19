package mlb

type ScheduleLineups struct {
	HomePlayers []LineupPlayer
	AwayPlayers []LineupPlayer
}

type LineupPlayer struct {
	Id              int
	FullName        string
	FirstName       string
	LastName        string
	UseName         string
	PrimaryPosition PlayerPosition
}

type PlayerPosition struct {
	Code         string
	Name         string
	Type         string
	Abbreviation string
}
