package mocking

import (
	"gnt-cc/rapi_client"

	"github.com/stretchr/testify/mock"
)

type rapiClient struct {
	mock.Mock
}

func NewRAPIClient() *rapiClient {
	return new(rapiClient)
}

func (mock *rapiClient) Get(clusterName string, slug string) (rapi_client.Response, error) {
	args := mock.Called(clusterName, slug)

	return args.Get(0).(rapi_client.Response), args.Error(1)
}

func (mock *rapiClient) Post(clusterName string, slug string, body interface{}) (rapi_client.Response, error) {
	args := mock.Called(clusterName, slug, body)

	return args.Get(0).(rapi_client.Response), args.Error(1)
}

func (mock *rapiClient) Put(clusterName string, slug string, body interface{}) (rapi_client.Response, error) {
	args := mock.Called(clusterName, slug, body)

	return args.Get(0).(rapi_client.Response), args.Error(1)
}
