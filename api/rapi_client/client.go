package rapi_client

import (
	"errors"
	"fmt"
	"gnt-cc/config"
	"io"
	"net/http"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type Client interface {
	Get(clusterName string, slug string) (Response, error)
	Post(clusterName string, slug string, body interface{}) (Response, error)
}

type rapiClient struct {
	clusterUrls map[string]string
	http        HTTPClient
}

type Response struct {
	Status int
	Body   string
}

func New(httpClient HTTPClient, clusterConfigs []config.ClusterConfig) (*rapiClient, error) {
	urlMap, err := validateAndCreateClusterUrls(clusterConfigs)

	if err != nil {
		return nil, fmt.Errorf("invalid cluster config: %s", err)
	}

	return &rapiClient{
		clusterUrls: urlMap,
		http:        httpClient,
	}, nil
}

func validateAndCreateClusterUrls(clusterConfigs []config.ClusterConfig) (map[string]string, error) {
	urlMap := make(map[string]string)

	for _, c := range clusterConfigs {
		if c.Name == "" {
			return nil, errors.New("empty field 'Name'")
		}

		if _, exists := urlMap[c.Name]; exists {
			return nil, fmt.Errorf("duplicate cluster name '%s'", c.Name)
		}

		urlMap[c.Name] = createClusterURL(c)
	}

	return urlMap, nil
}

func createClusterURL(config config.ClusterConfig) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d",
		getProtocol(config.SSL),
		config.Username,
		config.Password,
		config.Hostname,
		config.Port,
	)
}

func getProtocol(useSSL bool) string {
	if useSSL {
		return "https"
	}

	return "http"
}
