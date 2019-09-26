package sleeper

import (
	"fmt"
	"net/http"
)

// StatsService deals with interacting with the Stats segment of the Sleeper API
type StatsService struct {
	client *Client
}

// Get gets all stats
func (statsService *StatsService) Get(sport, seasonType, season string) (*http.Response, error) {
	path := fmt.Sprintf("stats/%s/%s/%s", sport, seasonType, season)
	req, err := statsService.client.newRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	resp, err := statsService.client.do(req)
	return resp, err
}
