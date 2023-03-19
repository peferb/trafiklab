package sl

import (
	. "github.com/peferb/trafiklab/sl/api"
	"io"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	statusSuccessJson = `"StatusCode":0`
	statusSuccessXml  = "<StatusCode>0</StatusCode>"
)

func getEnvVar(envKey string, t *testing.T) string {
	apiKey := os.Getenv(envKey)
	if len(apiKey) == 0 {
		t.Fatalf("Missing API key '%s' in os env", envKey)
	}
	return apiKey
}

func testStatusCode(testName string, t *testing.T, envVar string, api string, format string, apiParams url.Values, resType string, shouldContain string) {
	apiKey := getEnvVar(envVar, t)
	departures := NewApi(api, format, apiKey)
	if resType == "string" {
		data, err := departures.GetString(apiParams)
		if err != nil || !strings.Contains(data, shouldContain) {
			t.Fatalf(
				`%s %v key = "%s" exptected to contain "%s", got: data = %s, error = %s`,
				testName,
				format,
				apiKey,
				shouldContain,
				data,
				err,
			)
		}
	} else if resType == "bytes" {
		data, err := departures.GetBytes(apiParams)
		if err != nil || !strings.Contains(string(data), shouldContain) {
			t.Fatalf(
				`%s %v key = "%s" exptected to contain "%s", got: data = %s, error = %s`,
				testName,
				format,
				apiKey,
				shouldContain,
				data,
				err,
			)
		}
	} else if resType == "response" {
		res, err := departures.GetResponse(apiParams)
		if res.StatusCode != 200 || err != nil {
			t.Fatalf("%s not 200, response='%v', error='%s'", testName, res, err.Error())
		}

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("%s could not read body '%s'", testName, err.Error())
		}
		defer res.Body.Close()
		body := string(bodyBytes)
		if err != nil || !strings.Contains(body, shouldContain) {
			t.Fatalf(
				`%s %v key = "%s" exptected to contain "%s", got: body = %s, error = %s`,
				testName,
				format,
				apiKey,
				shouldContain,
				body,
				err,
			)
		}
	} else {
		t.Fatalf(`%s Unkown res type "%s"`, testName, resType)
	}
}

/*
 * Route-planner v3.1 Gis Route
 */

// TestRoutePlannerV3_1GisRouteStringJson tries to fetch Gis Route data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1GisRouteStringJson(t *testing.T) {
	apiParams := url.Values{"ploy": {"1"}, "lang": {"sv"}, "ctx": {"G|1|G@F|A=1@O=Stockholm%20City@X=18059500@Y=59331143@U=74@L=400105313@|A=2@O=Stockholm,%20Drottninggatan%2022@l=@X=18065001@Y=59330415@u=0@|12022019|84500|85100|ft|ft@0@1000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7bt@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7tt@0@5000@120@-1@100@1@2500@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7%7C}"}}
	testStatusCode("TestRoutePlannerV3_1GisRouteStringJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1GisRoute, JSON, apiParams, "string", `"GisRoute":`)
}

// TestRoutePlannerV3_1GisRouteStringXml tries to fetch Gis Route data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1GisRouteStringXml(t *testing.T) {
	apiParams := url.Values{"ploy": {"1"}, "lang": {"sv"}, "ctx": {"G|1|G@F|A=1@O=Stockholm%20City@X=18059500@Y=59331143@U=74@L=400105313@|A=2@O=Stockholm,%20Drottninggatan%2022@l=@X=18065001@Y=59330415@u=0@|12022019|84500|85100|ft|ft@0@1000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7bt@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7tt@0@5000@120@-1@100@1@2500@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7%7C}"}}
	testStatusCode("TestRoutePlannerV3_1GisRouteStringXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1GisRoute, XML, apiParams, "string", `<GisRoute`)
}

