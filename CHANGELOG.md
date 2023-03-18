# peferb/trafiklab changelog

## v.0.2.3
### Changed all parameters to stlib types
Now DataFormat type and "Api" Type is replaced by regular strings.

I came to his decision after reading [Jack Lindamoods article "Aspects of a good Go library"](https://medium.com/@cep21/aspects-of-a-good-go-library-7082beabb403).

### NewApi returns Api...
... and not Struct{Api} as before. Don't even know what Struct{Api} is. The previous return type was the result of a 
late late night in a new lang 😅

## v0.2.1 / v0.2.2
_(duplicated version because of me messing things up)_

### Breaking changes
* Replaced the package import usage with `sl.NewApi()`.
* Moved data structure selection from func name to param. 

#### Example new usage 
```go
import (
	"github.com/peferb/trafiklab/sl"
	. "github.com/peferb/trafiklab/sl/api"
)

func main() {
	departuresApi := sl.NewApi(RealTimeDeparturesV4, JSON, "you api key here")
}
```

### New functionality
* Added "SL Nearby stops v2.0 (Närliggande hållplatser 2)"
* Added "SL Stops and lines v2.0 AKA LineData (SL Hållplatser och Linjer 2)"

### Misc
* Much smaller codebase / lib size
* Less bloated dev interface
* Added tests

## v0.1.0

#### Includes packages for:
* SL Departures v4.0 (Realtidsinformation 4)
* SL Service alerts v2.0 AKA Deviations (Störningsinformation 2)
* SL Stop lookup v1.0 AKA TypeAhead (Platsuppslag)
* SL Traffic status v2.0 AKA Traffic Situation (SL Trafikläget 2)

#### Example usage
```go
import (
	realtime "github.com/peferb/trafiklab/sl/realtime_departures_v4"
)

func main() {
	bytes, err := realtime.GetBytesJson("your api key here", req.URL.Query())
}
```