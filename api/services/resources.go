package services

import (
	"fmt"
	"gnt-cc/config"
	"gnt-cc/model"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type ResourcesService struct {
	InstanceRepository instanceRepository
	NodeRepository     nodeRepository
}

type clusterResources struct {
	instances   []string
	nodes       []string
	clusterName string
}

func (s *ResourcesService) CollectAll() CollectResults {
	c := config.Get()
	channel := make(chan clusterResources)
	var wg sync.WaitGroup

	results := CollectResults{
		Nodes:     []model.ClusterResource{},
		Instances: []model.ClusterResource{},
		Clusters:  []model.Resource{},
	}

	for _, cluster := range c.Clusters {
		wg.Add(1)
		go s.asyncCollectFromCluster(cluster.Name, channel, &wg)

		results.Clusters = append(results.Clusters, model.Resource{
			Name: cluster.Name,
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

	return results
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
		instances: instances,
		nodes:     nodes,
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