// TestRoutePlannerV3_1GisRouteBytesJson tries to fetch Gis Route data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1GisRouteBytesJson(t *testing.T) {
	apiParams := url.Values{"ploy": {"1"}, "lang": {"sv"}, "ctx": {"G|1|G@F|A=1@O=Stockholm%20City@X=18059500@Y=59331143@U=74@L=400105313@|A=2@O=Stockholm,%20Drottninggatan%2022@l=@X=18065001@Y=59330415@u=0@|12022019|84500|85100|ft|ft@0@1000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7bt@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7tt@0@5000@120@-1@100@1@2500@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7%7C}"}}
	testStatusCode("TestRoutePlannerV3_1GisRouteBytesJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1GisRoute, JSON, apiParams, "bytes", `"GisRoute":`)
}

// TestRoutePlannerV3_1GisRouteBytesXml tries to fetch Gis Route data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1GisRouteBytesXml(t *testing.T) {
	apiParams := url.Values{"ploy": {"1"}, "lang": {"sv"}, "ctx": {"G|1|G@F|A=1@O=Stockholm%20City@X=18059500@Y=59331143@U=74@L=400105313@|A=2@O=Stockholm,%20Drottninggatan%2022@l=@X=18065001@Y=59330415@u=0@|12022019|84500|85100|ft|ft@0@1000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7bt@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7tt@0@5000@120@-1@100@1@2500@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7%7C}"}}
	testStatusCode("TestRoutePlannerV3_1GisRouteBytesXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1GisRoute, XML, apiParams, "bytes", `<GisRoute`)
}

// TestRoutePlannerV3_1GisRouteResponseJson tries to fetch Gis Route data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1GisRouteResponseJson(t *testing.T) {
	apiParams := url.Values{"ploy": {"1"}, "lang": {"sv"}, "ctx": {"G|1|G@F|A=1@O=Stockholm%20City@X=18059500@Y=59331143@U=74@L=400105313@|A=2@O=Stockholm,%20Drottninggatan%2022@l=@X=18065001@Y=59330415@u=0@|12022019|84500|85100|ft|ft@0@1000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7bt@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7tt@0@5000@120@-1@100@1@2500@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$%C2%A7%7C}"}}
	testStatusCode("TestRoutePlannerV3_1GisRouteResponseJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1GisRoute, JSON, apiParams, "response", `"GisRoute":`)
}

/*
 * Route-planner v3.1 XSD
 */

// TestRoutePlannerV3_1XsdStringJson tries to fetch XSD data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1XsdStringJson(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TestRoutePlannerV3_1XsdStringJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Xsd, JSON, apiParams, "string", `<xs:annotation><xs:documentation>`)
}

// TestRoutePlannerV3_1XsdStringXml tries to fetch XSD data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1XsdStringXml(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TestRoutePlannerV3_1XsdStringXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Xsd, XML, apiParams, "string", `<xs:annotation><xs:documentation>`)
}

// TestRoutePlannerV3_1XsdBytesJson tries to fetch XSD data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1XsdBytesJson(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TestRoutePlannerV3_1XsdBytesJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Xsd, JSON, apiParams, "bytes", `<xs:annotation><xs:documentation>`)
}

// TestRoutePlannerV3_1XsdBytesXml tries to fetch XSD data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1XsdBytesXml(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TestRoutePlannerV3_1XsdBytesXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Xsd, XML, apiParams, "bytes", `<xs:annotation><xs:documentation>`)
}

// TestRoutePlannerV3_1XsdResponseJson tries to fetch XSD data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1XsdResponseJson(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TestRoutePlannerV3_1XsdResponseJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Xsd, JSON, apiParams, "response", `<xs:annotation><xs:documentation>`)
}

// TestRoutePlannerV3_1XsdResponseXml tries to fetch XSD data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1XsdResponseXml(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TestRoutePlannerV3_1XsdResponseXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Xsd, XML, apiParams, "response", `<xs:annotation><xs:documentation>`)
}

/*
 * Route-planner v3.1 Reconstruction
 */

