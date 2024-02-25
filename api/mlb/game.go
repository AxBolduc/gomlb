package mlb

import (
	"fmt"
	"log"
	"time"
)

type GameTeams struct {
	Away GameTeam
	Home GameTeam
}

type Game struct {
	GamePk    int
	GameGuid  string
	GameDate  string
	Teams     GameTeams
	Linescore ScheduleLinescore
	Lineups   ScheduleLineups
	Status    GameStatus
}

type GameStatus struct {
	AbstractGameState string
	CodedGameState    string
	DetailedState     string
	StatusCode        string
	StartTimeTBD      bool
	AbstractGameCode  string
}

func (g Game) FilterValue() string {
	return fmt.Sprintf("%s v %s", g.Teams.Away.Team.Name, g.Teams.Home.Team.Name)
}

func (g Game) Title() string {
	date, err := time.Parse(time.RFC3339, g.GameDate)
	if err != nil {
		log.Printf("Failed to parse game time, defaulting")
		date = time.Now()
	}

	dateString := date.Local().Format(time.Kitchen)
	if g.Status.CodedGameState == "I" {
		dateString = "LIVE"
	}

	return fmt.Sprintf("%s (%d-%d) v %s (%d-%d) - (%s)",
		g.Teams.Away.Team.Name,
		g.Teams.Away.LeagueRecord.Wins,
		g.Teams.Away.LeagueRecord.Losses,
		g.Teams.Home.Team.Name,
		g.Teams.Home.LeagueRecord.Wins,
		g.Teams.Home.LeagueRecord.Losses,
		dateString,
	)
}

func (g Game) Description() string {
	// TODO figure something out to put here
	return g.Status.DetailedState
}
