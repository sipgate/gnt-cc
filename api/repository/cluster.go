package repository

import "gnt-cc/config"

type ClusterRepository struct {
}

func (*ClusterRepository) GetAllNames() []string {
	c := config.Get()
	clusters := make([]string, len(c.Clusters))
	for i, cluster := range c.Clusters {
		clusters[i] = cluster.Name
	}
	return clusters
}
