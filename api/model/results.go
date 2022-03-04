package model

type InstanceResult struct {
	Found       bool
	NetworkPort int
	Instance    GntInstance
}

type NodeResult struct {
	Found bool
	Node  GntNodeWithInstances
}

type GroupResult struct {
	Found bool
	Group GntGroup
}

type JobResult struct {
	Found bool
	Job   GntJob
}

type SearchResult struct {
	ClusterName string
	Name        string
}

type SearchResults struct {
	Nodes     []SearchResult
	Instances []SearchResult
	Clusters  []struct {
		Name string
	}
}
