package runalyze

import (
	"context"
	"net/http"
	"time"
)

// HeartRateMax represents the data of a max heart rate entry
type HeartRateMax struct {
	DateTime  time.Time `json:"date_time"`
	HeartRate int       `json:"heart_rate"`
}

// CreateHeartRateMax create a new max heart rate entry
func (c *Client) CreateHeartRateMax(ctx context.Context, hrm HeartRateMax) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics/heartRateMax", hrm)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
