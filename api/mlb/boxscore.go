package mlb

type Boxscore struct {
	Teams BoxscoreTeams
}

type BoxscoreTeams struct {
	Away BoxscoreTeam
	Home BoxscoreTeam
}

type BoxscoreTeam struct {
	Players      TeamPlayers
	Batters      []int
	Pitchers     []int
	Bench        []int
	Bullpen      []int
	BattingOrder []int
	Info         []BoxscoreInfo
	Note         []BoxscoreNote
}

type BoxscoreNote struct {
	Label string
	Value string
}

type BoxscoreInfo struct {
	FieldList []BoxscoreNote
}

type BoxscorePlayer struct {
	Person       Person
	JerseyNumber string
	Position     PlayerPosition
	Status       BoxscorePlayerStatus
	ParentTeamId int
	Stats        BoxscorePlayerStats
	SeasonStats  BoxscorePlayerSeasonStats
	GameStatus   BoxscorePlayerGameStatus
}

type BoxscorePlayerStatus struct {
	Code        string
	Description string
}

type BoxscorePlayerStats struct {
	Batting  BoxscoreBattingStats
	Pitching BoxscorePitchingStats
	Fielding BoxscoreFieldingStats
}

type BoxscoreBattingStats struct {
	Note                 string
	Summary              string
	GamesPlayed          int
	FlyOuts              int
	GroundOuts           int
	Runs                 int
	Doubles              int
	Triples              int
	HomeRuns             int
	StrikeOuts           int
	BaseOnBalls          int
	IntentionalWalks     int
	Hits                 int
	HitByPitch           int
	AtBats               int
	CaughtStealing       int
	StolenBases          int
	StolenBasePercentage string
	GroundIntoDoublePlay int
	GroundIntoTriplePlay int
	PlateAppearances     int
	TotalBases           int
	Rbi                  int
	LeftOnBase           int
	SacBunts             int
	SacFlies             int
	CatchersInterference int
	Pickoffs             int
	AtBatsPerHomeRun     string
}

type BoxscoreFieldingStats struct {
	CaughtStealing       int
	StolenBases          int
	StolenBasePercentage string
	Assists              int
	PutOuts              int
	Errors               int
	Fielding             string
	PassedBall           int
	Pickoffs             int
}

type BoxscorePitchingStats struct {
	Note                   string
	Summary                string
	GamesPlayed            int
	GamesStarted           int
	FlyOuts                int
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
	AtBats                 int
	CaughtStealing         int
	StolenBases            int
	StolenBasePercentage   string
	NumberOfPitches        int
	InningsPitched         string
	Wins                   int
	Losses                 int
	Saves                  int
	SaveOpportunities      int
	Holds                  int
	BlownSaves             int
	EarnedRuns             int
	BattersFaced           int
	Outs                   int
	GamesPitched           int
	CompleteGames          int
	Shutdouts              int
	PitchesThrown          int
	Balls                  int
	Strikes                int
	StrikePercentage       string
	HitBatsmen             int
	Balks                  int
	WildPitches            int
	Pickoffs               int
	Rbi                    int
	GamesFinished          int
	RunsScoredPer9         string
	HomeRunsPer9           string
	InheritedRunners       int
	InheritedRunnersScored int
	CatchersInterference   int
	SacBunts               int
	SacFlies               int
	PassedBall             int
}

type BoxscorePlayerSeasonStats struct {
	Batting  BoxscoreBattingSeasonStats
	Pitching BoxscorePitchingSeasonStats
	Fielding BoxscoreFieldingSeasonStats
}

type BoxscoreBattingSeasonStats struct {
	GamesPlayed          int
	FlyOuts              int
	GroundOuts           int
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
	GroundIntoTriplePlay int
	PlateAppearances     int
	TotalBases           int
	Rbi                  int
	LeftOnBase           int
	SacBunts             int
	SacFlies             int
	Babip                string
	CatchersInterference int
	Pickoffs             int
	AtBatsPerHomeRun     string
}

type BoxscoreFieldingSeasonStats struct {
	GamesStarted         int
	CaughtStealing       int
	StolenBases          int
	StolenBasePercentage string
	Assists              int
	PutOuts              int
	Errors               int
	Chances              int
	Fielding             string
	PassedBall           int
	Pickoffs             int
}

type BoxscorePitchingSeasonStats struct {
	GamesPlayed            int
	GamesStarted           int
	FlyOuts                int
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
	AtBats                 int
	Obp                    string
	CaughtStealing         int
	StolenBases            int
	StolenBasePercentage   string
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
	Shutdouts              int
	PitchesThrown          int
	Balls                  int
	Strikes                int
	StrikePercentage       string
	HitBatsmen             int
	Balks                  int
	WildPitches            int
	Pickoffs               int
	GroundOutsToAirOuts    string
	Rbi                    int
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
	PassedBall             int
}

type BoxscorePlayerGameStatus struct {
	IsCurrentBatter  bool
	IsCurrentPitcher bool
	IsOnBench        bool
	IsSubstitute     bool
}

type Person struct {
	Id       int
	FullName string
}

type TeamPlayers map[string]BoxscorePlayer
