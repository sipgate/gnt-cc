package model

type InstanceResult struct {
	Found       bool
	NetworkPort int
	Instance    GntInstance
}

type NodeResult struct {
	Found bool
	Node  GntNode
}