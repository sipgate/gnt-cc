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
		c.Error(err)
		return
	}

	c.JSON(200, model.AllInstancesResponse{
		Cluster:           clusterName,
		NumberOfInstances: len(instances),
		Instances:         instances,
	})
}

// Start godoc
// @Summary Start an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobIDResponse
// @Router /clusters/{cluster}/instances/{instance}/start [post]
func (controller *InstanceController) Start(c *gin.Context) {
	controller.SimpleAction(c, "startup")
}

// Restart godoc
// @Summary Restart an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobIDResponse
// @Router /clusters/{cluster}/instances/{instance}/restart [post]
func (controller *InstanceController) Restart(c *gin.Context) {
	controller.SimpleAction(c, "reboot")
}

// Shutdown godoc
// @Summary Shutdown an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobIDResponse
// @Router /clusters/{cluster}/instances/{instance}/shutdown [post]
func (controller *InstanceController) Shutdown(c *gin.Context) {
	controller.SimpleAction(c, "shutdown")
}

// Migrate godoc
// @Summary Migrate an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobIDResponse
// @Router /clusters/{cluster}/instances/{instance}/migrate [post]
func (controller *InstanceController) Migrate(c *gin.Context) {
	controller.SimpleAction(c, "migrate")
}

// Failover godoc
// @Summary Failover an instance in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobIDResponse
// @Router /clusters/{cluster}/instances/{instance}/failover [post]
func (controller *InstanceController) Failover(c *gin.Context) {
	controller.SimpleAction(c, "failover")
}

func (controller *InstanceController) SimpleAction(c *gin.Context, action string) {
	clusterName := c.Param("cluster")
	instanceName := c.Param("instance")

	jobID, err := controller.Actions.PerformSimpleInstanceAction(clusterName, instanceName, action)

	if err != nil {
		c.Error(err)
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
		c.AbortWithStatusJSON(400, model.ErrorResponse{Message: "instance name is required"})
		return
	}

	result, err := controller.Repository.Get(clusterName, instanceName)

	if err != nil {
		c.Error(err)
		return
	}

	if !result.Found {
		c.AbortWithStatusJSON(404, model.ErrorResponse{Message: fmt.Sprintf(MsgInstanceNotFound, instanceName, clusterName)})
		return
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
		c.AbortWithStatusJSON(400, model.ErrorResponse{Message: "instance name is required"})
		return
	}

	result, err := controller.Repository.Get(clusterName, instanceName)

	if err != nil {
		c.Error(err)
		return
	}

	if !result.Found {
		c.AbortWithStatusJSON(404, model.ErrorResponse{Message: fmt.Sprintf(MsgInstanceNotFound, instanceName, clusterName)})
	}

	instance := result.Instance

	if !instance.IsRunning {
		c.AbortWithStatusJSON(400, model.ErrorResponse{Message: "instance is not running"})
		return
	}

	primaryNode := instance.PrimaryNode
	port := result.NetworkPort

	err = websocket.PassThrough(c.Writer, c.Request, primaryNode, port)

	if err != nil {
		c.Error(err)
		return
	}
}