// TestRoutePlannerV3_1ReconstructionStringJson tries to fetch Reconstruction data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1ReconstructionStringJson(t *testing.T) {
	apiParams := url.Values{"ctx": {"T%24A%3D1%40O%3DT-Centralen%40L%3D400101052%40a%3D128%40%24A%3D1%40O%3DSlussen%40L%3D400101012%40a%3D128%40%24201706140859%24201706140902%24%20%24"}}
	testStatusCode("TestRoutePlannerV3_1ReconstructionStringJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Reconstruction, JSON, apiParams, "string", `"Trip":[`)
}

// TestRoutePlannerV3_1ReconstructionStringXml tries to fetch Reconstruction data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1ReconstructionStringXml(t *testing.T) {
	apiParams := url.Values{"ctx": {"T%24A%3D1%40O%3DT-Centralen%40L%3D400101052%40a%3D128%40%24A%3D1%40O%3DSlussen%40L%3D400101012%40a%3D128%40%24201706140859%24201706140902%24%20%24"}}
	testStatusCode("TestRoutePlannerV3_1ReconstructionStringXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Reconstruction, XML, apiParams, "string", `<TripList`)
}

// TestRoutePlannerV3_1ReconstructionBytesJson tries to fetch Reconstruction data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1ReconstructionBytesJson(t *testing.T) {
	apiParams := url.Values{"ctx": {"T%24A%3D1%40O%3DT-Centralen%40L%3D400101052%40a%3D128%40%24A%3D1%40O%3DSlussen%40L%3D400101012%40a%3D128%40%24201706140859%24201706140902%24%20%24"}}
	testStatusCode("TestRoutePlannerV3_1ReconstructionBytesJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Reconstruction, JSON, apiParams, "bytes", `"Trip":[`)
}

// TestRoutePlannerV3_1ReconstructionBytesXml tries to fetch Reconstruction data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1ReconstructionBytesXml(t *testing.T) {
	apiParams := url.Values{"ctx": {"T%24A%3D1%40O%3DT-Centralen%40L%3D400101052%40a%3D128%40%24A%3D1%40O%3DSlussen%40L%3D400101012%40a%3D128%40%24201706140859%24201706140902%24%20%24"}}
	testStatusCode("TestRoutePlannerV3_1ReconstructionBytesXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Reconstruction, XML, apiParams, "bytes", `<TripList`)
}

// TestRoutePlannerV3_1ReconstructionResponseJson tries to fetch Reconstruction data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1ReconstructionResponseJson(t *testing.T) {
	apiParams := url.Values{"ctx": {"T%24A%3D1%40O%3DT-Centralen%40L%3D400101052%40a%3D128%40%24A%3D1%40O%3DSlussen%40L%3D400101012%40a%3D128%40%24201706140859%24201706140902%24%20%24"}}
	testStatusCode("TestRoutePlannerV3_1ReconstructionResponseJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Reconstruction, JSON, apiParams, "response", `"Trip":[`)
}

// TestRoutePlannerV3_1ReconstructionResponseXml tries to fetch Reconstruction data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1ReconstructionResponseXml(t *testing.T) {
	apiParams := url.Values{"ctx": {"T%24A%3D1%40O%3DT-Centralen%40L%3D400101052%40a%3D128%40%24A%3D1%40O%3DSlussen%40L%3D400101012%40a%3D128%40%24201706140859%24201706140902%24%20%24"}}
	testStatusCode("TestRoutePlannerV3_1ReconstructionResponseXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Reconstruction, XML, apiParams, "response", `<TripList`)
}

/*
 * Route-planner v3.1 Trip
 */

// TestRoutePlannerV3_1TripStringJson tries to fetch Trip data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1TripStringJson(t *testing.T) {
	apiParams := url.Values{"originExtId": {"9710"}, "destExtId": {"9001"}, "lang": {"en"}}
	testStatusCode("TestRoutePlannerV3_1TripStringJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Trip, JSON, apiParams, "string", `"Trip":`)
}

