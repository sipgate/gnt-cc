package actions

import (
	"fmt"
	"gnt-cc/rapi_client"
	"strconv"
	"strings"
)

type InstanceActions struct {
	RAPIClient rapi_client.Client
}

func (repo *InstanceActions) Start(clusterName string, instanceName string) (int, error) {
	slug := fmt.Sprintf("/2/instances/%s/startup", instanceName)
	response, err := repo.RAPIClient.Put(clusterName, slug, nil)

	if err != nil {
		return 0, err
	}

	jobID, err := strconv.Atoi(strings.TrimSpace(response.Body))

	if err != nil {
		return 0, fmt.Errorf("cannot parse RAPI response")
	}

	return jobID, nil
}

func (repo *InstanceActions) Reboot(clusterName string, instanceName string) (int, error) {
	slug := fmt.Sprintf("/2/instances/%s/reboot", instanceName)
	response, err := repo.RAPIClient.Post(clusterName, slug, nil)

	if err != nil {
		return 0, err
	}

	jobID, err := strconv.Atoi(strings.TrimSpace(response.Body))

	if err != nil {
		return 0, fmt.Errorf("cannot parse RAPI response")
	}

	return jobID, nil
}

func (repo *InstanceActions) Shutdown(clusterName string, instanceName string) (int, error) {
	slug := fmt.Sprintf("/2/instances/%s/shutdown", instanceName)
	response, err := repo.RAPIClient.Put(clusterName, slug, nil)

	if err != nil {
		return 0, err
	}

	jobID, err := strconv.Atoi(strings.TrimSpace(response.Body))

	if err != nil {
		return 0, fmt.Errorf("cannot parse RAPI response")
	}

	return jobID, nil
}
