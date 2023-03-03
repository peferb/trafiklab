package trafiklab

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	return string(bodyBytes), nil
}

func GetSLRealtimeDeparturesV4(query url.Values) (string, error) {
	query.Set("key", os.Getenv("SL_REAL_TIME_DEPARTURES_V4"))
	u, _ := url.Parse("https://api.sl.se/api2/realtimedeparturesV4.json")
	u.RawQuery = query.Encode()
	return get(u.String())
}
