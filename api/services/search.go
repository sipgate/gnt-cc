package services

import (
	log "github.com/sirupsen/logrus"
	"gnt-cc/config"
	"gnt-cc/model"
	"strings"
	"time"
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
		instanceStart := time.Now()
		results, err := service.InstanceRepository.GetAllNames(cluster.Name)
		if err != nil {
			log.Errorf("search service: error talking to cluster '%s': %e", cluster.Name, err)
		} else {
			filteredInstances = append(filteredInstances, filterSearchResults(query, results, cluster.Name)...)
		}
		instanceElapsed := time.Since(instanceStart)

		// nodes
		nodeStart := time.Now()
		results, err = service.NodeRepository.GetAllNames(cluster.Name)
		if err != nil {
			log.Errorf("search service: error talking to cluster '%s': %e", cluster.Name, err)
		} else {
			filteredNodes = append(filteredNodes, filterSearchResults(query, results, cluster.Name)...)
		}
		nodeElapsed := time.Since(nodeStart)

		// clusters
		clusterStart := time.Now()
		if stringContainsIgnoreCase(cluster.Name, query) {
			filteredClusters = append(filteredClusters, model.ClusterSearchResult{
				Name: cluster.Name,
			})
		}
		clusterElapsed := time.Since(clusterStart)
		log.Debugf("search service stats for '%s': Instances: %s, Nodes: %s, Clusters: %s", cluster.Name, instanceElapsed, nodeElapsed, clusterElapsed)
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
