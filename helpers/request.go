package helpers

import (
	"net/http"
	"time"
)

func BuildGetRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = map[string][]string{
		"Accept": {"application/json"},
	}
	return request, nil
}

func SendRequest(url string) (*http.Response, error) {
	client := http.Client{Timeout: 10 * time.Second}

	request, err := BuildGetRequest(url)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
