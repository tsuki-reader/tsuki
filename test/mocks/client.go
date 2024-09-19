package mocks

import (
	"context"
	"encoding/json"
	"os"

	"github.com/machinebox/graphql"
)

type MockClient struct {
	ResponseFile string
}

func (m *MockClient) Run(ctx context.Context, req *graphql.Request, resp interface{}) error {
	data, err := os.ReadFile(m.ResponseFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, resp)
}
