package anilist

import (
	"context"
	"net/http"

	"tsuki/core"
	"tsuki/external/queries"

	"github.com/machinebox/graphql"
)

type HeaderTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

type ClientInterface interface {
	Run(ctx context.Context, req *graphql.Request, resp interface{}) error
}

type GraphQLClient struct {
	*graphql.Client
}

type GraphQLVariable struct {
	Key   string
	Value interface{}
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

// Send a request to Anilist. Optionally pass a custom GraphQL client to send the request.
// TODO: What happens when anilist responds with an error code?
func BuildAndSendRequest[T any](queryName string, token string, customClient ClientInterface, variables ...GraphQLVariable) (*T, error) {
	request, err := buildRequest(queryName)
	if err != nil {
		// This should never happen. If it does, it points to an implementation error.
		core.CONFIG.Logger.Fatal("Could not find Anilist query")
	}

	var client ClientInterface
	if customClient != nil {
		client = customClient
	} else {
		client = buildClient(token)
	}

	for _, variable := range variables {
		request.Var(variable.Key, variable.Value)
	}

	ctx := context.Background()
	var response T
	if err := client.Run(ctx, request, &response); err != nil {
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

func buildClient(token string) ClientInterface {
	transport := http.DefaultTransport
	if token != "" {
		transport = &HeaderTransport{
			Headers: map[string]string{
				"Authorization": "Bearer " + token,
			},
		}
	}

	httpClient := http.Client{
		Transport: transport,
	}

	return GraphQLClient{graphql.NewClient("https://graphql.anilist.co", graphql.WithHTTPClient(&httpClient))}
}
