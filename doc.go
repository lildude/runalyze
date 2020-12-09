/*
Package runalyze provides a client for using the Runalyze Personal API.

Usage:

	import "github.com/lildude/runalyze"

Construct a new Runalyze client, then call various methods on the API to access
different functions of the Runalyze API. All of the API calls require an access
token which should be passed when initialising the client. For example:

	ctx := context.Background()
	client := runalyze.NewClient(nil, "TOKEN")

	// Create a new sleep entry
	startSleep, _ := time.Parse(time.RFC3339, "2020-11-07T23:00:00Z")
	sleep := runalyze.Sleep{
		DateTime:           startSleep,
		Duration:           460,
	}
	resp, err := client.CreateSleep(ctx, sleep)

The Runalyze Personal API documentation is available at https://runalyze.com/doc/personal.

*/
package runalyze
