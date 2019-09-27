// Package sleeper contains ways to interact with the Sleeper API
package sleeper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.sleeper.app"
	defaultVersion = "v1"
)

// Client is the API client
type Client struct {
	// Base url for API requests
	BaseURL *url.URL

	// HTTP client used to communicate with the DO API.
	httpClient *http.Client

	Stats   *StatsService
	Players *PlayersService
}

// NewAPIClient create a Sleeper API Client
func NewAPIClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	client := &Client{BaseURL: baseURL, httpClient: httpClient}
	client.Stats = &StatsService{client: client}
	client.Players = &PlayersService{client: client}
	return client
}

func (client *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: fmt.Sprintf("/%s/%s", defaultVersion, path)}
	u := client.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ObjectResponse is a wrapper for the response from the Sleeper API for all endpoints
type ObjectResponse struct {
	// The actual response body
	Body map[string]interface{}

	// The response status code
	StatusCode int
}

// ArrayResponse is a wrapper for the response from the Sleeper API for all endpoints
type ArrayResponse struct {
	// The actual response body
	Body []interface{}

	// The response status code
	StatusCode int
}

func (client *Client) doHelper(req *http.Request) (*http.Response, []byte, error) {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("%s", body)
	}

	bodyBytes := []byte(body)
	return resp, bodyBytes, nil
}

func (client *Client) doWithObjectResponse(req *http.Request) (*ObjectResponse, error) {
	resp, bodyBytes, err := client.doHelper(req)

	var f interface{}
	jsonErr := json.Unmarshal(bodyBytes, &f)
	if jsonErr != nil {
		return nil, err
	}
	bodyInterface := f.(map[string]interface{})
	sleeperResponse := &ObjectResponse{Body: bodyInterface, StatusCode: resp.StatusCode}
	return sleeperResponse, nil
}

func (client *Client) doWithArrayResponse(req *http.Request) (*ArrayResponse, error) {
	resp, bodyBytes, err := client.doHelper(req)

	var f interface{}
	jsonErr := json.Unmarshal(bodyBytes, &f)
	if jsonErr != nil {
		return nil, err
	}
	bodyInterface := f.([]interface{})
	sleeperResponse := &ArrayResponse{Body: bodyInterface, StatusCode: resp.StatusCode}
	return sleeperResponse, nil
}
