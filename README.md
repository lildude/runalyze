# Runalyze

An unofficial Go client for the [Runalyze Personal API](https://runalyze.com/help/article/personal-api).

## Installation

Use Go to fetch the latest version of the package.

```shell
go get -u 'github.com/lildude/runalyze'
```

## Usage

You will need an access token to query the API. You can generate tokens in your account at [Personal API](https://runalyze.com/settings/personal-api).

With your token, you can then access the API using approach similar to this:

```go
package main

import (
  "context"
  "fmt"
  "os"
  "time"

  "github.com/joho/godotenv"
  "github.com/lildude/runalyze"
)

func main() {
  godotenv.Load(".env")
  ctx := context.Background()
  cl := runalyze.NewClient(nil, os.Getenv("RUNALYZE_ACCESS_TOKEN"))

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
    fmt.Println("Whoops!")
  }
  fmt.Println(resp.Status)
}
```
