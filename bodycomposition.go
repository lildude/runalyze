package runalyze

import (
	"context"
	"net/http"
	"time"
)

// BodyComposition represents the data of a body composition entry
type BodyComposition struct {
	DateTime         time.Time `json:"date_time"`
	Weight           float32   `json:"weight"`
	FatPercentage    float32   `json:"fat_percentage,omitempty"`
	WaterPercentage  float32   `json:"water_percentage,omitempty"`
	MusclePercentage float32   `json:"muscle_percentage,omitempty"`
	BonePercentage   float32   `json:"bone_percentage,omitempty"`
}

// CreateBodyComposition creates a new body composition entry
func (c *Client) CreateBodyComposition(ctx context.Context, bc BodyComposition) (*http.Response, error) {
	req, err := c.NewRequest("POST", "metrics/bodyComposition", bc)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
