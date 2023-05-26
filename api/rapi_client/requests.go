package rapi_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"syscall"
)

func (client rapiClient) Get(clusterName string, slug string) (Response, error) {
	clusterURL, exists := client.clusterUrls[clusterName]

	if !exists {
		return Response{}, fmt.Errorf("cluster not found: %s", clusterName)
	}

	httpResponse, err := client.http.Get(clusterURL + slug)

	if errors.Is(err, syscall.ECONNREFUSED) {
		return Response{}, fmt.Errorf("could not connect to cluster '%s'", clusterName)
	}

	if err != nil {
		return Response{}, fmt.Errorf("could not send request to cluster '%s': %s", clusterName, err)
	}

	return parseRAPIResponse(httpResponse)
}

func (client rapiClient) Post(clusterName string, slug string, body interface{}) (Response, error) {
	return client.modify(clusterName, slug, body, http.MethodPost)
}

func (client rapiClient) Put(clusterName string, slug string, body interface{}) (Response, error) {
	return client.modify(clusterName, slug, body, http.MethodPut)
}

func (client rapiClient) modify(clusterName string, slug string, body interface{}, method string) (Response, error) {
	clusterURL, exists := client.clusterUrls[clusterName]

	if !exists {
		return Response{}, fmt.Errorf("cluster '%s' not found", clusterName)
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return Response{}, fmt.Errorf("could not prepare request to cluster '%s': %s", clusterName, err)
	}

	request, err := http.NewRequest(
		method,
		clusterURL+slug,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return Response{}, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.http.Do(request)

	if err != nil {
		return Response{}, fmt.Errorf("request error: %s", err)
	}

	return parseRAPIResponse(response)
}

func parseRAPIResponse(httpResponse *http.Response) (Response, error) {
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		return Response{}, fmt.Errorf("could not parse RAPI response: %s", err)
	}

	return Response{
		Status: httpResponse.StatusCode,
		Body:   string(body),
	}, nil
}
