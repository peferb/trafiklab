package realtime_departures_v4

import (
	"encoding/json"
	sl "github.com/peferb/trafiklab/sl"
	"net/url"
)

// ApiUrl "Realtime Departures V4 (Realtidsinformation 4)"
//
// Access
// https://developer.trafiklab.se/api/sl-realtidsinformation-4
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/departures-4/
const ApiUrl = "https://api.sl.se/api2/realtimedeparturesV4.json"

func GetBytes(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrl)
}

func Get(apiKey string, query url.Values, response *Response) error {
	bytes, err := sl.GetBytes(
		apiKey,
		query,
		ApiUrl,
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
		ApiUrl,
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
		ApiUrl,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
