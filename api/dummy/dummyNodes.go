package dummy

import (
	"fmt"
	"gnt-cc/model"
	"math/rand"
)

func GetNodes(count int) []model.GntNode {
	dummyNodes := make([]model.GntNode, count)

	for i := 0; i < count; i++ {
		dummyNodes[i] = GetNode(fmt.Sprintf("dummy_node_%d", i))
	}

	return dummyNodes
}

func GetNode(name string) model.GntNode {
	return model.GntNode{
		Name:        name,
		MemoryTotal: 2000,
		MemoryFree:  rand.Intn(2000),
	}
}
