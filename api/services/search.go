package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gnt-cc/config"
	"gnt-cc/model"
	"strings"
	"sync"
	"time"
)

type SearchService struct {
	InstanceRepository instanceRepository
	NodeRepository     nodeRepository
}

//TODO: Add tests
func (service *SearchService) Search(query string) (model.SearchResults, error) {
	c := config.Get()

	var results model.SearchResults

	channel := make(chan clusterSearchResults)
	var waitGroup sync.WaitGroup
	for _, cluster := range c.Clusters {
		waitGroup.Add(1)
		go service.asyncSearchInCluster(query, cluster.Name, channel, &waitGroup)

		// clusters
		if stringContainsIgnoreCase(cluster.Name, query) {
			results.Clusters = append(results.Clusters, model.ClusterSearchResult{
				Name: cluster.Name,
			})
		}
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for r := range channel {
		for _, entry := range r.instancesNames {
			results.Instances = append(results.Instances, model.ResourceSearchResult{
				ClusterName: r.clusterName,
				Name:        entry,
			})
		}

		for _, entry := range r.nodesNames {
			results.Nodes = append(results.Nodes, model.ResourceSearchResult{
				ClusterName: r.clusterName,
				Name:        entry,
			})
		}
	}

	return results, nil
}

func (service *SearchService) asyncSearchInCluster(query string, clusterName string, channel chan clusterSearchResults, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	result, err := service.searchInCluster(query, clusterName)
	if err != nil {
		log.Error(err)
		return
	}
	channel <- result
}

func (service *SearchService) searchInCluster(query string, clusterName string) (clusterSearchResults, error) {
	results := clusterSearchResults{
		clusterName: clusterName,
	}

	// instances
	instanceStart := time.Now()
	instanceResults, err := service.InstanceRepository.GetAllNames(clusterName)
	if err != nil {
		return clusterSearchResults{}, fmt.Errorf("error talking to cluster '%s': %e", clusterName, err)
	}
	results.instancesNames = filterSearchResults(query, instanceResults)
	instanceElapsed := time.Since(instanceStart)

	// nodes
	nodeStart := time.Now()
	nodeResults, err := service.NodeRepository.GetAllNames(clusterName)
	if err != nil {
		return clusterSearchResults{}, fmt.Errorf("error talking to cluster '%s': %e", clusterName, err)
	}
	results.nodesNames = filterSearchResults(query, nodeResults)
	nodeElapsed := time.Since(nodeStart)

	log.Debugf("search service stats for '%s': Instances: %s, Nodes: %s", clusterName, instanceElapsed, nodeElapsed)
	return results, nil
}

func filterSearchResults(filter string, list []string) []string {
	var filteredList []string
	for _, str := range list {
		if stringContainsIgnoreCase(str, filter) {
			filteredList = append(filteredList, str)
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
