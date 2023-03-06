// Package realtime_departures_v4 "Realtime Departures V4 (Realtidsinformation 4)"
//
// Access
// https://developer.trafiklab.se/api/sl-realtidsinformation-4
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/
package realtime_departures_v4

import (
	"encoding/json"
	sl "github.com/peferb/trafiklab/sl"
	"net/url"
)

const ApiUrlJson = "https://api.sl.se/api2/realtimedeparturesV4.json"
const ApiUrlXml = "https://api.sl.se/api2/realtimedeparturesV4.xml"

func GetBytesJson(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrlJson)
}

func GetBytesXml(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrlXml)
}

func Get(apiKey string, query url.Values, response *Response) error {
	bytes, err := sl.GetBytes(
		apiKey,
		query,
		ApiUrlJson,
	)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &response)
}

func GetJson(apiKey string, query url.Values) (string, error) {
	bytes, err := sl.GetBytes(
		apiKey,
		query,
		ApiUrlJson,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetXml(apiKey string, query url.Values) (string, error) {
	bytes, err := sl.GetBytes(
		apiKey,
		query,
		ApiUrlXml,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
