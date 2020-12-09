package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var metricTestCases = []struct {
	name   string
	m      Metrics
	status int
}{
	{
		name: "full metrics data",
		m: Metrics{
			Sleep: Sleep{
				DateTime:           time.Now(),
				Duration:           460,
				RemDuration:        80,
				LightSleepDuration: 70,
				DeepSleepDuration:  70,
				AwakeDuration:      70,
				Quality:            8,
			},
			BodyComposition: BodyComposition{
				DateTime:         time.Now(),
				Weight:           82.5,
				FatPercentage:    11.23,
				WaterPercentage:  1.5,
				MusclePercentage: 19.7,
				BonePercentage:   98,
			},
			BloodPressure: BloodPressure{
				DateTime:  time.Now(),
				Systolic:  120,
				Diastolic: 80,
				HeartRate: 70,
			},
			HeartRateRest: HeartRateRest{
				DateTime:  time.Now(),
				HeartRate: 100,
			},
			HeartRateMax: HeartRateMax{
				DateTime:  time.Now(),
				HeartRate: 192,
			},
			MentalState: MentalState{
				DateTime: time.Now(),
				Fatigue:  8,
				Stress:   5,
				Mood:     2,
			},
		},
		status: http.StatusCreated,
	},
	{
		name: "partial metrics data",
		m: Metrics{
			Sleep: Sleep{
				DateTime: time.Now(),
				Duration: 460,
			},
			BodyComposition: BodyComposition{
				DateTime: time.Now(),
				Weight:   82.5,
			},
			HeartRateRest: HeartRateRest{
				DateTime:  time.Now(),
				HeartRate: 100,
			},
			HeartRateMax: HeartRateMax{
				DateTime:  time.Now(),
				HeartRate: 192,
			},
		},
		status: http.StatusCreated,
	},
	{
		name: "partial metrics data with invalid entry",
		m: Metrics{
			HeartRateMax: HeartRateMax{
				DateTime:  time.Now(),
				HeartRate: 999,
			},
		},
		status: http.StatusBadRequest,
	},
}

func TestMetrics(t *testing.T) {
	for _, tc := range metricTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testMetrics(st, tc.m, tc.status)
		})
	}
}

func testMetrics(t *testing.T, m Metrics, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqMetrics = Metrics{}
		err := json.NewDecoder(r.Body).Decode(&reqMetrics)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqMetrics, m)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateMetrics(context.Background(), m)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
