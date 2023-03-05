package realtime_departures_v4

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type DeparturesResponse struct {
	StatusCode    int `json:"StatusCode"`
	Message       any `json:"Message"`
	ExecutionTime int `json:"ExecutionTime"`
	ResponseData  struct {
		LatestUpdate string `json:"LatestUpdate"`
		DataAge      int    `json:"DataAge"`
		Metros       []any  `json:"Metros"`
		Buses        []any  `json:"Buses"`
		Trains       []struct {
			SecondaryDestinationName any    `json:"SecondaryDestinationName"`
			GroupOfLine              string `json:"GroupOfLine"`
			TransportMode            string `json:"TransportMode"`
			LineNumber               string `json:"LineNumber"`
			Destination              string `json:"Destination"`
			JourneyDirection         int    `json:"JourneyDirection"`
			StopAreaName             string `json:"StopAreaName"`
			StopAreaNumber           int    `json:"StopAreaNumber"`
			StopPointNumber          int    `json:"StopPointNumber"`
			StopPointDesignation     string `json:"StopPointDesignation"`
			TimeTabledDateTime       string `json:"TimeTabledDateTime"`
			ExpectedDateTime         string `json:"ExpectedDateTime"`
			DisplayTime              string `json:"DisplayTime"`
			JourneyNumber            int    `json:"JourneyNumber"`
			Deviations               any    `json:"Deviations"`
		} `json:"Trains"`
		Trams               []any `json:"Trams"`
		Ships               []any `json:"Ships"`
		StopPointDeviations []any `json:"StopPointDeviations"`
	} `json:"ResponseData"`
}

func get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return bodyBytes, nil
}

func Get(query url.Values, response *DeparturesResponse) error {
	query.Set("key", os.Getenv("SL_REAL_TIME_DEPARTURES_V4"))
	u, _ := url.Parse("https://api.sl.se/api2/realtimedeparturesV4.json")
	u.RawQuery = query.Encode()

	body, err := get(u.String())
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &response)
}
