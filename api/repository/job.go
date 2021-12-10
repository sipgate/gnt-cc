package repository

import (
	"encoding/json"
	"fmt"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
	"strings"
)

type JobRepository struct {
	RAPIClient rapi_client.Client
}

func (repo *JobRepository) GetAll(clusterName string) ([]model.GntJob, error) {
	response, err := repo.RAPIClient.Get(clusterName, "/2/jobs?bulk=1")

	if err != nil {
		return nil, err
	}

	var jobData []rapiJobResponse
	err = json.Unmarshal([]byte(response.Body), &jobData)

	if err != nil {
		return nil, err
	}

	jobs := make([]model.GntJob, len(jobData))

	for i := range jobData {
		job := &jobData[i]
		timestamps := parseJobTimestamps(job)

		jobs[i] = model.GntJob{
			ID:         job.ID,
			Summary:    parseJobSummary(job.Summary),
			ReceivedAt: timestamps.receivedAt,
			StartedAt:  timestamps.startedAt,
			EndedAt:    timestamps.endedAt,
			Status:     job.Status,
		}
	}

	return jobs, nil
}

func (repo *JobRepository) Get(clusterName, jobID string) (model.JobResult, error) {
	slug := fmt.Sprintf("/2/jobs/%s", jobID)
	response, err := repo.RAPIClient.Get(clusterName, slug)

	if err != nil {
		return model.JobResult{}, err
	}

	if response.Status == 404 {
		return model.JobResult{
			Found: false,
		}, nil
	}

	var job rapiJobResponse
	err = json.Unmarshal([]byte(response.Body), &job)

	if err != nil {
		return model.JobResult{}, err
	}

	timestamps := parseJobTimestamps(&job)
	jobLog, err := parseJobLog(&job)

	if err != nil {
		return model.JobResult{}, err
	}

	return model.JobResult{
		Found: true,
		Job: model.GntJob{
			ID:         job.ID,
			Summary:    parseJobSummary(job.Summary),
			ReceivedAt: timestamps.receivedAt,
			StartedAt:  timestamps.startedAt,
			EndedAt:    timestamps.endedAt,
			Status:     job.Status,
			Log:        jobLog,
		},
	}, nil
}

type timestamps struct {
	startedAt  int
	endedAt    int
	receivedAt int
}

func parseJobSummary(summary []string) string {
	if len(summary) == 0 {
		return "no summary provided"
	}

	return summary[0]
}

func parseJobTimestamps(job *rapiJobResponse) timestamps {
	startedAt := -1
	endedAt := -1
	receivedAt := -1

	if len(job.StartTs) > 0 {
		startedAt = job.StartTs[0]
	}

	if len(job.EndTs) > 0 {
		endedAt = job.EndTs[0]
	}

	if len(job.ReceivedTs) > 0 {
		receivedAt = job.ReceivedTs[0]
	}

	return timestamps{
		startedAt:  startedAt,
		endedAt:    endedAt,
		receivedAt: receivedAt,
	}
}

func parseJobLog(job *rapiJobResponse) ([]model.GntJobLogEntry, error) {
	if len(job.OpLog) == 0 {
		return []model.GntJobLogEntry{}, nil
	}

	returnedEntries := job.OpLog[0]
	opLogEntries := make([]model.GntJobLogEntry, len(returnedEntries))

	for i, logEntry := range returnedEntries {
		serial, ok := logEntry[0].(float64)
		if !ok {
			return []model.GntJobLogEntry{}, makeOpLogParseError(job.ID, fmt.Sprintf("serial not a float64, but a %T", logEntry[0]))
		}

		timings, ok := logEntry[1].([]interface{})
		if !ok {
			return []model.GntJobLogEntry{}, makeOpLogParseError(job.ID, fmt.Sprintf("timings not an array, but a %T", logEntry[1]))
		}

		timingsStart, ok := timings[0].(float64)
		if !ok {
			return []model.GntJobLogEntry{}, makeOpLogParseError(job.ID, fmt.Sprintf("timingsStart not a float64, but a %T", timings[0]))
		}

		msg, ok := logEntry[3].(string)
		if !ok {
			return []model.GntJobLogEntry{}, makeOpLogParseError(job.ID, fmt.Sprintf("message not a string, but a %T", logEntry[3]))
		}

		opLogEntries[i] = model.GntJobLogEntry{
			Serial:    int(serial),
			StartedAt: int(timingsStart),
			Message:   strings.TrimPrefix(msg, "* "),
		}
	}

	return opLogEntries, nil
}

func makeOpLogParseError(jobID int, reason string) error {
	return fmt.Errorf("cannot parse oplog of job [%d]: %s", jobID, reason)
}
