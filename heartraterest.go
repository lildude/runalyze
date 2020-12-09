package runalyze

import (
	"context"
	"net/http"
	"time"
)

// HeartRateRest represents the data of a max heart rate entry
type HeartRateRest struct {
	DateTime  time.Time `json:"date_time"`
	HeartRate int       `json:"heart_rate"`
}

// CreateHeartRateRest create a new max heart rate entry
func (c *Client) CreateHeartRateRest(ctx context.Context, hrr HeartRateRest) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics/heartRateRest", hrr)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
