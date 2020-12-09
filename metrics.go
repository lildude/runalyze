package runalyze

import (
	"context"
	"net/http"
)

// Metrics represents the data of all pssible metrics
type Metrics struct {
	Sleep           Sleep           `json:"sleep,omitempty"`
	BodyComposition BodyComposition `json:"bodyComposition,omitempty"`
	BloodPressure   BloodPressure   `json:"bloodPressure,omitempty"`
	HeartRateRest   HeartRateRest   `json:"heartRateRest,omitempty"`
	HeartRateMax    HeartRateMax    `json:"heartRateMax,omitempty"`
	MentalState     MentalState     `json:"mental,omitempty"`
}

// CreateMetrics creates a new metrics entry
func (c *Client) CreateMetrics(ctx context.Context, m Metrics) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics", m)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
