package repositories

import (
	"encoding/json"
	"gomlb/api/mlb"
	"time"
)

var instance *ScheduleRepository

type ScheduleRepository struct {
	client *mlb.Client
}

func NewScheduleRepository() *ScheduleRepository {
	if instance == nil {
		instance = &ScheduleRepository{
			client: mlb.NewDefaultClient(),
		}
	}

	return instance
}

func (repo *ScheduleRepository) GetScheduleForDate(date time.Time) (*mlb.Schedule, error) {
	queryParams := make(map[string]string)
	queryParams["hydrate"] = "linescore,lineups"
	queryParams["date"] = date.Format(time.DateOnly)

	responseBytes, err := repo.client.Get("/schedule/games", queryParams)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var res mlb.Schedule
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		panic(err)
	}

	return &res, nil

}
