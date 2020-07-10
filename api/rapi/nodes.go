package rapi

import (
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/model"
	"gnt-cc/utils"
)

func filterInstances(arr []model.GntInstance, cond func(model.GntInstance) bool) []model.GntInstance {
	result := []model.GntInstance{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func GetNodes(clusterConfig config.ClusterConfig) ([]model.GntNode, error) {
	response, err := Get(clusterConfig, "/2/nodes?bulk=1")

	if err != nil {
		return nil, err
	}

	var nodeData []Node
	err = json.Unmarshal([]byte(response), &nodeData)

	nodes := make([]model.GntNode, len(nodeData))

	for i, node := range nodeData {
		nodes[i] = model.GntNode{
			Name:        node.Name,
			MemoryTotal: node.Mtotal,
			MemoryFree:  node.Mfree,
		}
	}

	return nodes, nil
}

func GetNode(clusterConfig config.ClusterConfig, nodeName string) (model.GntNode, error) {
	response, err := Get(clusterConfig, fmt.Sprintf("/2/nodes/%s", nodeName))

	if err != nil {
		return model.GntNode{}, err
	}

	clusterInstances, err := GetInstances(clusterConfig)

	if err != nil {
		return model.GntNode{}, err
	}

	var nodeData Node
	err = json.Unmarshal([]byte(response), &nodeData)

	if err != nil {
		return model.GntNode{}, err
	}

	return model.GntNode{
		Name:        nodeData.Name,
		MemoryTotal: nodeData.Mtotal,
		MemoryFree:  nodeData.Mfree,
		PrimaryInstances: filterInstances(clusterInstances, func(instance model.GntInstance) bool {
			return instance.PrimaryNode == nodeName
		}),
		SecondaryInstances: filterInstances(clusterInstances, func(instance model.GntInstance) bool {
			return utils.IsInSlice(nodeName, instance.SecondaryNodes)
		}),
	}, nil
}
