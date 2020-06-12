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
	Cluster  string      `json:"cluster"`
	Instance GntInstance `json:"instance"`
}

type CreateInstanceResponse struct {
	Cluster  string `json:"cluster"`
	Instance string `json:"instance"`
	Status   string `json:"status"`
	JobId    int    `json:"jobId"`
}

type AllJobsResponse struct {
	Cluster              string         `json:"cluster"`
	NumberOfJobs         int            `json:"number_ob_jobs"`
	NumberOfJobsByStatus JobStatusCount `json:"number_of_jobs_by_status"`
	Jobs                 []GntJob       `json:"jobs"`
}

type JobResponse struct {
	Cluster string `json:"cluster"`
}

type AllNodesResponse struct {
	Cluster       string    `json:"cluster"`
	NumberOfNodes int       `json:"number_of_nodes"`
	Nodes         []GntNode `json:"nodes"`
}

type NodeResponse struct {
	Cluster string  `json:"cluster"`
	Node    GntNode `json:"node"`
}
