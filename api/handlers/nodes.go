package handlers

import (
	"errors"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/dummy"
	"gnt-cc/httputil"
	"gnt-cc/model"
	"gnt-cc/rapi"

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
func FindAllNodes(context *gin.Context) {
	clusterName := context.Param("cluster")
	if !config.ClusterExists(clusterName) {
		httputil.NewError(context, 404, errors.New("cluster not found"))
		return
	}

	var nodes []model.GntNode

	if config.Get().DummyMode {
		nodes = dummy.GetNodes(20)
	} else {
		var err error
		nodes, err = rapi.GetNodes(clusterName)

		if err != nil {
			httputil.NewError(context, 500, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
	}

	context.JSON(200, model.AllNodesResponse{
		Cluster:       clusterName,
		NumberOfNodes: len(nodes),
		Nodes:         nodes,
	})
}

func FindNode(context *gin.Context) {
	clusterName := context.Param("cluster")
	nodeName := context.Param("node")

	if !config.ClusterExists(clusterName) {
		httputil.NewError(context, 404, errors.New("cluster not found"))
		return
	}

	var node model.GntNode

	if config.Get().DummyMode {
		node = dummy.GetNode(nodeName)
	} else {
		var err error
		node, err = rapi.GetNode(clusterName, nodeName)

		if err != nil {
			httputil.NewError(context, 500, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
	}

	context.JSON(200, model.NodeResponse{
		Cluster: clusterName,
		Node:    node,
	})
}
