package runalyze

import (
	"context"
	"net/http"
	"time"
)

// MentalState represents the data of a mental state entry
type MentalState struct {
	DateTime time.Time `json:"date_time"`
	Fatigue  int       `json:"fatigue"`
	Stress   int       `json:"stress"`
	Mood     int       `json:"mood"`
}

// CreateMental creates a new mental state entry
func (c *Client) CreateMental(ctx context.Context, m MentalState) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics/mental", m)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
