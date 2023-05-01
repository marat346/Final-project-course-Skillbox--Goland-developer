package sps

import (
	"encoding/json"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// SupportRequest - sends the request to the specified address.
// Gets data and makes a list.
// Checks if the data is correct.
func SupportRequest(addr string) ([]SupportData, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return []SupportData{}, err
	}

	if resp.StatusCode != 200 {
		return []SupportData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []SupportData{}, err
	}

	data := make([]SupportData, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []SupportData{}, err
	}

	return data, nil
}
