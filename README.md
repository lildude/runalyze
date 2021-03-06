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
  cl := runalyze.NewClient(os.Getenv("RUNALYZE_ACCESS_TOKEN"))

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

## Releasing

This project uses [GoReleaser](https://goreleaser.com) via GitHub Actions to make the releases quick and easy. When I'm ready for a new release, I push a new tag and the workflow takes care of things.

## Todo

- Add support for uploading activities
