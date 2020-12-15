# Runalyze

![Tests Status Badge](https://github.com/lildude/runalyze/workflows/Tests/badge.svg)

An unofficial Go client for the [Runalyze Personal API](https://runalyze.com/help/article/personal-api).

## Installation

Use Go to fetch the latest version of the package.

```shell
go get -u 'github.com/lildude/runalyze'
```

## Usage

You will need an access token to query the API. You can generate tokens in your account at [Personal API](https://runalyze.com/settings/personal-api).

When using the client, you will also need to supply a name of your application. This is used in the user agent when querying the API and is used to identify applications that are accessing the API and enable Runalyze to contact the application author if there are problems. So pick a name that stands out!

With your token and application name, you can then access the API using approach similar to this:

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
  cfg := runalyze.Configuration{
    AppName: "My Cool App/3.2.1",
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
    fmt.Println(err)
  }
  fmt.Println(resp.Status)
}
```

## Motivation

I've written this myself from scratch, heavily influenced by the methods implemented in <https://github.com/billglover/starling>, as a project to help practice Go, but also because I found the Swagger Codegen'd code quite cumbersome, ugly and without any tests. The generated code is also not Go modules friendly and requires bundling with your application. You can see what I've managed to generate in the repo at <https://github.com/lildude/runalyze-generated>.

## Todo

- Add support for uploading activities
