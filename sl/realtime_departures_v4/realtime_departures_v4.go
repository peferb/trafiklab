package realtime_departures_v4

import (
	"encoding/json"
	. "github.com/peferb/trafiklab/sl"
	"net/url"
)

func Get(aipKey string, query url.Values, response *Response) error {
	bytes, err := GetBytes(
		aipKey,
		query,
		"https://api.sl.se/api2/realtimedeparturesV4.json",
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
		"https://api.sl.se/api2/realtimedeparturesV4.json",
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
		"https://api.sl.se/api2/realtimedeparturesV4.xml",
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
