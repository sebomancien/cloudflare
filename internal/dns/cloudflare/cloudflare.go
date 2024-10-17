package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Email  string
	ApiKey string
	ZoneId string
}

type Record struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

func GetRecords(config *Config) ([]Record, error) {
	var records []Record
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/", config.ZoneId)
	err := send(http.MethodGet, url, nil, &records, config)
	return records, err
}

func PutRecord(record Record, config *Config) error {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s/", config.ZoneId, record.Id)
	err := send[interface{}](http.MethodPut, url, record, nil, config)
	return err
}

func send[T any](method string, url string, request any, response *T, config *Config) error {
	// Convert the JSON into a request body
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// Create the request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Add the header identification
	req.Header.Set("X-Auth-Email", config.Email)
	req.Header.Set("X-Auth-Key", config.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch DNS records, status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Define a common structure for all responses body
	var data struct {
		Results T    `json:"result"`
		Success bool `json:"success"`
		Errors  []struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
		Messages []struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"messages"`
	}

	// Parse the response JSON
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	if !data.Success {
		return fmt.Errorf("server error: %v", data.Errors)
	}

	if response != nil {
		*response = data.Results
	}

	return nil
}
