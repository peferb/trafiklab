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

// JSON Dataformat
const JSON = "json"

// XML Dataformat
const XML = "xml"

// RoutePlannerV3_1Trip Trip part of "Route Planner V3.1 AKA Travel Planner V3.1 (SL Reseplanerare 3.1)"
//
// Access
// https://developer.trafiklab.se/api/sl-reseplanerare-31
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/route-planner-31/
const RoutePlannerV3_1Trip = "api.sl.se/api2/TravelplannerV3_1/trip"

// RoutePlannerV3_1GisRoute Gis Route part of "Route Planner V3.1 AKA Travel Planner V3.1 (SL Reseplanerare 3.1)"
//
// Access
// https://developer.trafiklab.se/api/sl-reseplanerare-31
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/route-planner-31/
const RoutePlannerV3_1GisRoute = "api.sl.se/api2/TravelplannerV3_1/gisroute"

// http://api.sl.se/api2/TravelplannerV3_1/gisroute.&lt;FORMAT&gt;?key=&lt;DIN APINYCKEL&gt;&amp;ctx=

// RoutePlannerV3_1JourneyDetail Journey Detail part of"Route Planner V3.1 AKA Travel Planner V3.1 (SL Reseplanerare 3.1)"
//
// Access
// https://developer.trafiklab.se/api/sl-reseplanerare-31
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/route-planner-31/
const RoutePlannerV3_1JourneyDetail = "api.sl.se/api2/TravelplannerV3_1/journeydetail"

// RealTimeDeparturesV4 "Realtime Departures V4 (Realtidsinformation 4)"
//
// Access
// https://developer.trafiklab.se/api/sl-realtidsinformation-4
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/
const RealTimeDeparturesV4 = "api.sl.se/api2/realtimedeparturesV4"

// ServiceAlertsV2 "SL Stop lookup v1.0 AKA deviations (Platsuppslag)"
//
// Access
// https://developer.trafiklab.se/api/sl-platsuppslag
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/
const ServiceAlertsV2 = "api.sl.se/api2/deviations"

// StopLookupV1 "SL Stop lookup v1.0 AKA TypeAhead (Platsuppslag)"
//
// Access
// https://developer.trafiklab.se/api/sl-platsuppslag
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/
const StopLookupV1 = "api.sl.se/api2/typeahead"

// TrafficStatusV2 "SL Traffic status v2.0 AKA Traffic Situation (SL Trafikläget 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-trafiklaget-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/sl-traffic-status-2/
const TrafficStatusV2 = "api.sl.se/api2/trafficsituation"

// StopsAndLinesV2 "SL Stops and lines v2.0 AKA LineData (SL Hållplatser och Linjer 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-hallplatser-och-linjer-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stops-and-lines-2/
const StopsAndLinesV2 = "api.sl.se/api2/LineData"

// NearbyStopsV2 "SL Nearby stops v2.0 (Närliggande hållplatser 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-narliggande-hallplatser-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/nearby-stops-2/
const NearbyStopsV2 = "api.sl.se/api2/nearbystopsv2"
