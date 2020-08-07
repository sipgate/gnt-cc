package rapi

import (
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"strings"
)

type QueryField struct {
	Doc   string `json:"doc"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type QueryResponse struct {
	Fields []QueryField      `json:"fields"`
	Data   [][][]interface{} `json:"data"`
}

type QueryMap map[string]interface{}

func GetQuery(clusterConfig config.ClusterConfig, resource string, fields []string) ([]QueryMap, error) {
	queryString := fmt.Sprintf("/2/query/%s?fields=%s", resource, strings.Join(fields, ","))

	httpResponse, err := Get(clusterConfig, queryString)

	if err != nil {
		return nil, err
	}

	var queryResponse QueryResponse
	err = json.Unmarshal([]byte(httpResponse), &queryResponse)

	var queryEntries = make([]QueryMap, len(queryResponse.Data))

	for index, entry := range queryResponse.Data {
		queryEntries[index] = make(QueryMap)

		for fieldIndex, field := range queryResponse.Fields {
			queryEntries[index][field.Name] = entry[fieldIndex][1]
		}
	}

	return queryEntries, nil
}
