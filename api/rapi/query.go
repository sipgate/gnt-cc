package rapi

import (
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"strings"
)

type ResourceTransformCallback func(ResourceFieldValueMap) error

type ResourceFieldValueMap map[string]interface{}

func queryAndTransformResources(config queryRequestConfig, transformCallback ResourceTransformCallback) error {
	queryUrl := buildQueryUrl(config.resourceType, config.fields)
	queryResponse, err := performAndParseQueryRequest(config.clusterConfig, queryUrl)

	if err != nil {
		return err
	}

	var resourceFieldValueMap = make(ResourceFieldValueMap)

	for _, resource := range queryResponse.Data {
		for fieldIndex, field := range queryResponse.Fields {
			resourceFieldValueMap[field.Name] = resource[fieldIndex][1]
		}

		err = transformCallback(resourceFieldValueMap)

		if err != nil {
			return err
		}
	}

	return nil
}

type queryField struct {
	Doc   string `json:"doc"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type queryResponse struct {
	Fields []queryField      `json:"fields"`
	Data   [][][]interface{} `json:"data"`
}

type queryRequestConfig struct {
	clusterConfig config.ClusterConfig
	resourceType  string
	fields        []string
}

func buildQueryUrl(resourceType string, fields []string) string {
	return fmt.Sprintf("/2/query/%s?fields=%s", resourceType, strings.Join(fields, ","))
}

func performAndParseQueryRequest(clusterConfig config.ClusterConfig, queryUrl string) (queryResponse, error) {
	jsonResponse, err := Get(clusterConfig, queryUrl)

	if err != nil {
		return queryResponse{}, err
	}

	var response queryResponse
	err = json.Unmarshal([]byte(jsonResponse), &response)

	return response, err
}
