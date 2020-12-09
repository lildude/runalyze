package runalyze

import (
	"context"
	"net/http"
	"time"
)

// Sleep represents the data of a sleep entry
type Sleep struct {
	DateTime           time.Time `json:"date_time"`
	Duration           int       `json:"duration"`
	RemDuration        int       `json:"rem_duration,omitempty"`
	LightSleepDuration int       `json:"light_sleep_duration,omitempty"`
	DeepSleepDuration  int       `json:"deep_sleep_duration,omitempty"`
	AwakeDuration      int       `json:"awake_duration,omitempty"`
	Quality            int       `json:"quality,omitempty"`
}

// CreateSleep creates a new sleeping entry
func (c *Client) CreateSleep(ctx context.Context, s Sleep) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics/sleep", s)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
