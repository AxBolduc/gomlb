package mlb

type GameTeam struct {
	Score        int
	LeagueRecord TeamRecord
	Team         struct {
		Name string
	}
}

type TeamRecord struct {
	Wins   int
	Losses int
	Pct    string
}

type Team struct {
	Id   int
	Name string
}
