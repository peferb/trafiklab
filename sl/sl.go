package sl

import (
	"fmt"
	. "github.com/peferb/trafiklab/sl/api"
	"io"
	"net/http"
	"net/url"
)

func NewApi(apiType string, dataFormat string, apiKey string) Api {
	apiUrl := fmt.Sprintf("%s.%s", apiType, dataFormat)

	return Api{
		GetBytes: func(query url.Values) ([]byte, error) {
			return getBytes(apiKey, query, apiUrl)
		},
		GetString: func(query url.Values) (string, error) {
			bytes, err := getBytes(apiKey, query, apiUrl)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		},
		GetResponse: func(query url.Values) (*http.Response, error) {
			return getResponse(apiKey, query, apiUrl)
		},
	}
}

func getResponse(apiKey string, query url.Values, apiUrl string) (*http.Response, error) {
	query.Set("key", apiKey)
	u, _ := url.Parse(apiUrl)
	u.RawQuery = query.Encode()

	return http.Get(u.String())
}

func getBytes(apiKey string, query url.Values, apiUrl string) ([]byte, error) {
	res, err := getResponse(apiKey, query, apiUrl)
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
