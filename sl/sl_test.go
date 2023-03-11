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
