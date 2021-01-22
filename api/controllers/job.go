package controllers

import (
	"github.com/gin-gonic/gin"
	"gnt-cc/model"
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
