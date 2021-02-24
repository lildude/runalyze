package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var sleepTestCases = []struct {
	name   string
	sleep  Sleep
	status int
}{
	{
		name: "full sleep data",
		sleep: Sleep{
			DateTime:           time.Now(),
			Duration:           460,
			RemDuration:        80,
			LightSleepDuration: 70,
			DeepSleepDuration:  70,
			AwakeDuration:      70,
			HrAverage:          89,
			HrLowest:           60,
			Quality:            8,
		},
		status: http.StatusCreated,
	},
	{
		name: "partial sleep data",
		sleep: Sleep{
			DateTime: time.Now(),
			Duration: 460,
			Quality:  2,
		},
		status: http.StatusCreated,
	},
	{
		name: "invalid (too high) sleep data",
		sleep: Sleep{
			DateTime: time.Now(),
			Duration: 1460,
		},
		status: http.StatusBadRequest,
	},
}

func TestSleep(t *testing.T) {
	for _, tc := range sleepTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testSleep(st, tc.sleep, tc.status)
		})
	}
}

func testSleep(t *testing.T, sleep Sleep, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics/sleep", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqSleep = Sleep{}
		err := json.NewDecoder(r.Body).Decode(&reqSleep)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqSleep, sleep)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateSleep(context.Background(), sleep)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
