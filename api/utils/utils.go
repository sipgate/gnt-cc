package utils

import (
	"gnt-cc/config"
)

func IsValidCluster(clusterName string) bool {
	for _, cluster := range []config.GanetiCluster(config.Get().Clusters) {
		if cluster.Name == clusterName {
			return true
		}
	}
	return false
}
