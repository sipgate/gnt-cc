package controllers

import (
	"fmt"
	"gnt-cc/model"
	"gnt-cc/websocket"

	"github.com/gin-gonic/gin"
)

type InstanceController struct {
	Repository instanceRepository
	Actions    instanceActions
}

// GetAll godoc
// @Summary Get all instances in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.AllInstancesResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/instances [get]
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

// Reboot godoc
// @Summary Reboot an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobIDResponse
// @Router /clusters/{cluster}/instances/{instance}/reboot [post]
func (controller *InstanceController) Reboot(c *gin.Context) {
	clusterName := c.Param("cluster")
	instanceName := c.Param("instance")

	jobID, err := controller.Actions.Reboot(clusterName, instanceName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	c.JSON(200, model.JobIDResponse{
		JobID: jobID,
	})
}

// Get godoc
// @Summary Get an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.InstanceResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/instances/{instance} [get]
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

// OpenInstanceConsole godoc
// @Summary Open a websocket connection to a cluster instance console
// @Description ...
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/instances/{instance}/console [get]
func (controller *InstanceController) OpenInstanceConsole(c *gin.Context) {
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

	instance := result.Instance

	if !instance.IsRunning {
		c.AbortWithStatusJSON(400, createErrorBody("instance is not running"))
		return
	}

	primaryNode := instance.PrimaryNode
	port := result.NetworkPort

	err = websocket.PassThrough(c.Writer, c.Request, primaryNode, port)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}
}
