# peferb/trafiklab 
#### _github.com/peferb/trafiklab@v0.2.1_

Golang wrapper of [Trafiklab.se](https://trafiklab.se/) which supplies traffic information about Swedish trains, buses, trams and ships.

Implemented to match the official specifications as much as possible while trying to keep it simple and small and also
easy to use. 

**GO GET IT!**
```shell 
$ go get github.com/peferb/trafiklab@v0.2.1
```

## CLI example

```go
package main

import (
	"bytes"
	"encoding/json"
	"github.com/peferb/trafiklab/sl"
	. "github.com/peferb/trafiklab/sl/api"
	"log"
	"net/url"
	"os"
)


func main() {
	departuresApi := sl.NewApi(RealTimeDeparturesV4, JSON, "you api key here")
	departures, err := departuresApi.GetBytes(url.Values{"siteid": {os.Args[1]}})
	if err != nil {
		log.Println(err)
		return
	}

	prettyJson := &bytes.Buffer{}
	err = json.Indent(prettyJson, departures, "", "  ")
	log.Println(prettyJson.String(), err)
	// Example run:
	// $ go run . 9710
}
```


## Server example
```go
package main

import (
	"github.com/peferb/trafiklab/sl"
	. "github.com/peferb/trafiklab/sl/api"
	"log"
	"net/http"
)

func departuresHandler(w http.ResponseWriter, req *http.Request) {
	departuresApi := sl.NewApi(RealTimeDeparturesV4, JSON, "your api key here")
	
	bytes, err := departuresApi.GetBytes(req.URL.Query())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/realtime-departures", departuresHandler)
	log.Println("Listening for requests at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
	// Example url: 
	//localhost:8000/realtime-departures?siteid=9710
}
```

## Setup

* Sign up at [trafiklab.se](https://trafiklab.se)
* Generate keys
* Import wanted API from below

### SL Departures v4.0 (Realtidsinformation 4)
Package: github.com/peferb/trafiklab/sl/realtime_departures_v4
* [Generate API key](https://developer.trafiklab.se/api/sl-realtidsinformation-4)
* [API Specification](https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/)

### SL Service alerts v2.0 AKA Deviations (Störningsinformation 2)
Package: github.com/peferb/trafiklab/sl/service_alerts_v2
* [Generate API key](https://developer.trafiklab.se/api/sl-storningsinformation-2)
* [API Specification](https://www.trafiklab.se/api/trafiklab-apis/sl/service-alerts-2/)

### SL Stop lookup v1.0 AKA TypeAhead (Platsuppslag)
Package: github.com/peferb/trafiklab/sl/stop_lookup_v1
* [Generate API key](https://developer.trafiklab.se/api/sl-platsuppslag)
* [API Specification](https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/)

### SL Traffic status v2.0 AKA Traffic Situation (SL Trafikläget 2) 
Package: github.com/peferb/trafiklab/sl/traffic_status_v2
* [Generate API key](https://developer.trafiklab.se/api/sl-trafiklaget-2)
* [API Specification](https://www.trafiklab.se/api/trafiklab-apis/sl/sl-traffic-status-2/)

## TODO
_Nothing to see here ;)_
* Option to get data as structs 
* Wrap more API:s
  1. Rest of the SL API:s
  2. Remaining API:s
* Tests...?! 