// TestRoutePlannerV3_1TripStringXml tries to fetch Trip data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1TripStringXml(t *testing.T) {
	apiParams := url.Values{"originExtId": {"9710"}, "destExtId": {"9001"}, "lang": {"en"}}
	testStatusCode("TestRoutePlannerV3_1TripStringXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Trip, XML, apiParams, "string", `<TripList`)
}

// TestRoutePlannerV3_1TripBytesJson tries to fetch Trip data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1TripBytesJson(t *testing.T) {
	apiParams := url.Values{"originExtId": {"9710"}, "destExtId": {"9001"}, "lang": {"en"}}
	testStatusCode("TestRoutePlannerV3_1TripBytesJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Trip, JSON, apiParams, "bytes", `"Trip":`)
}

// TestRoutePlannerV3_1TripBytesXml tries to fetch Trip data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1TripBytesXml(t *testing.T) {
	apiParams := url.Values{"originExtId": {"9710"}, "destExtId": {"9001"}, "lang": {"en"}}
	testStatusCode("TestRoutePlannerV3_1TripBytesXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Trip, XML, apiParams, "bytes", `<TripList`)
}

// TestRoutePlannerV3_1TripResponseJson tries to fetch Trip data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1TripResponseJson(t *testing.T) {
	apiParams := url.Values{"originExtId": {"9710"}, "destExtId": {"9001"}, "lang": {"en"}}
	testStatusCode("TestRoutePlannerV3_1TripResponseJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Trip, JSON, apiParams, "response", `"Trip":`)
}

// TestRoutePlannerV3_1TripResponseXml tries to fetch Trip data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1TripResponseXml(t *testing.T) {
	apiParams := url.Values{"originExtId": {"9710"}, "destExtId": {"9001"}, "lang": {"en"}}
	testStatusCode("TestRoutePlannerV3_1TripResponseXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1Trip, XML, apiParams, "response", `<TripList`)
}

/*
 * Route-planner v3.1 Journey Detail
 */

// TestRoutePlannerV3_1JourneyDetailStringJson tries to fetch Journey Detail data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1JourneyDetailStringJson(t *testing.T) {
	apiParams := url.Values{"id": {"1|14536|1|74|18032023"}}
	testStatusCode("TestRoutePlannerV3_1JourneyDetailStringJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1JourneyDetail, JSON, apiParams, "string", `"Stops":`)
}

// TestRoutePlannerV3_1JourneyDetailStringXml tries to fetch Journey Detail data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1JourneyDetailStringXml(t *testing.T) {
	apiParams := url.Values{"id": {"1|14536|1|74|18032023"}}
	testStatusCode("TestRoutePlannerV3_1JourneyDetailStringXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1JourneyDetail, XML, apiParams, "string", `<Stops>`)
}

// TestRoutePlannerV3_1JourneyDetailBytesJson tries to fetch Journey Detail data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1JourneyDetailBytesJson(t *testing.T) {
	apiParams := url.Values{"id": {"1|14536|1|74|18032023"}}
	testStatusCode("TestRoutePlannerV3_1JourneyDetailBytesJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1JourneyDetail, JSON, apiParams, "bytes", `"Stops":`)
}

// TestRoutePlannerV3_1JourneyDetailBytesXml tries to fetch Journey Detail data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1JourneyDetailBytesXml(t *testing.T) {
	apiParams := url.Values{"id": {"1|14536|1|74|18032023"}}
	testStatusCode("TestRoutePlannerV3_1JourneyDetailBytesXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1JourneyDetail, XML, apiParams, "bytes", `<Stops>`)
}

// TestRoutePlannerV3_1JourneyDetailResponseJson tries to fetch Journey Detail data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1JourneyDetailResponseJson(t *testing.T) {
	apiParams := url.Values{"id": {"1|14536|1|74|18032023"}}
	testStatusCode("TestRoutePlannerV3_1JourneyDetailResponseJson", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1JourneyDetail, JSON, apiParams, "response", `"Stops":`)
}

