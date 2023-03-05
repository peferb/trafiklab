package sl

import (
	"io"
	"net/http"
	"net/url"
)

func GetBytes(apiUrl string, query url.Values, key string) ([]byte, error) {
	query.Set("key", key)
	u, _ := url.Parse(apiUrl)
	u.RawQuery = query.Encode()

	res, err := http.Get(u.String())
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
