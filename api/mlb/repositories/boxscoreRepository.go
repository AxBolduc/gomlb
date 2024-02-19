package repositories

import (
	"encoding/json"
	"fmt"
	"sports-cli/api/mlb"
)

type BoxscoreRepository struct {
	mlbClient *mlb.Client
}

func NewBoxscoreRepository() BoxscoreRepository {
	return BoxscoreRepository{
		mlbClient: mlb.NewDefaultClient(),
	}
}

func (repo BoxscoreRepository) GetBoxscoreFromGamePk(gamePk int) (*mlb.Boxscore, error) {
	responseBytes, err := repo.mlbClient.Get(fmt.Sprintf("game/%d/boxscore", gamePk), nil)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var res mlb.Boxscore
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
