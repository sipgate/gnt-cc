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

type jobOperationsPayloadType map[string]func(interface{}) (string, error)

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
	jobLog, err := parseJobLog(&job)

	if err != nil {
		return model.JobResult{}, err
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

func parseJobOpMessage(data interface{}) (string, error) {
	return fmt.Sprintf("%s", data), nil
}

func castRemoteImportStructureToDiskList(data interface{}) ([]interface{}, error) {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid structure received: %v, type: %T", data, data)
	}
	diskList, ok := dataMap["disks"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid structure received: %v, type: %T", dataMap["disks"], dataMap["disks"])
	}
	return diskList, nil
}

func parseJobOpRemoteImport(data interface{}) (string, error) {
	var disks []string

	diskList, err := castRemoteImportStructureToDiskList(data)
	if err != nil {
		return "", err
	}

	for _, diskEntry := range diskList {
		diskEntryCasted, ok := diskEntry.([]interface{})
		if !ok {
			return "", fmt.Errorf("invalid structure received: %v", diskEntry)
		}
		ip, ok := diskEntryCasted[0].(string)
		if !ok {
			return "", fmt.Errorf("invalid remote disk IP: %v", diskEntryCasted[0])
		}
		port, ok := diskEntryCasted[1].(float64)
		if !ok {
			return "", fmt.Errorf("invalid remote disk port: %v", diskEntryCasted[1])
		}
		disks = append(disks, fmt.Sprintf("%s:%.0f", ip, port))
	}
	return fmt.Sprintf("Importing Disks from: %v", disks), nil
}

func parseJobOpUnknownType(data interface{}) (string, error) {
	return fmt.Sprintf("Received unknown data: '%v'", data), nil
}

func parseJobLog(job *rapiJobResponse) ([]model.GntJobLogEntry, error) {
	returnedEntries := job.OpLog[0]
	opLogEntries := make([]model.GntJobLogEntry, len(returnedEntries))
	jobOpTypesMap := jobOperationsPayloadType{
		"message":       parseJobOpMessage,
		"remote-import": parseJobOpRemoteImport,
	}

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

		msgType, ok := logEntry[2].(string)
		if !ok {
			return []model.GntJobLogEntry{}, makeOpLogParseError(job.ID, fmt.Sprintf("message-type not a string, but a %T", logEntry[3]))
		}

		parser, exists := jobOpTypesMap[msgType]
		var err error
		var msg string
		if !exists {
			msg, err = parseJobOpUnknownType(logEntry[3])
		} else {
			msg, err = parser(logEntry[3])
		}
		if err != nil {
			return []model.GntJobLogEntry{}, makeOpLogParseError(job.ID, fmt.Sprintf("failed to parse oplog message payload: %e", err))
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
