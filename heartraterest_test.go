package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var hrrTestCases = []struct {
	name   string
	hrr    HeartRateRest
	status int
}{
	{
		name: "valid heart rate rest data",
		hrr: HeartRateRest{
			DateTime:  time.Now(),
			HeartRate: 100,
		},
		status: http.StatusCreated,
	},
	{
		name: "invalid (too high) heart rate rest data",
		hrr: HeartRateRest{
			DateTime:  time.Now(),
			HeartRate: 999,
		},
		status: http.StatusBadRequest,
	},
	{
		name: "invalid (too low) heart rate rest data",
		hrr: HeartRateRest{
			DateTime:  time.Now(),
			HeartRate: 2,
		},
		status: http.StatusBadRequest,
	},
}

func TestCreateHeartRateRest(t *testing.T) {
	for _, tc := range hrrTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testCreateHeartRateRest(st, tc.hrr, tc.status)
		})
	}
}

func testCreateHeartRateRest(t *testing.T, hrr HeartRateRest, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics/heartRateRest", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqHrr = HeartRateRest{}
		err := json.NewDecoder(r.Body).Decode(&reqHrr)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqHrr, hrr)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateHeartRateRest(context.Background(), hrr)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
