package mocking

import (
	"github.com/stretchr/testify/mock"
	"gnt-cc/query"
	"gnt-cc/rapi_client"
)

type queryPerformer struct {
	mock.Mock
}

func NewQueryPerformer() *queryPerformer {
	return &queryPerformer{}
}

func (mock *queryPerformer) Perform(client rapi_client.Client, config query.RequestConfig) ([]query.Resource, error) {
	args := mock.Called(client, config)

	return args.Get(0).([]query.Resource), args.Error(1)
}
