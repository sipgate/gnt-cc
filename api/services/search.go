package services

import (
	"gnt-cc/config"
	"gnt-cc/model"
	"strings"
)

type SearchService struct {
	InstanceRepository instanceRepository
}

//TODO: Add tests
func (service *SearchService) Search(query string) (*model.SearchResults, error) {
	c := config.Get()
	filteredInstances := []model.SearchResult{}
	filteredNodes := []model.SearchResult{}
	filteredClusters := []struct {
		Name string
	}{}

	for _, cluster := range c.Clusters {
		// instances
		//TODO: make this resilient against failing clusters
		results, err := service.InstanceRepository.GetAllNames(cluster.Name)
		if err != nil {
			return nil, err
		}
		filteredInstances = append(filteredInstances, filterSearchResults(query, results, cluster.Name)...)

		// nodes
		//TODO: add node search

		// clusters
		if stringContainsIgnoreCase(cluster.Name, query) {
			filteredClusters = append(filteredClusters, struct {
				Name string
			}{
				cluster.Name,
			})
		}
	}
	return &model.SearchResults{
		Nodes:     filteredNodes,
		Instances: filteredInstances,
		Clusters:  filteredClusters,
	}, nil
}

func filterSearchResults(filter string, list []string, clusterName string) []model.SearchResult {
	var filteredList []model.SearchResult
	for _, str := range list {
		if stringContainsIgnoreCase(str, filter) {
			filteredList = append(filteredList, model.SearchResult{
				ClusterName: clusterName,
				Name:        str,
			})
		}
	}
	return filteredList
}

func stringContainsIgnoreCase(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}
