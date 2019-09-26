# sleeper-api-go
Golang API client for the Sleeper Fantasy Football app

[![CircleCI](https://circleci.com/gh/jcobian/sleeper-api-go.svg?style=svg)](https://circleci.com/gh/jcobian/sleeper-api-go)

[Sleeper API Docs](https://docs.sleeper.app/)


## Usage
```go
import "github.com/jcobian/sleeper-api-go"
```

### TL;DR
```go
client := sleeper.NewAPIClient(nil)
resp, err := client.Stats.Get("nfl", "regular", "2019")
```
