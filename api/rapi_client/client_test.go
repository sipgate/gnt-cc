package rapi_client_test

import (
	"fmt"
	"gnt-cc/config"
	"gnt-cc/mocking"
	"gnt-cc/rapi_client"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestConfigNoSSL(name string) config.ClusterConfig {
	return config.ClusterConfig{
		Name:     name,
		Hostname: "test.gnt",
		Port:     5080,
		Username: "test",
		Password: "supersecret",
		SSL:      false,
	}
}

func createTestConfigSSL(name string) config.ClusterConfig {
	return config.ClusterConfig{
		Name:     name,
		Hostname: "test.gnt",
		Port:     5080,
		Username: "test",
		Password: "supersecret",
		SSL:      true,
	}
}

func getDefaultTestClusters() []config.ClusterConfig {
	return []config.ClusterConfig{createTestConfigSSL("test")}
}

type errorReader struct{}

func (m *errorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("error")
}

func makeResponseWithBodyReaderReturningAnError() *http.Response {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(&errorReader{}),
		Header:     make(http.Header),
	}
}

func TestCreatingRAPIClientShouldNotBePossible_WhenAClusterConfigHasNoName(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	client, err := rapi_client.New(httpClient, []config.ClusterConfig{{
		Hostname: "test",
	}, {
		Name: "test",
	}})

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestCreatingRAPIClientShouldNotBePossible_WhenClusterNamesAreNotUnique(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	client, err := rapi_client.New(httpClient, []config.ClusterConfig{{
		Name: "test",
	}, {
		Name: "test",
	}})

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestCreatingRAPIClientShouldBePossible_WhenAllClusterConfigsAreValid(t *testing.T) {
	httpClient := mocking.NewHTTPClient()
	client, err := rapi_client.New(httpClient, []config.ClusterConfig{
		createTestConfigNoSSL("test1"),
		createTestConfigSSL("test2"),
	})

	assert.Nil(t, err)
	assert.NotNil(t, client)
}
