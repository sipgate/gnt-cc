package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gnt-cc/model"
	"gnt-cc/utils"
)

type NodeController struct {
	Repository         nodeRepository
	InstanceRepository instanceRepository
}

// GetAll godoc
// @Summary Get all nodes in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.AllNodesResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/nodes [get]
func (controller *NodeController) GetAll(c *gin.Context) {
	clusterName := c.Param("cluster")

	nodes, err := controller.Repository.GetAll(clusterName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	c.JSON(200, model.AllNodesResponse{
		Cluster:       clusterName,
		NumberOfNodes: len(nodes),
		Nodes:         nodes,
	})
}

// GetAll godoc
// @Summary Get a node in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.NodeResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/nodes/{node} [get]
func (controller *NodeController) Get(c *gin.Context) {
	clusterName := c.Param("cluster")
	nodeName := c.Param("node")

	if nodeName == "" {
		c.AbortWithStatusJSON(400, createErrorBody("node name is required"))
		return
	}

	nodeResult, err := controller.Repository.Get(clusterName, nodeName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	if !nodeResult.Found {
		c.AbortWithStatusJSON(404, createErrorBody(fmt.Sprintf(MsgNodeNotFound, nodeName, clusterName)))
	}

	instances, err := controller.InstanceRepository.GetAll(clusterName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	node := nodeResult.Node

	primaryInstances := filterInstances(instances, func(instance model.GntInstance) bool {
		return instance.PrimaryNode == node.Name
	})
	secondaryInstances := filterInstances(instances, func(instance model.GntInstance) bool {
		return utils.IsInSlice(node.Name, instance.SecondaryNodes)
	})

	c.JSON(200, model.NodeResponse{
		Cluster:            clusterName,
		Node:               node,
		PrimaryInstances:   instanceArrayNotNil(primaryInstances),
		SecondaryInstances: instanceArrayNotNil(secondaryInstances),
	})
}

func filterInstances(arr []model.GntInstance, closure func(model.GntInstance) bool) []model.GntInstance {
	var result []model.GntInstance
	for i := range arr {
		if closure(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func instanceArrayNotNil(array []model.GntInstance) []model.GntInstance {
	if array == nil {
		return []model.GntInstance{}
	}
	return array
}
