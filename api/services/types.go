package services

import "gnt-cc/model"

type (
	clusterRepository interface {
		GetAllNames() []string
	}

	instanceRepository interface {
		GetAllNames(clusterName string) ([]string, error)
	}

	nodeRepository interface {
		GetAllNames(clusterName string) ([]string, error)
	}

	CollectResults struct {
		Instances []model.ClusterResource
		Nodes     []model.ClusterResource
		Clusters  []model.Resource
	}

	resourcesService interface {
		CollectAll() CollectResults
	}
)
