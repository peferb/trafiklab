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

func testStatusCode(testName string, t *testing.T, envVar string, api Type, format DataFormat, apiParams url.Values, resType string, shouldContain string) {
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
