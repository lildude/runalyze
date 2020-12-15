package runalyze_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lildude/runalyze"
)

// Note: the examples listed here are compiled but not executed while testing.
// See the documentation on [Testing](https://golang.org/pkg/testing/#hdr-Examples)
// for further details.

func Example_sleep() {
	godotenv.Load(".env")
	ctx := context.Background()
	cfg := runalyze.Configuration{
		AppName: "go-runalyze Testing",
		Token: os.Getenv("RUNALYZE_ACCESS_TOKEN"),
	}
	cl := runalyze.NewClient(cfg)

	startSleep, _ := time.Parse(time.RFC3339, "2020-11-07T23:00:00Z")
	sleep := runalyze.Sleep{
		DateTime:           startSleep,
		Duration:           460,
		RemDuration:        80,
		LightSleepDuration: 70,
		DeepSleepDuration:  70,
		AwakeDuration:      70,
		Quality:            8,
	}
	resp, err := cl.CreateSleep(ctx, sleep)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Status)
}

func Example_bloodPressure() {
	godotenv.Load(".env")
	ctx := context.Background()
	cfg := runalyze.Configuration{
		AppName: "go-runalyze Testing",
		Token: os.Getenv("RUNALYZE_ACCESS_TOKEN"),
	}
	cl := runalyze.NewClient(cfg)

	date, _ := time.Parse(time.RFC3339, "2020-11-07T23:00:00Z")
	bp := runalyze.BloodPressure{
		DateTime:  date,
		Systolic:  120,
		Diastolic: 80,
		HeartRate: 70,
	}
	resp, err := cl.CreateBloodPressure(ctx, bp)
	if err != nil {
		fmt.Println("Whoops!")
	}
	fmt.Println(resp.Status)
}
