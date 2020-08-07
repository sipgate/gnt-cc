package handlers

import (
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/model"
	"gnt-cc/rapi"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllJobs godoc
// @Summary Find all jobs
// @Description ...
// @Produce json
// @Success 200 {object} model.AllJobsResponse
// @Failure 404 {object} httputil.HTTPError404
// @Router /clusters/{cluster}/jobs [get]
func GetAllJobs(c *gin.Context) {
	clusterConfig, clusterErr := config.GetClusterConfig(c.Param("cluster"))

	if clusterErr != nil {
		c.AbortWithError(404, clusterErr)
		return
	}

	content, err := rapi.Get(clusterConfig, "/2/jobs?bulk=1")
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var jobsData rapi.JobsBulk
	json.Unmarshal([]byte(content), &jobsData)

	gntJobs := make([]model.GntJob, len(jobsData))
	for _, job := range jobsData {
		gntJobs = append(gntJobs, model.GntJob{
			ID:     job.ID,
			Status: job.Status,
		})
	}

	var jobsCount model.JobStatusCount

	for _, job := range jobsData {
		switch job.Status {
		case "canceled":
			jobsCount.Canceled++
		case "error":
			jobsCount.Error++
		case "pending":
			jobsCount.Pending++
		case "queued":
			jobsCount.Queued++
		case "success":
			jobsCount.Success++
		case "waiting":
			jobsCount.Waiting++
		}
	}
	c.JSON(200, model.AllJobsResponse{
		Cluster:              clusterConfig.Name,
		NumberOfJobs:         len(jobsData),
		NumberOfJobsByStatus: jobsCount,
		Jobs:                 gntJobs,
	})

}

// GetJob godoc
// @Summary Find single job
// @Description ...
// @Produce json
// @Success 200 {object} model.JobResponse
// @Failure 404 {object} httputil.HTTPError404
// @Router /clusters/{cluster}/job/{jobId} [get]
func GetJob(c *gin.Context) {
	clusterConfig, clusterErr := config.GetClusterConfig(c.Param("cluster"))

	if clusterErr != nil {
		c.AbortWithError(404, clusterErr)
		return
	}

	jobID := c.Param("jobId")

	waitForFinish := false
	q := c.Request.URL.Query()
	blockParam, found := q["block"]
	if found {
		blockVal, err := strconv.ParseBool(blockParam[0])
		if err == nil && blockVal {
			waitForFinish = true
		}
	}

	content, err := rapi.Get(clusterConfig, "/2/jobs/"+jobID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	var jobData rapi.Job
	json.Unmarshal([]byte(content), &jobData)

	if waitForFinish {
		sanityCounter := 0

		for range time.Tick(time.Second) {
			if sanityCounter > 300 {
				c.AbortWithError(500, fmt.Errorf("timed out waiting for jobId %s to finish/fail (300 seconds)", jobID))
				return
			} else if jobData.Status != "success" && jobData.Status != "finished" {
				content, err := rapi.Get(clusterConfig, "/2/jobs/"+jobID)
				if err != nil {
					c.AbortWithError(500, err)
					return
				}
				json.Unmarshal([]byte(content), &jobData)
				sanityCounter++
			} else {
				break
			}
		}
	}

	c.JSON(200, gin.H{
		"cluster": clusterConfig.Name,
		"job":     jobData,
	})
}