// TestRoutePlannerV3_1JourneyDetailResponseXml tries to fetch Journey Detail data from RoutePlannerV3_1 API
func TestRoutePlannerV3_1JourneyDetailResponseXml(t *testing.T) {
	apiParams := url.Values{"id": {"1|14536|1|74|18032023"}}
	testStatusCode("TestRoutePlannerV3_1JourneyDetailResponseXml", t, "SL_ROUTE_PLANNER_V3_1", RoutePlannerV3_1JourneyDetail, XML, apiParams, "response", `<Stops>`)
}

/*
 * RealTimeDeparturesV4
 */

// TestRealTimeDeparturesV4StringJson tries to fetch data from RealTimeDeparturesV4 API
func TestRealTimeDeparturesV4StringJson(t *testing.T) {
	apiParams := url.Values{"siteid": {"9710"}}
	testStatusCode("TestRealTimeDeparturesV4StringJson", t, "SL_REAL_TIME_DEPARTURES_V4", RealTimeDeparturesV4, JSON, apiParams, "string", statusSuccessJson)
}

// TestRealTimeDeparturesV4StringXml tries to fetch data from RealTimeDeparturesV4 API
func TestRealTimeDeparturesV4StringXml(t *testing.T) {
	apiParams := url.Values{"siteid": {"9710"}}
	testStatusCode("TestRealTimeDeparturesV4StringXml", t, "SL_REAL_TIME_DEPARTURES_V4", RealTimeDeparturesV4, XML, apiParams, "string", statusSuccessXml)
}

// TestRealTimeDeparturesV4BytesJson tries to fetch data from RealTimeDeparturesV4 API
func TestRealTimeDeparturesV4BytesJson(t *testing.T) {
	apiParams := url.Values{"siteid": {"9710"}}
	testStatusCode("TestRealTimeDeparturesV4BytesJson", t, "SL_REAL_TIME_DEPARTURES_V4", RealTimeDeparturesV4, JSON, apiParams, "bytes", statusSuccessJson)
}

// TestRealTimeDeparturesV4BytesXml tries to fetch data from RealTimeDeparturesV4 API
func TestRealTimeDeparturesV4BytesXml(t *testing.T) {
	apiParams := url.Values{"siteid": {"9710"}}
	testStatusCode("TestRealTimeDeparturesV4BytesXml", t, "SL_REAL_TIME_DEPARTURES_V4", RealTimeDeparturesV4, XML, apiParams, "bytes", statusSuccessXml)
}

// TestRealTimeDeparturesV4ResponseJson tries to fetch data from RealTimeDeparturesV4 API
func TestRealTimeDeparturesV4ResponseJson(t *testing.T) {
	apiParams := url.Values{"siteid": {"9710"}}
	testStatusCode("TestRealTimeDeparturesV4ResponseJson", t, "SL_REAL_TIME_DEPARTURES_V4", RealTimeDeparturesV4, JSON, apiParams, "response", statusSuccessJson)
}

// TestRealTimeDeparturesV4ResponseXml tries to fetch data from RealTimeDeparturesV4 API
func TestRealTimeDeparturesV4ResponseXml(t *testing.T) {
	apiParams := url.Values{"siteid": {"9710"}}
	testStatusCode("TestRealTimeDeparturesV4ResponseXml", t, "SL_REAL_TIME_DEPARTURES_V4", RealTimeDeparturesV4, XML, apiParams, "response", statusSuccessXml)
}

/*
* ServiceAlertsV2
 */

// TestServiceAlertsV2StringJson tries to fetch data from ServiceAlertsV2 API
func TestServiceAlertsV2StringJson(t *testing.T) {
	apiParams := url.Values{"siteid": {"9000"}}
	testStatusCode("TestServiceAlertsV2StringJson", t, "SL_SERVICE_ALERTS_V2", ServiceAlertsV2, JSON, apiParams, "string", statusSuccessJson)
}

// TestServiceAlertsV2StringXml tries to fetch data from ServiceAlertsV2 API
func TestServiceAlertsV2StringXml(t *testing.T) {
	apiParams := url.Values{"siteid": {"9000"}}
	testStatusCode("TestServiceAlertsV2StringXml", t, "SL_SERVICE_ALERTS_V2", ServiceAlertsV2, XML, apiParams, "string", statusSuccessXml)
}

