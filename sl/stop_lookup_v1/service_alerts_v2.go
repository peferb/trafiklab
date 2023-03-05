// Package stop_lookup_v1 "SL Stop lookup v1.0 AKA TypeAhead (Platsuppslag)"
//
// Access
// https://developer.trafiklab.se/api/sl-platsuppslag
//
// Documentation
// https://www.trafiklab.se/api/trafiklab-apis/sl/stop-lookup/
package stop_lookup_v1

import (
	sl "github.com/peferb/trafiklab/sl"
	"net/url"
)

const ApiUrlJson = "https://api.sl.se/api2/typeahead.json"
const ApiUrlXML = "https://api.sl.se/api2/typeahead.xml"

// GetBytes JSON
func GetBytes(apiKey string, query url.Values) ([]byte, error) {
	return sl.GetBytes(apiKey, query, ApiUrlJson)
}

// TODO Struct setter
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
