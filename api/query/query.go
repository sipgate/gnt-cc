package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"gnt-cc/rapi_client"
	"strings"
)

var validResources = []string{
	"node",
	"group",
	"network",
	"lock",
	"filter",
	"instance",
	"job",
	"export",
}

type (
	Performer     struct{}
	RequestConfig struct {
		ClusterName  string
		ResourceType string
		Fields       []string
	}
	Resource         map[string]interface{}
	fieldDescription struct {
		Doc   string `json:"doc"`
		Kind  string `json:"kind"`
		Name  string `json:"name"`
		Title string `json:"title"`
	}
	responseBody struct {
		Fields []fieldDescription `json:"fields"`
		Data   [][][]interface{}  `json:"data"`
	}
)

func (p *Performer) Perform(client rapi_client.Client, config RequestConfig) ([]Resource, error) {
	if err := validateRequestConfig(config); err != nil {
		return nil, err
	}

	httpResponse, err := client.Get(
		config.ClusterName,
		buildQuerySlug(config),
	)

	if err != nil {
		return nil, err
	}

	parsedBody, err := parseQueryResponseBody(httpResponse.Body)

	if err != nil {
		return nil, err
	}

	return createResourcesArray(parsedBody), nil
}

func parseQueryResponseBody(response string) (*responseBody, error) {
	var parsed responseBody
	err := json.Unmarshal([]byte(response), &parsed)

	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

func createResourcesArray(parsedResponse *responseBody) []Resource {
	var resources = make([]Resource, len(parsedResponse.Data))

	for i, values := range parsedResponse.Data {
		resources[i] = createResource(parsedResponse.Fields, values)
	}

	return resources
}

func createResource(fields []fieldDescription, values [][]interface{}) Resource {
	resource := make(Resource)

	for i, field := range fields {
		resource[field.Name] = values[i][1]
	}

	return resource
}

func validateRequestConfig(config RequestConfig) error {
	if !isValidResourceType(config.ResourceType) {
		return fmt.Errorf("invalid resource type: %s", config.ResourceType)
	}

	if len(config.Fields) == 0 {
		return fmt.Errorf("fields are required")
	}

	if config.ClusterName == "" {
		return errors.New("cluster name is required")
	}

	return nil
}

func isValidResourceType(resource string) bool {
	for _, r := range validResources {
		if resource == r {
			return true
		}
	}
	return false
}

func buildQuerySlug(config RequestConfig) string {
	queryParams := fmt.Sprintf("?fields=%s", strings.Join(config.Fields, ","))
	return fmt.Sprintf("/2/query/%s%s", config.ResourceType, queryParams)
}
