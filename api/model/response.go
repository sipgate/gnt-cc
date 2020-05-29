package model

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
}

type CreateInstanceResponse struct {
	Cluster  string `json:"cluster"`
	Instance string `json:"instance"`
	Status   string `json:"status"`
	JobId    string `json:"jobId"`
}

type AllJobsResponse struct {
	Cluster              string              `json:"cluster"`
	NumberOfJobsByStatus int                 `json:"number_of_jobs_by_status"`
}

type JobResponse struct {
	Cluster string   `json:"cluster"`
}

type AllNodesResponse struct {
	Cluster       string         `json:"cluster"`
	NumberOfNodes int            `json:"number_of_nodes"`
}
