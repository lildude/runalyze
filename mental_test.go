package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var msTestCases = []struct {
	name   string
	ms     MentalState
	status int
}{
	{
		name: "full mental state data",
		ms: MentalState{
			DateTime: time.Now(),
			Fatigue:  8,
			Stress:   5,
			Mood:     2,
		},
		status: http.StatusCreated,
	},
	{
		name: "invalid (fatigue too low) mental state data",
		ms: MentalState{
			DateTime: time.Now(),
			Fatigue:  -1,
			Stress:   5,
			Mood:     2,
		},
		status: http.StatusBadRequest,
	},
	{
		name: "invalid (Stress too high) mental state data",
		ms: MentalState{
			DateTime: time.Now(),
			Fatigue:  -1,
			Stress:   11,
			Mood:     2,
		},
		status: http.StatusBadRequest,
	},
}

func TestCreateMental(t *testing.T) {
	for _, tc := range msTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testCreateMental(st, tc.ms, tc.status)
		})
	}
}

func testCreateMental(t *testing.T, ms MentalState, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics/mental", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqMs = MentalState{}
		err := json.NewDecoder(r.Body).Decode(&reqMs)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqMs, ms)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateMental(context.Background(), ms)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
