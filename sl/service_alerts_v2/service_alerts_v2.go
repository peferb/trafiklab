// Package service_alerts_v2 "SL Service alerts v2.0 AKA Deviations (Störningsinformation 2)"
//
// Access
// https://developer.trafiklab.se/api/sl-storningsinformation-2
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/service-alerts-2/
package service_alerts_v2

import (
	sl "github.com/peferb/trafiklab/sl"
	"net/url"
)

const ApiUrlJson = "https://api.sl.se/api2/deviations.json"
const ApiUrlXML = "https://api.sl.se/api2/deviations.xml"

// GetBytes JSON
func GetBytes(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrlJson)
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
		ApiUrlXML,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
