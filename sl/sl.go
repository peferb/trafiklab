package sl

import (
	"io"
	"net/http"
	"net/url"
)

func GetResponse(apiKey string, query url.Values, apiUrl string) (*http.Response, error) {
	query.Set("key", apiKey)
	u, _ := url.Parse(apiUrl)
	u.RawQuery = query.Encode()

	return http.Get(u.String())
}

func GetBytes(apiKey string, query url.Values, apiUrl string) ([]byte, error) {
	res, err := GetResponse(apiKey, query, apiUrl)
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
