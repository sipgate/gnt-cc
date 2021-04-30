package repository

import (
	"encoding/json"
	"fmt"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
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

	for i, job := range jobData {
		ts := parseJobTimestamp(&job)

		jobs[i] = model.GntJob{
			ID:         job.ID,
			Summary:    job.Summary[0],
			ReceivedAt: job.ReceivedTs[0],
			StartedAt:  ts.startedAt,
			EndedAt:    ts.endedAt,
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

	ts := parseJobTimestamp(&job)
	log, err := parseJobLog(&job)

	if err != nil {
		return model.JobResult{}, err
	}

	return model.JobResult{
		Found: true,
		Job: model.GntJob{
			ID:         job.ID,
			Summary:    job.Summary[0],
			ReceivedAt: job.ReceivedTs[0],
			StartedAt:  ts.startedAt,
			EndedAt:    ts.endedAt,
			Status:     job.Status,
			Log: log,
		},
	}, nil
}

type timestamps struct {
		startedAt int
		endedAt int
}

func parseJobTimestamp(job *rapiJobResponse) timestamps {
	startedAt := -1
	endedAt := -1

	if len(job.StartTs) > 0 {
		startedAt = job.StartTs[0]
	}

	if len(job.EndTs) > 0 {
		endedAt = job.EndTs[0]
	}

	return timestamps{
		startedAt: startedAt,
		endedAt: endedAt,
	}
}

type dataFloat64 struct {
	data float64
}

type dataString struct {
	data string
}

type dataTimestamps struct {
	data []dataFloat64
}

type dataInterface struct {
	data []interface{}
}

func parseJobLog(job *rapiJobResponse) ([]model.GntJobLogEntry, error) {
	entries, ok := job.OpLog[0].(dataInterface)

	if !ok {
		return []model.GntJobLogEntry{}, fmt.Errorf("cannot parse job oplog [jobID=%d]", job.ID)
	}

	opLogEntries := make([]model.GntJobLogEntry, len(entries.data))

	for i, v := range entries.data {
		entry, ok := v.(dataInterface)

		if !ok {
			return []model.GntJobLogEntry{}, fmt.Errorf("cannot parse job oplog, cannot parse entry [jobID=%d]", job.ID)
		}

		logEntry := entry.data

		serial, ok := logEntry[0].(dataFloat64)
		if !ok {
			return []model.GntJobLogEntry{}, fmt.Errorf("cannot parse job oplog, serial could not be parsed [jobID=%d]", job.ID)
		}

		ts, ok := logEntry[1].(dataTimestamps)
		if !ok {
			return []model.GntJobLogEntry{}, fmt.Errorf("cannot parse job oplog, ts could not be parsed [jobID=%d]", job.ID)
		}

		msg, ok := logEntry[3].(dataString)
		if !ok {
			return []model.GntJobLogEntry{}, fmt.Errorf("cannot parse job oplog, message could not be parsed [jobID=%d]", job.ID)
		}

		opLogEntries[i] = model.GntJobLogEntry{
			Serial: int(serial.data),
			StartedAt: int(ts.data[0].data),
			EndedAt: int(ts.data[1].data),
			Message: msg.data,
		}
	}

	return opLogEntries, nil
}

