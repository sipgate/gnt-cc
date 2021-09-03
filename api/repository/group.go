package repository

import (
	"encoding/json"
	"fmt"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
)

type GroupRepository struct {
	RAPIClient rapi_client.Client
}

func (repo GroupRepository) GetAll(clusterName string) ([]model.GntGroup, error) {
	slug := fmt.Sprintf("/2/groups?bulk=1")
	response, err := repo.RAPIClient.Get(clusterName, slug)

	if err != nil {
		return []model.GntGroup{}, err
	}

	var groupsResponse rapiGroupsResponse
	err = json.Unmarshal([]byte(response.Body), &groupsResponse)

	if err != nil {
		return []model.GntGroup{}, err
	}

	groups := make([]model.GntGroup, len(groupsResponse))

	for i, groupResponse := range groupsResponse {
		groups[i] = model.GntGroup{
			UUID: groupResponse.UUID,
			Name: groupResponse.Name,
		}
	}

	return groups, nil
}
