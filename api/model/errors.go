package model

type BaseError interface {
	error
}

type ClusterNotFoundError interface {
	BaseError
	ClusterName() string
}

type NodeNotFoundError interface {
	BaseError
	NodeName() string
	ClusterName() string
}

type InstanceNotFoundError interface {
	BaseError
	InstanceName() string
	ClusterName() string
}
