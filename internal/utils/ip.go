package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

// Get the current public IP address
func GetPublicIp() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ifconfig.me", nil)
	if err != nil {
		return "", err
	}

	// Mimic curl to avoid the website returning the whole HTML page
	req.Header.Set("User-Agent", "curl/7.64.1")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Validate that the IP address is in the correct format
	ip := net.ParseIP(string(body))
	if ip == nil {
		return "", fmt.Errorf("failed to parse ip address: %s", body)
	}

	return ip.String(), nil
}
