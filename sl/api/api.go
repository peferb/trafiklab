package api

import (
	"net/http"
	"net/url"
)

type Api struct {
	GetBytes    func(query url.Values) ([]byte, error)
	GetString   func(query url.Values) (string, error)
	GetResponse func(query url.Values) (*http.Response, error)
}

type DataFormat string

const JSON = "json"
const XML = "xml"

// Type of SL API
//
// SL has multiple different API:s to fetch data from. Read about them here:
// https://www.trafiklab.se/api/trafiklab-apis/sl/
type Type string

// RealTimeDeparturesV4 "Realtime Departures V4 (Realtidsinformation 4)"
//
// Access
// https://developer.trafiklab.se/api/sl-realtidsinformation-4
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/
const RealTimeDeparturesV4 Type = "https://api.sl.se/api2/realtimedeparturesV4"

// ServiceAlertsV2 "SL Stop lookup v1.0 AKA deviations (Platsuppslag)"
//
// Access
// https://developer.trafiklab.se/api/sl-platsuppslag
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/
const ServiceAlertsV2 Type = "https://api.sl.se/api2/deviations"

// StopLookupV1 "SL Stop lookup v1.0 AKA TypeAhead (Platsuppslag)"
//
// Access
// https://developer.trafiklab.se/api/sl-platsuppslag
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/
const StopLookupV1 Type = "https://api.sl.se/api2/typeahead"

// TrafficStatusV2 "SL Traffic status v2.0 AKA Traffic Situation (SL Trafikläget 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-trafiklaget-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/sl-traffic-status-2/
const TrafficStatusV2 Type = "https://api.sl.se/api2/trafficsituation"

// StopsAndLinesV2 "SL Stops and lines v2.0 AKA LineData (SL Hållplatser och Linjer 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-hallplatser-och-linjer-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stops-and-lines-2/
const StopsAndLinesV2 Type = "https://api.sl.se/api2/LineData"
