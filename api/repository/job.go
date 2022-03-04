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
		timestamps := parseJobTimestamp(job)

		jobs[i] = model.GntJob{
			ID:         job.ID,
			Summary:    job.Summary[0],
			ReceivedAt: job.ReceivedTs[0],
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

	timestamps := parseJobTimestamp(&job)
	jobLog, err := parseToJobLog(job.OpLog)

	if err != nil {
		return model.JobResult{}, makeOpLogParseError(job.ID, err)
	}

	return model.JobResult{
		Found: true,
		Job: model.GntJob{
			ID:         job.ID,
			Summary:    job.Summary[0],
			ReceivedAt: job.ReceivedTs[0],
			StartedAt:  timestamps.startedAt,
			EndedAt:    timestamps.endedAt,
			Status:     job.Status,
			Log:        jobLog,
		},
	}, nil
}

type timestamps struct {
	startedAt int
	endedAt   int
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
		endedAt:   endedAt,
	}
}

func parseToJobLog(log [][]json.RawMessage) (*[]model.GntJobLogEntry, error) {
	opLogEntries := make([]model.GntJobLogEntry, len(log[0]))
	for i, entryRaw := range log[0] {
		entry, err := parseToJobLogEntry(entryRaw)

		if err != nil {
			return nil, err
		}

		opLogEntries[i] = entry
	}

	return &opLogEntries, nil
}

func parseToJobLogEntry(raw json.RawMessage) (model.GntJobLogEntry, error) {
	var entry rapiOpLogEntry
	err := entry.parse(raw)

	if err != nil {
		return model.GntJobLogEntry{}, err
	}

	var message string
	if entry.PayloadType == "message" {
		message, err = parseOpLogMessage(entry.Payload)
	} else if entry.PayloadType == "remote-import" {
		message, err = parseOpLogRemoteImport(entry.Payload)
	} else {
		message = fmt.Sprintf("unknown oplog payload type %s", entry.PayloadType)
	}

	if err != nil {
		return model.GntJobLogEntry{}, err
	}

	return model.GntJobLogEntry{
		Serial:    entry.Serial,
		StartedAt: entry.Timestamps[0],
		Message:   strings.TrimPrefix(message, "* "),
	}, nil
}

func parseOpLogMessage(payload json.RawMessage) (string, error) {
	var message string
	err := json.Unmarshal(payload, &message)

	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(message, "* "), nil
}

func parseOpLogRemoteImport(payload json.RawMessage) (string, error) {
	var parsed rapiRemoteImportPayload
	err := json.Unmarshal(payload, &parsed)

	if err != nil {
		return "", err
	}

	var disks []string
	for _, v := range parsed.Disks {
		var disk rapiRemoteImportDisk
		err = disk.parse(v)

		if err != nil {
			return "", err
		}

		disks = append(disks, fmt.Sprintf("%s:%d", disk.IpAddress, disk.Port))
	}

	return fmt.Sprintf("Importing Disks from: %v", disks), nil
}

func makeOpLogParseError(jobID int, err error) error {
	return fmt.Errorf("cannot parse oplog of job [%d]: %e", jobID, err)
}
