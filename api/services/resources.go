package services

import (
	"fmt"
	"gnt-cc/model"
	"sort"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type ResourcesService struct {
	ClusterRepository  clusterRepository
	InstanceRepository instanceRepository
	NodeRepository     nodeRepository
}

type clusterResources struct {
	instances   []string
	nodes       []string
	clusterName string
}

func (s *ResourcesService) CollectAll() CollectResults {
	clusterNames := s.ClusterRepository.GetAllNames()
	channel := make(chan clusterResources)
	var wg sync.WaitGroup

	results := CollectResults{
		Nodes:     []model.ClusterResource{},
		Instances: []model.ClusterResource{},
		Clusters:  []model.Resource{},
	}

	for _, cluster := range clusterNames {
		wg.Add(1)
		go s.asyncCollectFromCluster(cluster, channel, &wg)

		results.Clusters = append(results.Clusters, model.Resource{
			Name: cluster,
		})
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for r := range channel {
		for _, name := range r.instances {
			results.Instances = append(results.Instances, model.ClusterResource{
				Name:        name,
				ClusterName: r.clusterName,
			})
		}

		for _, name := range r.nodes {
			results.Nodes = append(results.Nodes, model.ClusterResource{
				Name:        name,
				ClusterName: r.clusterName,
			})
		}
	}

	sortResourcesAlphabeticallyInPlace(results.Instances)
	sortResourcesAlphabeticallyInPlace(results.Nodes)

	return results
}

func sortResourcesAlphabeticallyInPlace(resources []model.ClusterResource) {
	sort.Slice(resources, func(i int, j int) bool {
		if resources[i].Name == resources[j].Name {
			return resources[i].ClusterName < resources[j].ClusterName
		}
		return resources[i].Name < resources[j].Name
	})
}

func (s *ResourcesService) asyncCollectFromCluster(clusterName string, c chan clusterResources, wg *sync.WaitGroup) {
	defer wg.Done()

	results, err := s.collectFromCluster(clusterName)

	if err != nil {
		log.Error(err)
	} else {
		c <- results
	}
}

func (s *ResourcesService) collectFromCluster(clusterName string) (clusterResources, error) {
	instances, durationInstances, err := collectResourcesFromCluster(clusterName, s.InstanceRepository.GetAllNames)

	if err != nil {
		return clusterResources{}, fmt.Errorf("cannot collect instances from cluster %s: %e", clusterName, err)
	}

	nodes, durationNodes, err := collectResourcesFromCluster(clusterName, s.NodeRepository.GetAllNames)

	if err != nil {
		return clusterResources{}, fmt.Errorf("cannot collect nodes from cluster %s: %e", clusterName, err)
	}

	log.Debugf("collected resources from cluster %s: Instances %s, Nodes: %s", clusterName, durationInstances, durationNodes)

	return clusterResources{
		instances:   instances,
		nodes:       nodes,
		clusterName: clusterName,
	}, nil

}

func collectResourcesFromCluster(clusterName string, method func(string) ([]string, error)) ([]string, time.Duration, error) {
	t1 := time.Now()

	names, err := method(clusterName)

	if err != nil {
		return []string{}, time.Since(t1), err
	}

	return names, time.Since(t1), nil
}
