package rapi

import (
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/model"
)

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

	var nodeData Node
	err = json.Unmarshal([]byte(response), &nodeData)

	if err != nil {
		return model.GntNode{}, err
	}

	return model.GntNode{
		Name:        nodeData.Name,
		MemoryTotal: nodeData.Mtotal,
		MemoryFree:  nodeData.Mfree,
	}, nil
}
