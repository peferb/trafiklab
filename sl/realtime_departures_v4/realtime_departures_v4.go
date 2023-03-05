package realtime_departures_v4

import (
	"encoding/json"
	. "github.com/peferb/trafiklab/sl"
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

func Get(aipKey string, query url.Values, response *Response) error {
	bytes, err := GetBytes(
		aipKey,
		query,
		ApiUrl,
	)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &response)
}

func GetJson(aipKey string, query url.Values) (string, error) {
	bytes, err := GetBytes(
		aipKey,
		query,
		ApiUrl,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetXml(aipKey string, query url.Values) (string, error) {
	bytes, err := GetBytes(
		aipKey,
		query,
		ApiUrl,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
