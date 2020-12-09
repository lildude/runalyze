package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var bpTestCases = []struct {
	name   string
	bp     BloodPressure
	status int
}{
	{
		name: "full blood pressure data",
		bp: BloodPressure{
			DateTime:  time.Now(),
			Systolic:  120,
			Diastolic: 80,
			HeartRate: 70,
		},
		status: http.StatusCreated,
	},
	{
		name: "partial blood pressure data",
		bp: BloodPressure{
			DateTime:  time.Now(),
			Systolic:  120,
			Diastolic: 80,
		},
		status: http.StatusCreated,
	},
	{
		name: "missing required blood pressure data",
		bp: BloodPressure{
			DateTime:  time.Now(),
			Diastolic: 80,
		},
		status: http.StatusBadRequest,
	},
}

func TestCreateBloodPressure(t *testing.T) {
	for _, tc := range bpTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testCreateBloodPressure(st, tc.bp, tc.status)
		})
	}
}

func testCreateBloodPressure(t *testing.T, bp BloodPressure, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics/bloodPressure", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqBp = BloodPressure{}
		err := json.NewDecoder(r.Body).Decode(&reqBp)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqBp, bp)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateBloodPressure(context.Background(), bp)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