// TestServiceAlertsV2BytesJson tries to fetch data from ServiceAlertsV2 API
func TestServiceAlertsV2BytesJson(t *testing.T) {
	apiParams := url.Values{"siteid": {"9000"}}
	testStatusCode("TestServiceAlertsV2BytesJson", t, "SL_SERVICE_ALERTS_V2", ServiceAlertsV2, JSON, apiParams, "bytes", statusSuccessJson)
}

// TestServiceAlertsV2BytesXml tries to fetch data from ServiceAlertsV2 API
func TestServiceAlertsV2BytesXml(t *testing.T) {
	apiParams := url.Values{"siteid": {"9000"}}
	testStatusCode("TestServiceAlertsV2BytesXml", t, "SL_SERVICE_ALERTS_V2", ServiceAlertsV2, XML, apiParams, "bytes", statusSuccessXml)
}

// TestServiceAlertsV2ResponseJson tries to fetch data from ServiceAlertsV2 API
func TestServiceAlertsV2ResponseJson(t *testing.T) {
	apiParams := url.Values{"siteid": {"9000"}}
	testStatusCode("TestServiceAlertsV2ResponseJson", t, "SL_SERVICE_ALERTS_V2", ServiceAlertsV2, JSON, apiParams, "response", statusSuccessJson)
}

// TestServiceAlertsV2ResponseXml tries to fetch data from ServiceAlertsV2 API
func TestServiceAlertsV2ResponseXml(t *testing.T) {
	apiParams := url.Values{"siteid": {"9000"}}
	testStatusCode("TestServiceAlertsV2ResponseXml", t, "SL_SERVICE_ALERTS_V2", ServiceAlertsV2, XML, apiParams, "response", statusSuccessXml)
}

/*
* StopLookupV1
 */

// TestStopLookupV1StringJson tries to fetch data from StopLookupV1 API
func TestStopLookupV1StringJson(t *testing.T) {
	apiParams := url.Values{"searchstring": {"stockholm"}}
	testStatusCode("StopLookupV1StringJson", t, "SL_STOP_LOOKUP_V1", StopLookupV1, JSON, apiParams, "string", statusSuccessJson)
}

// TestStopLookupV1StringXml tries to fetch data from StopLookupV1 API
func TestStopLookupV1StringXml(t *testing.T) {
	apiParams := url.Values{"searchstring": {"stockholm"}}
	testStatusCode("StopLookupV1StringXml", t, "SL_STOP_LOOKUP_V1", StopLookupV1, XML, apiParams, "string", statusSuccessXml)
}

// TestStopLookupV1BytesJson tries to fetch data from StopLookupV1 API
func TestStopLookupV1BytesJson(t *testing.T) {
	apiParams := url.Values{"searchstring": {"stockholm"}}
	testStatusCode("StopLookupV1BytesJson", t, "SL_STOP_LOOKUP_V1", StopLookupV1, JSON, apiParams, "bytes", statusSuccessJson)
}

// TestStopLookupV1BytesXml tries to fetch data from StopLookupV1 API
func TestStopLookupV1BytesXml(t *testing.T) {
	apiParams := url.Values{"searchstring": {"stockholm"}}
	testStatusCode("StopLookupV1BytesXml", t, "SL_STOP_LOOKUP_V1", StopLookupV1, XML, apiParams, "bytes", statusSuccessXml)
}

// TestStopLookupV1ResponseJson tries to fetch data from StopLookupV1 API
func TestStopLookupV1ResponseJson(t *testing.T) {
	apiParams := url.Values{"searchstring": {"stockholm"}}
	testStatusCode("StopLookupV1ResponseJson", t, "SL_STOP_LOOKUP_V1", StopLookupV1, JSON, apiParams, "response", statusSuccessJson)
}

