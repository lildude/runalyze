package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var hrmTestCase = []struct {
	name   string
	hrm    HeartRateMax
	status int
}{
	{
		name: "valid heart rate max data",
		hrm: HeartRateMax{
			DateTime:  time.Now(),
			HeartRate: 192,
		},
		status: http.StatusCreated,
	},
	{
		name: "invalid (too high) heart rate max data",
		hrm: HeartRateMax{
			DateTime:  time.Now(),
			HeartRate: 999,
		},
		status: http.StatusBadRequest,
	},
	{
		name: "invalid (too low) heart rate max data",
		hrm: HeartRateMax{
			DateTime:  time.Now(),
			HeartRate: 2,
		},
		status: http.StatusBadRequest,
	},
}

func TestCreateHeartRateMax(t *testing.T) {
	for _, tc := range hrmTestCase {
		t.Run(tc.name, func(st *testing.T) {
			testCreateHeartRateMax(st, tc.hrm, tc.status)
		})
	}
}

func testCreateHeartRateMax(t *testing.T, hrm HeartRateMax, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics/heartRateMax", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqHrm = HeartRateMax{}
		err := json.NewDecoder(r.Body).Decode(&reqHrm)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqHrm, hrm)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateHeartRateMax(context.Background(), hrm)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
