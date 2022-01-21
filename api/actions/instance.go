package actions

import (
	"fmt"
	"gnt-cc/rapi_client"
	"strconv"
	"strings"
)

type rapiActionToMethodMap map[string](func(string, string, interface{}) (rapi_client.Response, error))

type InstanceActions struct {
	RAPIClient rapi_client.Client
}

func (actions *InstanceActions) PerformSimpleInstanceAction(clusterName string, instanceName string, rapiAction string) (int, error) {
	rapiActionToMethodMapping := rapiActionToMethodMap{
		"startup":  actions.RAPIClient.Put,
		"reboot":   actions.RAPIClient.Post,
		"shutdown": actions.RAPIClient.Put,
	}

	rapiMethod, exists := rapiActionToMethodMapping[rapiAction]

	if !exists {
		return 0, fmt.Errorf("cannot find rapiClient function for action '%s'", rapiAction)
	}

	slug := fmt.Sprintf("/2/instances/%s/%s", instanceName, rapiAction)
	response, err := rapiMethod(clusterName, slug, nil)

	if err != nil {
		return 0, err
	}

	jobID, err := strconv.Atoi(strings.TrimSpace(response.Body))

	if err != nil {
		return 0, fmt.Errorf("cannot parse RAPI response")
	}

	return jobID, nil
}
