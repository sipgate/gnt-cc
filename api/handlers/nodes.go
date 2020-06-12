package handlers

import (
	"errors"
	"fmt"
	"gnt-cc/httputil"
	"gnt-cc/model"
	"gnt-cc/utils"
	"math/rand"

	"github.com/gin-gonic/gin"
)

// FindAllNodes godoc
// @Summary Find all nodes
// @Description ...
// @Produce json
// @Success 200 {object} model.AllNodesResponse
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/nodes [get]
func FindAllNodes(c *gin.Context) {
	name := c.Param("cluster")
	if !utils.IsValidCluster(name) {
		httputil.NewError(c, 404, errors.New("cluster not found"))
	} else {
		// TODO: only sends dummy data for now
		//content, err := rapi.Get(name, "/2/nodes?bulk=1")
		//if err != nil {
		//	httputil.NewError(c, 502, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
		//	return
		//}
		//var nodesData rapi.NodesBulk
		/*json.Unmarshal([]byte(content), &nodesData)
		c.JSON(200, gin.H{
			"cluster":       name,
			"numberOfNodes": len(nodesData),
			"nodes":         nodesData,
		})*/
		dummyCount := 150
		dummyNodes := make([]model.GntNode, dummyCount)

		for i := 0; i < dummyCount; i++ {
			dummyNodes[i] = model.GntNode{
				Name:        fmt.Sprintf("dummy_node_%s_%d", name, i),
				MemoryTotal: 2000,
				MemoryFree:  rand.Intn(2000),
			}
		}

		c.JSON(200, model.AllNodesResponse{
			Cluster:       name,
			NumberOfNodes: 200,
			Nodes:         dummyNodes,
		})
	}
}