// TestStopLookupV1ResponseXml tries to fetch data from StopLookupV1 API
func TestStopLookupV1ResponseXml(t *testing.T) {
	apiParams := url.Values{"searchstring": {"stockholm"}}
	testStatusCode("StopLookupV1ResponseXml", t, "SL_STOP_LOOKUP_V1", StopLookupV1, XML, apiParams, "response", statusSuccessXml)
}

/*
* TrafficStatusV2
 */

// TestTrafficStatusV2StringJson tries to fetch data from TrafficStatusV2 API
func TestTrafficStatusV2StringJson(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TrafficStatusV2StringJson", t, "SL_TRAFFIC_STATUS_V2", TrafficStatusV2, JSON, apiParams, "string", statusSuccessJson)
}

// TestTrafficStatusV2StringXml tries to fetch data from TrafficStatusV2 API
func TestTrafficStatusV2StringXml(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TrafficStatusV2StringXml", t, "SL_TRAFFIC_STATUS_V2", TrafficStatusV2, XML, apiParams, "string", statusSuccessXml)
}

// TestTrafficStatusV2BytesJson tries to fetch data from TrafficStatusV2 API
func TestTrafficStatusV2BytesJson(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TrafficStatusV2BytesJson", t, "SL_TRAFFIC_STATUS_V2", TrafficStatusV2, JSON, apiParams, "bytes", statusSuccessJson)
}

// TestTrafficStatusV2BytesXml tries to fetch data from TrafficStatusV2 API
func TestTrafficStatusV2BytesXml(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TrafficStatusV2BytesXml", t, "SL_TRAFFIC_STATUS_V2", TrafficStatusV2, XML, apiParams, "bytes", statusSuccessXml)
}

// TestTrafficStatusV2ResponseJson tries to fetch data from TrafficStatusV2 API
func TestTrafficStatusV2ResponseJson(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TrafficStatusV2ResponseJson", t, "SL_TRAFFIC_STATUS_V2", TrafficStatusV2, JSON, apiParams, "response", statusSuccessJson)
}

// TestTrafficStatusV2ResponseXml tries to fetch data from TrafficStatusV2 API
func TestTrafficStatusV2ResponseXml(t *testing.T) {
	apiParams := url.Values{}
	testStatusCode("TrafficStatusV2ResponseXml", t, "SL_TRAFFIC_STATUS_V2", TrafficStatusV2, XML, apiParams, "response", statusSuccessXml)
}

/*
* StopsAndLinesV2
 */

// TestStopsAndLinesV2StringJson tries to fetch data from StopsAndLinesV2 API
func TestStopsAndLinesV2StringJson(t *testing.T) {
	apiParams := url.Values{"model": {"Transportmode"}}
	testStatusCode("StopsAndLinesV2StringJson", t, "SL_STOPS_AND_LINES_V2", StopsAndLinesV2, JSON, apiParams, "string", statusSuccessJson)
}

// TestStopsAndLinesV2StringXml tries to fetch data from StopsAndLinesV2 API
func TestStopsAndLinesV2StringXml(t *testing.T) {
	apiParams := url.Values{"model": {"Transportmode"}}
	testStatusCode("StopsAndLinesV2StringXml", t, "SL_STOPS_AND_LINES_V2", StopsAndLinesV2, XML, apiParams, "string", statusSuccessXml)
}

// TestStopsAndLinesV2BytesJson tries to fetch data from StopsAndLinesV2 API
func TestStopsAndLinesV2BytesJson(t *testing.T) {
	apiParams := url.Values{"model": {"Transportmode"}}
	testStatusCode("StopsAndLinesV2BytesJson", t, "SL_STOPS_AND_LINES_V2", StopsAndLinesV2, JSON, apiParams, "bytes", statusSuccessJson)
}

// TestStopsAndLinesV2BytesXml tries to fetch data from StopsAndLinesV2 API
func TestStopsAndLinesV2BytesXml(t *testing.T) {
	apiParams := url.Values{"model": {"Transportmode"}}
	testStatusCode("StopsAndLinesV2BytesXml", t, "SL_STOPS_AND_LINES_V2", StopsAndLinesV2, XML, apiParams, "bytes", statusSuccessXml)
}

