package rapi

import (
	"encoding/json"
	"fmt"
	"gnt-cc/model"
)

func GetNodes(clusterName string) ([]model.GntNode, error) {
	response, err := Get(clusterName, "/2/nodes?bulk=1")

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

func GetNode(clusterName string, nodeName string) (model.GntNode, error) {
	response, err := Get(clusterName, fmt.Sprintf("/2/node/%s", nodeName))

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
