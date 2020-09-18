package controllers

import (
	"github.com/gin-gonic/gin"
	"gnt-cc/model"
	"net/http"
)

type StatisticsController struct {
	InstanceRepository instanceRepository
	NodeRepository     nodeRepository
}

func (controller *StatisticsController) Get(c *gin.Context) {
	clusterName := c.Param("cluster")

	instances, err := controller.InstanceRepository.GetAll(clusterName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	nodes, err := controller.NodeRepository.GetAll(clusterName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	response := model.StatisticsResponse{}

	for _, instance := range instances {
		response.Instances.Count++
		response.Instances.CPUCount += instance.CpuCount
		response.Instances.MemoryTotal += instance.MemoryTotal
	}

	for _, node := range nodes {
		response.Nodes.Count++
		response.Nodes.CPUCount += node.CPUCount
		response.Nodes.MemoryTotal += node.MemoryTotal
	}

	c.JSON(http.StatusOK, response)
}
