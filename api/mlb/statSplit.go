package mlb

type HittingStatsSplit struct {
	Season   string
	Stat     HittingStats
	Player   Person
	Sport    Sport
	GameType string
	Team     Team
	NumTeams int
}

type PitchingStatsSplit struct {
	Season   string
	Stat     PitchingStats
	Player   Person
	Sport    Sport
	GameType string
	Team     Team
	NumTeams int
}

type FieldingStatsSplit struct {
	Season   string
	Stat     FieldingStats
	Player   Person
	Sport    Sport
	GameType string
	Team     Team
	NumTeams int
}

type HittingStats struct {
	GamesPlayed          int
	GroundOuts           int
	AirOuts              int
	Runs                 int
	Doubles              int
	Triples              int
	HomeRuns             int
	StrikeOuts           int
	BaseOnBalls          int
	IntentionalWalks     int
	Hits                 int
	HitByPitch           int
	Avg                  string
	AtBats               int
	Obp                  string
	Slg                  string
	Ops                  string
	CaughtStealing       int
	StolenBases          int
	StolenBasePercentage string
	GroundIntoDoublePlay int
	NumberOfPitches      int
	PlateAppearances     int
	TotalBases           int
	Rbi                  int
	LeftOnBase           int
	SacBunts             int
	SacFlies             int
	Babip                string
	GroundOutsToAirouts  string
	CatchersInterference int
	AtBatsPerHomeRun     string
}

type FieldingStats struct {
	GamesPlayed        int
	GamesStarted       int
	Assists            int
	PutOuts            int
	Errors             int
	Chances            int
	Fielding           string
	Position           PlayerPosition
	RangeFactorPerGame string
	RangeFactorPer9Inn string
	Innings            string
	DoublePlays        int
	TriplePlays        int
	ThrowingErrors     int
}

type PitchingStats struct {
	GamesPlayed            int
	GamesStarted           int
	GroundOuts             int
	AirOuts                int
	Runs                   int
	Doubles                int
	Triples                int
	HomeRuns               int
	StrikeOuts             int
	BaseOnBalls            int
	IntentionalWalks       int
	Hits                   int
	HitByPitch             int
	Avg                    string
	AtBats                 int
	Obp                    string
	Slg                    string
	Ops                    string
	CaughtStealing         int
	StolenBases            int
	StolenBasePercentage   string
	GroundIntoDoublePlay   int
	NumberOfPitches        int
	Era                    string
	InningsPitched         string
	Wins                   int
	Losses                 int
	Saves                  int
	SaveOpportunities      int
	Holds                  int
	BlownSaves             int
	EarnedRuns             int
	Whip                   string
	BattersFaced           int
	Outs                   int
	GamesPitched           int
	CompleteGames          int
	Shutouts               int
	Strikes                int
	StrikePercentage       string
	HitBatsmen             int
	Balks                  int
	WildPitches            int
	Pickoffs               int
	TotalBases             int
	GroundOutsToAirouts    string
	WinPercentage          string
	PitchesPerInning       string
	GamesFinished          int
	StrikeoutWalkRatio     string
	StrikeoutsPer9Inn      string
	WalksPer9Inn           string
	HitsPer9Inn            string
	RunsScoredPer9         string
	HomeRunsPer9           string
	InheritedRunners       int
	InheritedRunnersScored int
	CatchersInterference   int
	SacBunts               int
	SacFlies               int
}

type Sport struct {
	Id           int
	Abbreviation string
}
