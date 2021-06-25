package model

type ClusterResponse struct {
	Cluster GntCluster `json:"cluster"`
}

type AllClustersResponse struct {
	Clusters []GntCluster `json:"clusters"`
}

type AllInstancesResponse struct {
	Cluster           string        `json:"cluster"`
	NumberOfInstances int           `json:"numberOfInstances"`
	Instances         []GntInstance `json:"instances"`
}

type InstanceResponse struct {
	Cluster  string      `json:"cluster"`
	Instance GntInstance `json:"instance"`
}

type AllNodesResponse struct {
	Cluster       string    `json:"cluster"`
	NumberOfNodes int       `json:"numberOfNodes"`
	Nodes         []GntNode `json:"nodes"`
}

type NodeResponse struct {
	Cluster            string               `json:"cluster"`
	Node               GntNodeWithInstances `json:"node"`
	PrimaryInstances   []GntInstance        `json:"primaryInstances"`
	SecondaryInstances []GntInstance        `json:"secondaryInstances"`
}

type AllJobsResponse struct {
	Cluster      string   `json:"cluster"`
	NumberOfJobs int      `json:"numberOfJobs"`
	Jobs         []GntJob `json:"jobs"`
}

type JobResponse struct {
	Cluster string `json:"cluster"`
	Job     GntJob `json:"job"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatisticsElement struct {
	Count       int `json:"count"`
	MemoryTotal int `json:"memoryTotal"`
	CPUCount    int `json:"cpuCount"`
}

type StatisticsResponse struct {
	Instances StatisticsElement `json:"instances"`
	Nodes     StatisticsElement `json:"nodes"`
	Master    string            `json:"master"`
}
