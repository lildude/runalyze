package runalyze

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var bcTestCases = []struct {
	name   string
	bc     BodyComposition
	status int
}{
	{
		name: "full body composition data",
		bc: BodyComposition{
			DateTime:         time.Now(),
			Weight:           82.5,
			FatPercentage:    11.23,
			WaterPercentage:  1.5,
			MusclePercentage: 19.7,
			BonePercentage:   98,
		},
		status: http.StatusCreated,
	},
	{
		name: "partial body composition data",
		bc: BodyComposition{
			DateTime:      time.Now(),
			Weight:        82.5,
			FatPercentage: 11.23,
		},
		status: http.StatusCreated,
	},
	{
		name: "invalid body composition data",
		bc: BodyComposition{
			DateTime:      time.Now(),
			FatPercentage: 11.23, // Weight is required
		},
		status: http.StatusBadRequest,
	},
}

func TestCreateBodyComposition(t *testing.T) {
	for _, tc := range bcTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testBodyComposition(st, tc.bc, tc.status)
		})
	}
}

func testBodyComposition(t *testing.T, bc BodyComposition, status int) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics/bodyComposition", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var reqBc = BodyComposition{}
		err := json.NewDecoder(r.Body).Decode(&reqBc)
		assert.NoError(t, err, "should send a request that the API can parse")
		assert.ObjectsAreEqual(reqBc, bc)

		w.WriteHeader(http.StatusCreated)
	})

	resp, err := client.CreateBodyComposition(context.Background(), bc)
	assert.NoError(t, err, "should not produce an error")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
