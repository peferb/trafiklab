# peferb/trafiklab 
#### _github.com/peferb/trafiklab@v0.2.4 ([changelog](https://github.com/peferb/trafiklab/blob/main/CHANGELOG.md))_

Golang wrapper of [Trafiklab.se](https://trafiklab.se/) which supplies traffic information about Swedish trains, buses, trams and ships.

Implemented to match the official specifications as much as possible while trying to keep it simple and small and also
easy to use. 

**GO GET IT!**
```shell 
$ go get github.com/peferb/trafiklab@v0.2.4
```

![i-like-trains](https://github.com/user-attachments/assets/7c227f60-8995-4b58-a6b3-d7e2bfbc3d86)


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
	// localhost:8000/realtime-departures?siteid=9710
}
```

## Setup

* Sign up at [trafiklab.se](https://trafiklab.se)
* Generate keys

### SL Departures v4.0 (Realtidsinformation 4)
* [Generate API key](https://developer.trafiklab.se/api/sl-realtidsinformation-4)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/)

### SL Service alerts v2.0 AKA Deviations (Störningsinformation 2)
* [Generate API key](https://developer.trafiklab.se/api/sl-storningsinformation-2)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/service-alerts-2/)

### SL Stop lookup v1.0 AKA TypeAhead (Platsuppslag)
* [Generate API key](https://developer.trafiklab.se/api/sl-platsuppslag)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/)

### SL Traffic status v2.0 AKA Traffic Situation (SL Trafikläget 2) 
* [Generate API key](https://developer.trafiklab.se/api/sl-trafiklaget-2)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/sl-traffic-status-2/)

### SL Stops and lines v2.0 AKA LineData (SL Hållplatser och Linjer 2)
* [Generate API key](https://developer.trafiklab.se/api/sl-hallplatser-och-linjer-2)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/stops-and-lines-2/)

### SL Nearby stops v2.0 (Närliggande hållplatser 2)
* [Generate API key](https://developer.trafiklab.se/api/sl-narliggande-hallplatser-2)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/nearby-stops-2/)

### SL Route Planner V3.1 AKA Travel Planner V3.1 (SL Reseplanerare 3.1)
⚠️ _"Route Planner V3.1" differs a bit from the other API:s by having multiple endpoints, and two of the endpoints only 
return data in XML format._
* [Generate API key](https://developer.trafiklab.se/api/sl-reseplanerare-31)
* [API Documentation](https://www.trafiklab.se/api/trafiklab-apis/sl/route-planner-31/)
