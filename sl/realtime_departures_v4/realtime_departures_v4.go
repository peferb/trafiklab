package realtime_departures_v4

import (
	"encoding/json"
	. "github.com/peferb/trafiklab/sl"
	"net/url"
	"os"
)

func Get(query url.Values, response *Response) error {
	bytes, err := GetBytes(
		"https://api.sl.se/api2/realtimedeparturesV4.json",
		query,
		os.Getenv("SL_REAL_TIME_DEPARTURES_V4"),
	)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &response)
}

func GetJson(query url.Values) (string, error) {
	bytes, err := GetBytes(
		"https://api.sl.se/api2/realtimedeparturesV4.json",
		query,
		os.Getenv("SL_REAL_TIME_DEPARTURES_V4"),
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetXml(query url.Values) (string, error) {
	bytes, err := GetBytes(
		"https://api.sl.se/api2/realtimedeparturesV4.xml",
		query,
		os.Getenv("SL_REAL_TIME_DEPARTURES_V4"),
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
