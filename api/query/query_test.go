package query_test

import (
	"errors"
	"gnt-cc/mocking"
	"gnt-cc/query"
	"gnt-cc/rapi_client"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var performer = &query.Performer{}

func createValidConfig() query.RequestConfig {
	return query.RequestConfig{
		ClusterName:  "test",
		ResourceType: "node",
		Fields:       []string{"one", "two"},
	}
}

func TestPerformReturnsError_WithInvalidRequestConfig(t *testing.T) {
	client := mocking.NewRAPIClient()
	invalidConfigs := []query.RequestConfig{{
		ClusterName:  "",
		ResourceType: "node",
		Fields:       []string{"one"},
	}, {
		ClusterName:  "test",
		ResourceType: "invalid",
		Fields:       []string{"one"},
	}, {
		ClusterName:  "test",
		ResourceType: "instance",
		Fields:       []string{},
	}}

	for _, c := range invalidConfigs {
		_, err := performer.Perform(client, c)
		assert.Error(t, err)
	}
}

func TestPerformReturnsNoError_WithValidRequestConfig(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Body: "{}"}, nil)
	_, err := performer.Perform(client, createValidConfig())
	assert.NoError(t, err)
}

func TestPerformReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected"))
	_, err := performer.Perform(client, createValidConfig())
	assert.EqualError(t, err, "expected")
}

func TestPerformReturnsError_WhenRAPIClientReturnsNon200StatusCode(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{
		Status: 502,
	}, nil)
	_, err := performer.Perform(client, createValidConfig())
	assert.EqualError(t, err, "RAPI returned status code 502")
}

func TestPerformReturnsError_WhenRAPIClientReturnsInvalidJSONBody(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Body: "{"}, nil)
	_, err := performer.Perform(client, createValidConfig())
	assert.Error(t, err)
}

func TestPerformCallsRAPIClientGetFuncCorrectly(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", "test", "/2/query/instance?fields=one,two").
		Once().Return(rapi_client.Response{Body: "{}"}, nil)
	config := query.RequestConfig{
		ClusterName:  "test",
		ResourceType: "instance",
		Fields:       []string{"one", "two"},
	}
	_, err := performer.Perform(client, config)
	assert.NoError(t, err)
	client.AssertExpectations(t)
}

func TestPerformReturnsResourcesArray(t *testing.T) {
	validResponse, err := ioutil.ReadFile("../testfiles/rapi_responses/valid_query_response.json")
	assert.NoError(t, err)

	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Body: string(validResponse)}, nil)
	resources, err := performer.Perform(client, createValidConfig())
	assert.NoError(t, err)
	assert.EqualValues(t, []query.Resource{{
		"name":  "bart",
		"pnode": "node4",
	}, {
		"name":  "smithers",
		"pnode": "node2",
	}}, resources)
}
