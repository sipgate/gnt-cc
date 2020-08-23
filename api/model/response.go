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
	Cluster            string        `json:"cluster"`
	Node               GntNode       `json:"node"`
	PrimaryInstances   []GntInstance `json:"primaryInstances"`
	SecondaryInstances []GntInstance `json:"secondaryInstances"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
