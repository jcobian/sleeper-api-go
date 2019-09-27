package sleeper

import (
	"fmt"
)

// StatsService deals with interacting with the Stats segment of the Sleeper API
type StatsService struct {
	client *Client
}

// GetAllStats gets all stats
func (statsService *StatsService) GetAllStats(sport, seasonType, season string) (*ObjectResponse, error) {
	path := fmt.Sprintf("stats/%s/%s/%s", sport, seasonType, season)
	req, err := statsService.client.newRequest("GET", path)

	if err != nil {
		return nil, err
	}

	resp, err := statsService.client.doWithObjectResponse(req)
	return resp, err
}

// GetStatsForWeek gets stats for a given week
func (statsService *StatsService) GetStatsForWeek(sport, seasonType, season, week string) (*ObjectResponse, error) {
	path := fmt.Sprintf("stats/%s/%s/%s/%s", sport, seasonType, season, week)
	req, err := statsService.client.newRequest("GET", path)

	if err != nil {
		return nil, err
	}

	resp, err := statsService.client.doWithObjectResponse(req)
	return resp, err
}

// GetAllProjections gets all stats
func (statsService *StatsService) GetAllProjections(sport, seasonType, season string) (*ObjectResponse, error) {
	path := fmt.Sprintf("projections/%s/%s/%s", sport, seasonType, season)
	req, err := statsService.client.newRequest("GET", path)

	if err != nil {
		return nil, err
	}

	resp, err := statsService.client.doWithObjectResponse(req)
	return resp, err
}

// GetProjectionsForWeek gets stats for a given week
func (statsService *StatsService) GetProjectionsForWeek(sport, seasonType, season, week string) (*ObjectResponse, error) {
	path := fmt.Sprintf("projections/%s/%s/%s/%s", sport, seasonType, season, week)
	req, err := statsService.client.newRequest("GET", path)

	if err != nil {
		return nil, err
	}

	resp, err := statsService.client.doWithObjectResponse(req)
	return resp, err
}
