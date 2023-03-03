package trafiklab

import (
	"net/http"
	"net/url"
	"os"
)

type SL struct{}

func (receiver SL) GetRealtimeDeparturesV4(query url.Values) (*http.Response, error) {
	query.Set("key", os.Getenv("SL_REAL_TIME_DEPARTURES_V4"))
	u, _ := url.Parse("https://api.sl.se/api2/realtimedeparturesV4.json")
	u.RawQuery = query.Encode()
	return http.Get(u.String())
}
