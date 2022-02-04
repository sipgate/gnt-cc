package rapi_client_test

import (
	"gnt-cc/config"
	"gnt-cc/rapi_client"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func getTestClusters() []config.ClusterConfig {
	return []config.ClusterConfig{
		{
			Name:     "test1",
			Hostname: "test1.gnt",
			Port:     5080,
			Username: "test",
			Password: "supersecret",
			SSL:      true,
		},
		{
			Name:     "test3",
			Hostname: "test3.other.gnt",
			Port:     5090,
			Username: "test",
			Password: "supersecret3",
			SSL:      false,
		},
	}
}

func TestGetMethodCallsCorrectURLs(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("GET", "=~.*",
		httpmock.NewStringResponder(200, "Test"))

	clusters := getTestClusters()
	client, err := rapi_client.New(clusters, mock)

	assert.NoError(t, err)

	for _, cluster := range clusters {
		_, err = client.Get(cluster.Name, "/test")
		assert.NoError(t, err)
	}

	info := mock.GetCallCountInfo()
	assert.Equal(t, 1, info["GET https://test:supersecret@test1.gnt:5080/test"])
	assert.Equal(t, 1, info["GET http://test:supersecret3@test3.other.gnt:5090/test"])
}

func TestPostMethodCallsCorrectURLs(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("POST", "=~.*",
		httpmock.NewStringResponder(200, "Test"))

	clusters := getTestClusters()
	client, err := rapi_client.New(clusters, mock)

	assert.NoError(t, err)

	for _, cluster := range clusters {
		_, err = client.Post(cluster.Name, "/test", "")
		assert.NoError(t, err)
	}

	info := mock.GetCallCountInfo()
	assert.Equal(t, 1, info["POST https://test:supersecret@test1.gnt:5080/test"])
	assert.Equal(t, 1, info["POST http://test:supersecret3@test3.other.gnt:5090/test"])
}

func TestPostMethodWillReturnError_WhenRequestBodyIsInvalid(t *testing.T) {
	invalidBody := struct {
		Test func() `json:"test"`
	}{Test: func() {}}

	mock := httpmock.NewMockTransport()

	client, err := rapi_client.New([]config.ClusterConfig{createTestConfigSSL("test1")}, mock)
	assert.NoError(t, err)
	_, err = client.Post("test1", "/test", invalidBody)

	assert.NotNil(t, err)
	info := mock.GetCallCountInfo()
	assert.Equal(t, 0, info["POST https://test:supersecret@test1.gnt:5080/test"])
}

func TestGetFuncWillCorrectlyParseResponseBody(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("GET", "=~.*",
		httpmock.NewStringResponder(200, `{ "test": "Test" }`))

	client, _ := rapi_client.New(getDefaultTestClusters(), mock)
	response, err := client.Get("test", "/")
	assert.NoError(t, err)
	assert.EqualValues(t, 200, response.Status)
	assert.EqualValues(t, `{ "test": "Test" }`, response.Body)
}

func TestPostFuncWillCorrectlyParseResponseBody(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("POST", "=~.*",
		httpmock.NewStringResponder(200, `{ "test": "Test" }`))

	client, _ := rapi_client.New(getDefaultTestClusters(), mock)
	response, err := client.Post("test", "/", "")
	assert.NoError(t, err)
	assert.EqualValues(t, 200, response.Status)
	assert.EqualValues(t, `{ "test": "Test" }`, response.Body)
}

func TestGetFuncWillReturnErrorHTTPStatusCode(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("GET", "=~.*",
		httpmock.NewStringResponder(404, ""))

	client, _ := rapi_client.New(getDefaultTestClusters(), mock)
	response, _ := client.Get("test", "/")
	assert.EqualValues(t, 404, response.Status)
}

func TestPostFuncWillReturnErrorHTTPStatusCode(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("POST", "=~.*",
		httpmock.NewStringResponder(404, ""))

	client, _ := rapi_client.New(getDefaultTestClusters(), mock)
	response, _ := client.Post("test", "/", "")
	assert.EqualValues(t, 404, response.Status)
}
