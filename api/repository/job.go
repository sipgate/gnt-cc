package repository

import (
	"encoding/json"
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
		startedAt := -1
		endedAt := -1

		if len(job.StartTs) > 0 {
			startedAt = job.StartTs[0]
		}

		if len(job.EndTs) > 0 {
			endedAt = job.EndTs[0]
		}

		jobs[i] = model.GntJob{
			ID:         job.ID,
			Summary:    job.Summary[0],
			ReceivedAt: job.ReceivedTs[0],
			StartedAt:  startedAt,
			EndedAt:    endedAt,
			Status:     job.Status,
		}
	}

	return jobs, nil
}
