package dummy

import (
	"fmt"
	"gnt-cc/model"
	"math/rand"
)

func GetNodes(count int) []model.GntNode {
	dummyNodes := make([]model.GntNode, count)

	for i := 0; i < count; i++ {
		dummyNodes[i] = model.GntNode{
			Name:        fmt.Sprintf("dummy_node_%d", i),
			MemoryTotal: 2000,
			MemoryFree:  rand.Intn(2000),
		}
	}

	return dummyNodes
}
