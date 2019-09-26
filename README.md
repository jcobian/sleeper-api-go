# sleeper-api-go
Golang API client for the Sleeper Fantasy Football app

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
