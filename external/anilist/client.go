package anilist

import (
	"context"
	"log"
	"net/http"
	"tsuki/external/queries"

	"github.com/machinebox/graphql"
)

type HeaderTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

func (t *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.Headers {
		req.Header.Set(key, value)
	}

	if t.Transport == nil {
		t.Transport = http.DefaultTransport
	}
	return t.Transport.RoundTrip(req)
}

var CLIENT *graphql.Client
var TOKEN string

func SetupClient(token string) {
	if CLIENT != nil && TOKEN == token {
		return
	}

	TOKEN = token
	httpClient := http.Client{
		Transport: &HeaderTransport{
			Headers: map[string]string{
				"Authorization": "Bearer " + TOKEN,
			},
		},
	}

	CLIENT = graphql.NewClient("https://graphql.anilist.co", graphql.WithHTTPClient(&httpClient))
}

// TODO: What happens when anilist responds with an error code?
func BuildAndSendRequest[T any](queryName string) (*T, error) {
	request, err := buildRequest(queryName)
	if err != nil {
		// This should never happen. If it does, it points to an implementation error.
		log.Fatal("Could not find Anilist query")
	}

	ctx := context.Background()
	var response T
	if err := CLIENT.Run(ctx, request, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func buildRequest(queryName string) (*graphql.Request, error) {
	bytes, err := queries.QUERIES.ReadFile(queryName + ".graphql")
	if err != nil {
		return nil, err
	}
	query := string(bytes)
	return graphql.NewRequest(query), nil
}
