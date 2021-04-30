package controllers

import (
	"fmt"
	"gnt-cc/model"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	Repository jobRepository
}

// GetAll godoc
// @Summary Get all jobs in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.AllJobsResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/jobs [get]
func (controller *JobController) GetAll(c *gin.Context) {
	clusterName := c.Param("cluster")

	jobs, err := controller.Repository.GetAll(clusterName)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	c.JSON(200, model.AllJobsResponse{
		Cluster:      clusterName,
		NumberOfJobs: len(jobs),
		Jobs:         jobs,
	})
}

// Get godoc
// @Summary Get a single job in a given cluster
// @Description ...
// @Produce json
// @Success 200 {object} model.JobResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/jobs [get]
func (controller *JobController) Get(c *gin.Context) {
	clusterName := c.Param("cluster")
	jobID := c.Param("job")

	
	if jobID == "" {
		c.AbortWithStatusJSON(400, createErrorBody("job ID is required"))
		return
	}
	
	result, err := controller.Repository.Get(clusterName, jobID)

	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	if !result.Found {
		c.AbortWithStatusJSON(404, createErrorBody(
			fmt.Sprintf(MsgJobNotFound, jobID, clusterName)))
	}

	c.JSON(200, model.JobResponse{
		Cluster:  clusterName,
		Job: result.Job,
	})
}
