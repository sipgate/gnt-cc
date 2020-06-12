package rapi

import (
	"encoding/json"
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

