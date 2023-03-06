// Package traffic_status_v2 "SL Traffic status v2.0 AKA Traffic Situation (SL Trafikläget 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-trafiklaget-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/sl-traffic-status-2/
package traffic_status_v2

import (
	sl "github.com/peferb/trafiklab/sl"
	"net/url"
)

const ApiUrlJson = "https://api.sl.se/api2/trafficsituation.json"
const ApiUrlXml = "https://api.sl.se/api2/trafficsituation.xml"

func GetBytesJson(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrlJson)
}

func GetBytesXml(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrlXml)
}

// TODO Struct getter
// TODO Response struct

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