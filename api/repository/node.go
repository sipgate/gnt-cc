package repository

import (
	"encoding/json"
	"fmt"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
)

type NodeRepository struct {
	RAPIClient rapi_client.Client
}

func (repo *NodeRepository) Get(clusterName string, nodeName string) (model.NodeResult, error) {
	slug := fmt.Sprintf("/2/nodes/%s", nodeName)
	response, err := repo.RAPIClient.Get(clusterName, slug)

	if err != nil {
		return model.NodeResult{}, err
	}

	if response.Status == 404 {
		return model.NodeResult{
			Found: false,
		}, nil
	}

	var node rapiNodeResponse
	err = json.Unmarshal([]byte(response.Body), &node)

	if err != nil {
		return model.NodeResult{}, err
	}

	return model.NodeResult{
		Found: true,
		Node: model.GntNode{
			Name:               node.Name,
			MemoryTotal:        node.Mtotal,
			MemoryFree:         node.Mfree,
			DiskTotal:          node.Dtotal,
			CPUCount:           node.Ctotal,
			PrimaryInstances:   node.PinstList,
			SecondaryInstances: node.SinstList,
		},
	}, nil
}

func (repo *NodeRepository) GetAll(clusterName string) ([]model.GntNode, error) {
	response, err := repo.RAPIClient.Get(clusterName, "/2/nodes?bulk=1")

	if err != nil {
		return nil, err
	}

	var nodeData []rapiNodeResponse
	err = json.Unmarshal([]byte(response.Body), &nodeData)

	if err != nil {
		return nil, err
	}

	nodes := make([]model.GntNode, len(nodeData))

	for i, node := range nodeData {
		nodes[i] = model.GntNode{
			Name:               node.Name,
			MemoryTotal:        node.Mtotal,
			MemoryFree:         node.Mfree,
			DiskTotal:          node.Dtotal,
			CPUCount:           node.Ctotal,
			PrimaryInstances:   node.PinstList,
			SecondaryInstances: node.SinstList,
		}
	}

	return nodes, nil
}
