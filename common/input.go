package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// FetchInput fetches input data from Advent of Code for the given year and day
// and returns the input as a string.
func FetchInput(year int, day int) (string, error) {
	session := os.Getenv("SESSION")
	if session == "" {
		return "", fmt.Errorf("SESSION environment variable is not set")
	}

	client := &http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
