package runalyze

import (
	"context"
	"net/http"
	"time"
)

// BloodPressure represents the data of a blood pressure entry
type BloodPressure struct {
	DateTime  time.Time `json:"date_time"`
	Systolic  int       `json:"systolic"`
	Diastolic int       `json:"diastolic"`
	HeartRate int       `json:"heart_rate,omitempty"`
}

// CreateBloodPressure creates a new blood pressure entry
func (c *Client) CreateBloodPressure(ctx context.Context, bp BloodPressure) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics/bloodPressure", bp)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
