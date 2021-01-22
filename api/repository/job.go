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
		jobs[i] = model.GntJob{
			ID:         job.ID,
			Summary:    job.Summary[0],
			ReceivedAt: job.ReceivedTs[0],
			StartedAt:  job.StartTs[0],
			EndedAt:    job.EndTs[0],
			Status:     job.Status,
		}
	}

	return jobs, nil
}
