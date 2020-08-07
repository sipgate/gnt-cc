package handlers

import (
	"gnt-cc/config"
	"gnt-cc/dummy"
	"gnt-cc/model"
	"gnt-cc/rapi"

	"github.com/gin-gonic/gin"
)

// GetAllNodes godoc
// @Summary Find all nodes
// @Description ...
// @Produce json
// @Success 200 {object} model.AllNodesResponse
// @Failure 404 {object} httputil.HTTPError404
// @Router /clusters/{cluster}/nodes [get]
func GetAllNodes(c *gin.Context) {
	clusterConfig, clusterErr := config.GetClusterConfig(c.Param("cluster"))

	if clusterErr != nil {
		c.AbortWithError(404, clusterErr)
		return
	}

	var nodes []model.GntNode

	if config.Get().DummyMode {
		nodes = dummy.GetNodes(20)
	} else {
		var err error
		nodes, err = rapi.GetNodes(clusterConfig)

		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

	c.JSON(200, model.AllNodesResponse{
		Cluster:       clusterConfig.Name,
		NumberOfNodes: len(nodes),
		Nodes:         nodes,
	})
}

func GetNode(c *gin.Context) {
	clusterConfig, clusterErr := config.GetClusterConfig(c.Param("cluster"))

	if clusterErr != nil {
		c.AbortWithError(404, clusterErr)
		return
	}

	nodeName := c.Param("node")

	var node model.GntNode

	if config.Get().DummyMode {
		node = dummy.GetNode(nodeName)
	} else {
		var err error
		node, err = rapi.GetNode(clusterConfig, nodeName)

		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

	c.JSON(200, model.NodeResponse{
		Cluster: clusterConfig.Name,
		Node:    node,
	})
}
