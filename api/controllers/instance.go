package controllers

import (
	"fmt"
	"gnt-cc/model"
	"gnt-cc/websocket"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type InstanceController struct {
	Repository instanceRepository
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
// @Failure 400 {object} httputil.ErrorResponse
// @Failure 404 {object} httputil.ErrorResponse
// @Failure 500 {object} httputil.ErrorResponse
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

	log.Error("Websockets 4 ALL")

	primaryNode := instance.PrimaryNode
	port := 11430 // TODO: get port from instance

	err = websocket.Handler(c.Writer, c.Request, primaryNode, port)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}
	c.JSON(200, struct{ Name string }{Name: "lala"})
}