// TestStopsAndLinesV2ResponseJson tries to fetch data from StopsAndLinesV2 API
func TestStopsAndLinesV2ResponseJson(t *testing.T) {
	apiParams := url.Values{"model": {"Transportmode"}}
	testStatusCode("StopsAndLinesV2ResponseJson", t, "SL_STOPS_AND_LINES_V2", StopsAndLinesV2, JSON, apiParams, "response", statusSuccessJson)
}

// TestStopsAndLinesV2ResponseXml tries to fetch data from StopsAndLinesV2 API
func TestStopsAndLinesV2ResponseXml(t *testing.T) {
	apiParams := url.Values{"model": {"Transportmode"}}
	testStatusCode("StopsAndLinesV2ResponseXml", t, "SL_STOPS_AND_LINES_V2", StopsAndLinesV2, XML, apiParams, "response", statusSuccessXml)
}

/*
 * NearbyStopsV2
 */

// TestNearbyStopsV2StringJson tries to fetch data from NearbyStopsV2 API
func TestNearbyStopsV2StringJson(t *testing.T) {
	apiParams := url.Values{"originCoordLat": {"59.329323"}, "originCoordLong": {"18.068581"}}
	testStatusCode("NearbyStopsV2StringJson", t, "SL_NEARBY_STOPS_V2", NearbyStopsV2, JSON, apiParams, "string", "stopLocationOrCoordLocation")
}

// TestNearbyStopsV2StringXml tries to fetch data from NearbyStopsV2 API
func TestNearbyStopsV2StringXml(t *testing.T) {
	apiParams := url.Values{"originCoordLat": {"59.329323"}, "originCoordLong": {"18.068581"}}
	testStatusCode("NearbyStopsV2StringXml", t, "SL_NEARBY_STOPS_V2", NearbyStopsV2, XML, apiParams, "string", "<StopLocation id=")
}

// TestNearbyStopsV2BytesJson tries to fetch data from NearbyStopsV2 API
func TestNearbyStopsV2BytesJson(t *testing.T) {
	apiParams := url.Values{"originCoordLat": {"59.329323"}, "originCoordLong": {"18.068581"}}
	testStatusCode("NearbyStopsV2BytesJson", t, "SL_NEARBY_STOPS_V2", NearbyStopsV2, JSON, apiParams, "bytes", "stopLocationOrCoordLocation")
}

// TestNearbyStopsV2BytesXml tries to fetch data from NearbyStopsV2 API
func TestNearbyStopsV2BytesXml(t *testing.T) {
	apiParams := url.Values{"originCoordLat": {"59.329323"}, "originCoordLong": {"18.068581"}}
	testStatusCode("NearbyStopsV2BytesXml", t, "SL_NEARBY_STOPS_V2", NearbyStopsV2, XML, apiParams, "bytes", "<StopLocation id=")
}

// TestNearbyStopsV2ResponseJson tries to fetch data from NearbyStopsV2 API
func TestNearbyStopsV2ResponseJson(t *testing.T) {
	apiParams := url.Values{"originCoordLat": {"59.329323"}, "originCoordLong": {"18.068581"}}
	testStatusCode("NearbyStopsV2ResponseJson", t, "SL_NEARBY_STOPS_V2", NearbyStopsV2, JSON, apiParams, "response", "stopLocationOrCoordLocation")
}

// TestNearbyStopsV2ResponseXml tries to fetch data from NearbyStopsV2 API
func TestNearbyStopsV2ResponseXml(t *testing.T) {
	apiParams := url.Values{"originCoordLat": {"59.329323"}, "originCoordLong": {"18.068581"}}
	testStatusCode("NearbyStopsV2ResponseXml", t, "SL_NEARBY_STOPS_V2", NearbyStopsV2, XML, apiParams, "response", "<StopLocation id=")
}
