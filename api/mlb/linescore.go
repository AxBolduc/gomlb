package mlb

type ScheduleLinescore struct {
	CurrentInning int
	IsTopInning   bool
	Innings       []LinescoreInning
	Teams         LinescoreTeams
	Balls         int
	Strikes       int
	Outs          int
}

type LinescoreInning struct {
	Num  int
	Home struct {
		Runs       int
		Hits       int
		Errors     int
		LeftOnBase int
	}
	Away struct {
		Runs       int
		Hits       int
		Errors     int
		LeftOnBase int
	}
}

type LinescoreTeams struct {
	Home LinescoreTeam
	Away LinescoreTeam
}

type LinescoreTeam struct {
	Runs       int
	Hits       int
	Errors     int
	LeftOnBase int
	IsWinner   bool
}
