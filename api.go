// Package sleeper contains ways to interact with the Sleeper API
package sleeper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	Stats *StatsService
}

// NewAPIClient create a Sleeper API Client
func NewAPIClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	client := &Client{BaseURL: baseURL, httpClient: httpClient}
	client.Stats = &StatsService{client: client}
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

func (client *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
