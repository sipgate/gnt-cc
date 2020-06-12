package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gnt-cc/httputil"
	"gnt-cc/model"
	"gnt-cc/rapi"
	"gnt-cc/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// FindAllJobs godoc
// @Summary Find all jobs
// @Description ...
// @Produce json
// @Success 200 {object} model.AllJobsResponse
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/jobs [get]
func FindAllJobs(c *gin.Context) {
	name := c.Param("cluster")
	if !utils.IsValidCluster(name) {
		httputil.NewError(c, 404, errors.New("cluster not found"))
	} else {
		content, err := rapi.Get(name, "/2/jobs?bulk=1")
		if err != nil {
			httputil.NewError(c, 502, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
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
			Cluster:              name,
			NumberOfJobs:         len(jobsData),
			NumberOfJobsByStatus: jobsCount,
			Jobs:                 gntJobs,
		})
	}
}

// FindJob godoc
// @Summary Find single job
// @Description ...
// @Produce json
// @Success 200 {object} model.JobResponse
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/job/{jobId} [get]
func FindJob(c *gin.Context) {
	name := c.Param("cluster")
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

	if !utils.IsValidCluster(name) {
		httputil.NewError(c, 404, errors.New("cluster not found"))
	} else {
		content, err := rapi.Get(name, "/2/jobs/"+jobID)
		if err != nil {
			httputil.NewError(c, 502, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
		var jobData rapi.Job
		json.Unmarshal([]byte(content), &jobData)
		if waitForFinish {
			sanityCounter := 0
			for range time.Tick(time.Second) {
				if sanityCounter > 300 {
					httputil.NewError(c, 502, errors.New(fmt.Sprintf("Timed out waiting for jobId %s to finish/fail (300 seconds).", jobID)))
					return
				} else if jobData.Status != "success" && jobData.Status != "finished" {
					content, err := rapi.Get(name, "/2/jobs/"+jobID)
					if err != nil {
						httputil.NewError(c, 502, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
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
			"cluster": name,
			"job":     jobData,
		})
	}
}
