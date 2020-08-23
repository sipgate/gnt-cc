package rapi_client_test

import (
	"bytes"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/mocking"
	"gnt-cc/rapi_client"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func prepareNonExistingClusterClient() rapi_client.Client {
	httpClient := mocking.NewHTTPClient()
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	return client
}

func TestGetFuncReturnsError_WhenANonExistingClusterIsPassedIn(t *testing.T) {
	client := prepareNonExistingClusterClient()
	_, err := client.Get("test1", "/")
	assert.NotNil(t, err)
}

func TestPostFuncReturnsError_WhenANonExistingClusterIsPassedIn(t *testing.T) {
	client := prepareNonExistingClusterClient()
	_, err := client.Post("test1", "/", "")
	assert.NotNil(t, err)
}

func getTestClustersWithURLs() (clusters []config.ClusterConfig, urls []string) {
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
		}, []string{
			"https://test:supersecret@test1.gnt:5080/",
			"http://test:supersecret3@test3.other.gnt:5090/",
		}
}

func TestGetMethodCallsCorrectURLs(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	clusters, urls := getTestClustersWithURLs()
	for _, url := range urls {
		httpClient.On("Get", url+"test").
			Once().
			Return(mocking.MakeSuccessResponse("Test"), nil)
	}
	client, err := rapi_client.New(httpClient, clusters)
	assert.NoError(t, err)
	for _, cluster := range clusters {
		_, err = client.Get(cluster.Name, "/test")
		assert.NoError(t, err)
	}
	httpClient.AssertExpectations(t)
}

func TestPostMethodCallsCorrectURLs(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	clusters, urls := getTestClustersWithURLs()
	for _, url := range urls {
		httpClient.On("Post", url+"test", mock.Anything, mock.Anything).
			Once().
			Return(mocking.MakeSuccessResponse("Test"), nil)
	}
	client, err := rapi_client.New(httpClient, clusters)
	assert.NoError(t, err)
	for _, cluster := range clusters {
		_, err = client.Post(cluster.Name, "/test", "")
		assert.NoError(t, err)
	}
	httpClient.AssertExpectations(t)
}

func TestPostMethodWillReturnError_WhenRequestBodyIsInvalid(t *testing.T) {
	invalidBody := struct {
		Test func() `json:"test"`
	}{Test: func() {}}

	httpClient := mocking.NewHTTPClient()
	client, err := rapi_client.New(httpClient, []config.ClusterConfig{createTestConfigSSL("test1")})
	assert.NoError(t, err)
	_, err = client.Post("test1", "/test", invalidBody)

	httpClient.AssertNotCalled(t, "Post")
	assert.NotNil(t, err)
}

func TestGetMethodWillReturnErrorThrownByHTTPClient(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.On("Get", mock.Anything).
		Once().
		Return(nil, fmt.Errorf("fail"))
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	_, err := client.Get("test", "/")
	assert.NotNil(t, err)
}

func TestPostMethodWillReturnErrorThrownByHTTPClient(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.On("Post", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, fmt.Errorf("fail"))
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	_, err := client.Post("test", "/", "")
	assert.NotNil(t, err)
}

func TestGetFuncWillCorrectlyParseResponseBody(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.
		On("Get", mock.Anything).
		Once().
		Return(mocking.MakeSuccessResponse(`{ "test": "Test" }`), nil)
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	response, err := client.Get("test", "/")
	assert.NoError(t, err)
	assert.EqualValues(t, 200, response.Status)
	assert.EqualValues(t, `{ "test": "Test" }`, response.Body)
}

func TestPostFuncWillCorrectlyParseResponseBody(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.
		On("Post", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(mocking.MakeSuccessResponse(`{ "test": "Test" }`), nil)
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	response, err := client.Post("test", "/", "")
	assert.NoError(t, err)
	assert.EqualValues(t, 200, response.Status)
	assert.EqualValues(t, `{ "test": "Test" }`, response.Body)
}

func TestPostFuncWillCorrectlyPrepareRequestBody(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.
		On("Post", mock.Anything, mock.Anything, bytes.NewBufferString(`{"test":"Test"}`)).
		Once().
		Return(mocking.MakeSuccessResponse(""), nil)
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	body := struct {
		Test string `json:"test"`
	}{
		Test: "Test",
	}
	_, _ = client.Post("test", "/", body)
	httpClient.AssertExpectations(t)
}

func TestGetFuncWillReturnErrorHTTPStatusCode(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.
		On("Get", mock.Anything).
		Once().
		Return(mocking.MakeNotFoundResponse(), nil)
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	response, _ := client.Get("test", "/")
	assert.EqualValues(t, 404, response.Status)
}

func TestPostFuncWillReturnErrorHTTPStatusCode(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	httpClient.
		On("Post", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(mocking.MakeNotFoundResponse(), nil)
	client, _ := rapi_client.New(httpClient, getDefaultTestClusters())
	response, _ := client.Post("test", "/", "")
	assert.EqualValues(t, 404, response.Status)
}

func TestGetFuncReturnsError_WhenBodyCannotBeRead(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	client, _ := rapi_client.New(httpClient, []config.ClusterConfig{createTestConfigSSL("test1")})
	httpClient.On("Get", mock.Anything).
		Once().
		Return(makeResponseWithBodyReaderReturningAnError(), nil)
	_, err := client.Get("test1", "/")
	assert.NotNil(t, err)
}

func TestPostFuncReturnsError_WhenBodyCannotBeRead(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	client, _ := rapi_client.New(httpClient, []config.ClusterConfig{createTestConfigSSL("test1")})
	httpClient.On("Post", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(makeResponseWithBodyReaderReturningAnError(), nil)
	_, err := client.Post("test1", "/", "")
	assert.NotNil(t, err)
}
