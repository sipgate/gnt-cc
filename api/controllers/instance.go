package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gnt-cc/model"
)

type InstanceController struct {
	Repository instanceRepository
}

// GetAll godoc
// @Summary Get all nodes in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.AllInstancesResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/nodes [get]
func (controller *InstanceController) GetAll(c *gin.Context) {
	clusterName := c.Param("cluster")
	instances, err := controller.Repository.GetAll(clusterName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	c.JSON(200, model.AllInstancesResponse{
		Cluster:           clusterName,
		NumberOfInstances: len(instances),
		Instances:         instances,
	})
}

// GetAll godoc
// @Summary Get an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.InstanceResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/nodes/{instance} [get]
func (controller *InstanceController) Get(c *gin.Context) {
	clusterName := c.Param("cluster")
	instanceName := c.Param("instance")

	if instanceName == "" {
		c.AbortWithStatusJSON(400, createErrorBody("instance name is required"))
		return
	}

	result, err := controller.Repository.Get(clusterName, instanceName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	if !result.Found {
		c.AbortWithStatusJSON(404, createErrorBody(
			fmt.Sprintf(MsgInstanceNotFound, instanceName, clusterName)))
	}

	c.JSON(200, model.InstanceResponse{
		Cluster:  clusterName,
		Instance: result.Instance,
	})
}
