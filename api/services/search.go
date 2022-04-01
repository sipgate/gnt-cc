package services

import (
	"gnt-cc/model"
	"sort"
	"strings"
)

type SearchService struct {
	ResourcesService resourcesService
}

const RESULTS_LIMIT = 5

func (s *SearchService) Search(query string) model.SearchResults {
	resources := s.ResourcesService.CollectAll()

	clusters := []model.Resource{}

	for _, cluster := range resources.Clusters {
		if stringContainsIgnoreCase(cluster.Name, query) {
			clusters = append(clusters, model.Resource{
				Name: cluster.Name,
			})
		}
	}

	instances := filterClusterResources(query, resources.Instances)
	nodes := filterClusterResources(query, resources.Nodes)

	sortResourcesAlphabeticallyInPlace(instances)
	sortResourcesAlphabeticallyInPlace(nodes)

	return model.SearchResults{
		Nodes:     nodes[0:min(len(nodes), RESULTS_LIMIT)],
		Instances: instances[0:min(len(instances), RESULTS_LIMIT)],
		Clusters:  clusters[0:min(len(clusters), RESULTS_LIMIT)],
	}
}

func sortResourcesAlphabeticallyInPlace(resources []model.ClusterResource) {
	sort.Slice(resources, func(i int, j int) bool {
		return resources[i].Name < resources[j].Name
	})
}

func filterClusterResources(filter string, list []model.ClusterResource) []model.ClusterResource {
	filtered := []model.ClusterResource{}

	for _, item := range list {
		if stringContainsIgnoreCase(item.Name, filter) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func stringContainsIgnoreCase(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
