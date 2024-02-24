package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/axbolduc/gomlb/api/mlb"
)

type PlayerStatsRepository struct {
	client *mlb.Client
}

var playerStatsRepoInstance *PlayerStatsRepository

func NewPlayerStatsRepository() *PlayerStatsRepository {
	if playerStatsRepoInstance == nil {
		playerStatsRepoInstance = &PlayerStatsRepository{
			client: mlb.NewDefaultClient(),
		}
	}

	return playerStatsRepoInstance
}

func (r PlayerStatsRepository) GetCareerStatsByPlayerId(playerId int) (*mlb.HittingStatsResponse, error) {
	queryParams := make(map[string]string)
	queryParams["stats"] = "career"
	queryParams["group"] = "hitting"

	endpoint := fmt.Sprintf("/people/%d/stats", playerId)
	responseBytes, err := r.client.Get(endpoint, queryParams)

	if err != nil {
		log.Printf("Call to get player's carrer stats failed: %s", err.Error())
		return nil, err
	}

	// unmarshall the json

	var res mlb.HittingStatsResponse
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		log.Fatalf("Failed to parse the json response when getting career stats for player with id %d", playerId)
	}

	return &res, nil

}

func (r PlayerStatsRepository) GetYearByYearHittingStatsByPlayerId(playerId int) (*mlb.HittingStatsResponse, error) {
	queryParams := make(map[string]string)
	queryParams["stats"] = "yearByYear"
	queryParams["group"] = "hitting"

	endpoint := fmt.Sprintf("/people/%d/stats", playerId)
	responseBytes, err := r.client.Get(endpoint, queryParams)

	if err != nil {
		log.Printf("Call to get player's yearByYear stats failed: %s", err.Error())
		return nil, err
	}

	// unmarshall the json

	var res mlb.HittingStatsResponse
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		log.Fatalf("Failed to parse the json response when getting yearByYear stats for player with id %d", playerId)
	}

	return &res, nil
}
