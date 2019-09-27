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
resp, err := client.Stats.GetAllStats("nfl", "regular", "2019")
// resp.Body, resp.StatusCode
```

### Responses
The different endpoints return arbitrary JSON that even varies within the different objects in the response. Hence, can't return a struct. Thus, each function returns either a map of strings to an aribtrary interface (usually another map) or an array of an arbitrary interface (usually a map).

For example:

```go
resp, err := client.Stats.GetAllStats("nfl", "regular", "2019")
statsForPlayer := resp.Body["5065"] // pick any key here in the response, these are player IDs. This is a map of this player's stats
```
