package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const UPTIMEROBOT_API_URL = "https://api.uptimerobot.com/v2"

type GetMetricsResponse struct {
	Stat       string     `json:"stat"`
	Pagination Pagination `json:"pagination"`
	Monitors   []Monitor  `json:"monitors"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Monitor struct {
	ID           int    `json:"id"`
	FriendlyName string `json:"friendly_name"`
	URL          string `json:"url"`
	Type         int    `json:"type"`
	Status       int    `json:"status"`
}

// UptimerobotClient is a client for Uptimerobot API
type UptimerobotClient struct {
	APIKey string
	Client *http.Client
}

// NewUptimerobotClient returns a new UptimerobotClient
func NewUptimerobotClient(apiKey string) *UptimerobotClient {
	return &UptimerobotClient{
		APIKey: apiKey,
		Client: &http.Client{},
	}
}

// GetMonitors returns a list of monitors
func (c *UptimerobotClient) GetMonitors() ([]Monitor, error) {
	url := fmt.Sprintf("%s/getMonitors?api_key=%s&format=json", UPTIMEROBOT_API_URL, c.APIKey)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(b))
	}

	var decoded GetMetricsResponse
	err = json.NewDecoder(res.Body).Decode(&decoded)
	if err != nil {
		return nil, err
	}

	return decoded.Monitors, nil
}
