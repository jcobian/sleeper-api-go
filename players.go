package sleeper

import (
	"fmt"
)

// PlayersService deals with interacting with the Players segment of the Sleeper API
type PlayersService struct {
	client *Client
}

// GetAll gets all players
func (playersService *PlayersService) GetAll(sport string) (*ObjectResponse, error) {
	path := fmt.Sprintf("players/%s", sport)
	req, err := playersService.client.newRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	resp, err := playersService.client.doWithObjectResponse(req)
	return resp, err
}

// GetAllTrending gets all trending players
// TODO: Fix this, seems to not be returning or unmarshalling right
func (playersService *PlayersService) GetAllTrending(sport, trendingType string, lookbackHours, limit int) (*ArrayResponse, error) {
	path := fmt.Sprintf("players/%s/trending/%s?lookback_hours=%d&limit=%d", sport, trendingType, lookbackHours, limit)
	req, err := playersService.client.newRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	resp, err := playersService.client.doWithArrayResponse(req)
	return resp, err
}
