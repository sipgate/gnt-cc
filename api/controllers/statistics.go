package controllers

import (
	"gnt-cc/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsController struct {
	InstanceRepository instanceRepository
	NodeRepository     nodeRepository
}

// Get godoc
// @Summary Get statistics for a given cluster
// @Description Get statistics for a given cluster
// @Produce json
// @Success 200 {object} model.StatisticsResponse
// @Router /clusters/{cluster}/statistics [get]
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

		if node.IsMaster {
			response.Master = node.Name
		}
	}

	c.JSON(http.StatusOK, response)
}
