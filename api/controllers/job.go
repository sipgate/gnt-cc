package controllers

import (
	"fmt"
	"gnt-cc/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	Repository jobRepository
}

type GetManyWithLogsOptions struct {
	IDs string `form:"ids"`
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
		c.Error(err)
		return
	}

	c.JSON(200, model.AllJobsResponse{
		NumberOfJobs: len(jobs),
		Jobs:         jobs,
	})
}

// GetAll godoc
// @Summary Get specific jobs in a given cluster including their logs
// @Description ...
// @Produce json
// @Success 200 {object} model.AllJobsResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clusters/{cluster}/jobs/many [get]
// @Param ids query []int true "job ids to retrieve"
func (controller *JobController) GetManyWithLogs(c *gin.Context) {
	clusterName := c.Param("cluster")

	var options GetManyWithLogsOptions
	if c.BindQuery(&options) != nil {
		c.AbortWithStatusJSON(400, model.ErrorResponse{Message: "ids parameter is required"})
		return
	}

	jobs := []model.GntJob{}
	for _, id := range strings.Split(options.IDs, ",") {
		job, err := controller.Repository.Get(clusterName, id)

		if err != nil {
			c.Error(err)
			return
		}

		jobs = append(jobs, job.Job)
	}

	c.JSON(200, model.AllJobsResponse{
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
// @Router /clusters/{cluster}/job/{job} [get]
func (controller *JobController) Get(c *gin.Context) {
	clusterName := c.Param("cluster")
	jobID := c.Param("job")

	if jobID == "" {
		c.AbortWithStatusJSON(400, model.ErrorResponse{Message: "job ID is required"})
		return
	}

	result, err := controller.Repository.Get(clusterName, jobID)

	if err != nil {
		c.Error(err)
		return
	}

	if !result.Found {
		c.AbortWithStatusJSON(404, model.ErrorResponse{Message: fmt.Sprintf(MsgJobNotFound, jobID, clusterName)})
		return
	}

	c.JSON(200, model.JobResponse{
		Job: result.Job,
	})
}
