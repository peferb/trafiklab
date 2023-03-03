# peferb/trafiklab 
#### _github.com/peferb/trafiklab@v0.0.8_

Trafiklab.se Golang SDK for Swedish trains etc.

_Only json, no xml._

## Setup

* Sign up at [trafiklab.se](https://trafiklab.se)
* Generate keys
    * [SL_REAL_TIME_DEPARTURES_V4](https://developer.trafiklab.se/api/sl-realtidsinformation-4)
* Init your golang project and:
```shell 
$ go get github.com/peferb/trafiklab@v0.0.8
```

## Usage

### SL Departures v4.0 (Realtidsinformation 4)
[API Specification and params](https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/)
```go
import (
    trafik "github.com/peferb/trafiklab"
    "log"
    "net/http"
    "net/url"
)

// Handler that just wraps the API ([spec](https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/))
func Handler(w http.ResponseWriter, req *http.Request) {
    trafik.SetSLRealtimeDeparturesV4Key("SL_REAL_TIME_DEPARTURES_V4")
    body, err := trafik.GetSLRealtimeDeparturesV4(req.URL.Query())
    if err != nil {
        w.WriteHeader(500)
        w.Write([]byte(err.Error()))
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte(body))
}

// Or just print the data...
func printBody() {
    trafik.SetSLRealtimeDeparturesV4Key("SL_REAL_TIME_DEPARTURES_V4")
    q := url.Values{}
    q.Set("siteid", "9710")
    q.Set("train", "true")
    q.Set("metro", "false")
    q.Set("tram", "false")
    data, err := GetSLRealtimeDeparturesV4(q)
    log.Println(body, err)
}
```