package model

import "gnt-cc/api/rapi"

type ClusterResponse struct {
	Cluster GntCluster `json:"cluster"`
}

type AllClustersResponse struct {
	Clusters []GntCluster `json:"clusters"`
}

type AllInstancesResponse struct {
	Cluster           string        `json:"cluster"`
	NumberOfInstances int           `json:"number_of_instances"`
	Instances         []GntInstance `json:"instances"`
}

type InstanceResponse struct {
	Cluster  string        `json:"cluster"`
	Instance rapi.Instance `json:"instance"`
}

type CreateInstanceResponse struct {
	Cluster  string `json:"cluster"`
	Instance string `json:"instance"`
	Status   string `json:"status"`
	JobId    string `json:"jobId"`
}

type AllJobsResponse struct {
	Cluster              string              `json:"cluster"`
	NumberOfJobs         rapi.JobStatusCount `json:"number_of_jobs"`
	NumberOfJobsByStatus int                 `json:"number_of_jobs_by_status"`
	Jobs                 rapi.JobsBulk       `json:"jobs"`
}

type JobResponse struct {
	Cluster string   `json:"cluster"`
	Job     rapi.Job `json:"job"`
}

type AllNodesResponse struct {
	Cluster       string         `json:"cluster"`
	NumberOfNodes int            `json:"number_of_nodes"`
	Nodes         rapi.NodesBulk `json:"nodes"`
}
