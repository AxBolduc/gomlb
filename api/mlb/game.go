package mlb

import (
	"fmt"
)

type GameTeams struct {
	Away GameTeam
	Home GameTeam
}

type Game struct {
	GamePk    int
	GameGuid  string
	Teams     GameTeams
	Linescore ScheduleLinescore
	Lineups   ScheduleLineups
}

func (g Game) FilterValue() string {
	return fmt.Sprintf("%s v %s", g.Teams.Away.Team.Name, g.Teams.Home.Team.Name)
}

func (g Game) Title() string {
	return fmt.Sprintf("%s v %s", g.Teams.Away.Team.Name, g.Teams.Home.Team.Name)
}

func (g Game) Description() string {
	return fmt.Sprintf("(%s) - (%s)",
		g.Teams.Away.LeagueRecord.Pct,
		g.Teams.Home.LeagueRecord.Pct)
}
