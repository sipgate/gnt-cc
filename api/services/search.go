package services

import (
	log "github.com/sirupsen/logrus"
	"gnt-cc/config"
	"gnt-cc/model"
	"strings"
)

type SearchService struct {
	InstanceRepository instanceRepository
	NodeRepository     nodeRepository
}

//TODO: Add tests
func (service *SearchService) Search(query string) (*model.SearchResults, error) {
	c := config.Get()
	filteredInstances := []model.ResourceSearchResult{}
	filteredNodes := []model.ResourceSearchResult{}
	filteredClusters := []model.ClusterSearchResult{}

	for _, cluster := range c.Clusters {
		// instances
		results, err := service.InstanceRepository.GetAllNames(cluster.Name)
		if err != nil {
			log.Errorf("search service: error talking to cluster '%s': %e", cluster.Name, err)
		} else {
			filteredInstances = append(filteredInstances, filterSearchResults(query, results, cluster.Name)...)
		}

		// nodes
		results, err = service.NodeRepository.GetAllNames(cluster.Name)
		if err != nil {
			log.Errorf("search service: error talking to cluster '%s': %e", cluster.Name, err)
		} else {
			filteredNodes = append(filteredNodes, filterSearchResults(query, results, cluster.Name)...)
		}

		// clusters
		if stringContainsIgnoreCase(cluster.Name, query) {
			filteredClusters = append(filteredClusters, model.ClusterSearchResult{
				Name: cluster.Name,
			})
		}
	}
	return &model.SearchResults{
		Nodes:     filteredNodes,
		Instances: filteredInstances,
		Clusters:  filteredClusters,
	}, nil
}

func filterSearchResults(filter string, list []string, clusterName string) []model.ResourceSearchResult {
	var filteredList []model.ResourceSearchResult
	for _, str := range list {
		if stringContainsIgnoreCase(str, filter) {
			filteredList = append(filteredList, model.ResourceSearchResult{
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
