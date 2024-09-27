package helpers

import "net/http"

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
