package services_test

import (
	"github.com/stretchr/testify/assert"
	"gnt-cc/mocking"
	"gnt-cc/model"
	"gnt-cc/services"
	"testing"
)

type resourcesServiceMock struct {
}

func (*resourcesServiceMock) CollectAll() services.CollectResults {
	return services.CollectResults{
		Nodes: []model.ClusterResource{
			{
				ClusterName: "Cluster01",
				Name:        "cluster01-node01",
			},
			{
				ClusterName: "Cluster01",
				Name:        "cluster01-node02",
			},
			{
				ClusterName: "cluster02",
				Name:        "cluster02-node01",
			},
		},
		Instances: []model.ClusterResource{
			{
				ClusterName: "Cluster01",
				Name:        "test-instance01",
			},
			{
				ClusterName: "cluster02",
				Name:        "test-instance02",
			},
			{
				ClusterName: "cluster02",
				Name:        "test-instance03",
			},
		},
		Clusters: []model.Resource{
			{
				Name: "Cluster01",
			},
			{
				Name: "cluster02",
			},
		},
	}
}

func TestSearchService_Search(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  model.SearchResults
	}{
		{
			name:  "No Search Results",
			query: "non-existing-data",
			want: model.SearchResults{
				Nodes:     []model.ClusterResource{},
				Instances: []model.ClusterResource{},
				Clusters:  []model.Resource{},
			},
		},
		{
			name:  "All instances containing 'test'",
			query: "test",
			want: model.SearchResults{
				Nodes: []model.ClusterResource{},
				Instances: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "test-instance01",
					},
					{
						ClusterName: "cluster02",
						Name:        "test-instance02",
					},
					{
						ClusterName: "cluster02",
						Name:        "test-instance03",
					},
				},
				Clusters: []model.Resource{},
			},
		},
		{
			name:  "All resources containing 'Cluster0'",
			query: "Cluster0",
			want: model.SearchResults{
				Nodes: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "cluster01-node01",
					},
					{
						ClusterName: "Cluster01",
						Name:        "cluster01-node02",
					},
					{
						ClusterName: "cluster02",
						Name:        "cluster02-node01",
					},
				},
				Instances: []model.ClusterResource{},
				Clusters: []model.Resource{
					{
						Name: "Cluster01",
					},
					{
						Name: "cluster02",
					},
				},
			},
		},
		{
			name:  "All resources containing 'a'",
			query: "a",
			want: model.SearchResults{
				Nodes: []model.ClusterResource{},
				Instances: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "test-instance01",
					},
					{
						ClusterName: "cluster02",
						Name:        "test-instance02",
					},
					{
						ClusterName: "cluster02",
						Name:        "test-instance03",
					},
				},
				Clusters: []model.Resource{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchService := services.SearchService{ResourcesService: &resourcesServiceMock{}}
			assert.Equalf(t, tt.want, searchService.Search(tt.query), "Search(%v)", tt.query)
		})
	}
}

func TestResourcesService_CollectAll(t *testing.T) {
	tests := []struct {
		name      string
		clusters  []string
		instances []string
		nodes     []string
		want      services.CollectResults
	}{
		{
			name:      "Return empty lists",
			clusters:  []string{},
			instances: []string{},
			nodes:     []string{},
			want: services.CollectResults{
				Instances: []model.ClusterResource{},
				Nodes:     []model.ClusterResource{},
				Clusters:  []model.Resource{},
			},
		},
		{
			name:      "Return one cluster with two nodes and three instances",
			clusters:  []string{"Cluster01"},
			instances: []string{"test-instance01", "test-instance02", "test-instance03"},
			nodes:     []string{"node01", "node02"},
			want: services.CollectResults{
				Instances: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "test-instance01",
					},
					{
						ClusterName: "Cluster01",
						Name:        "test-instance02",
					},
					{
						ClusterName: "Cluster01",
						Name:        "test-instance03",
					},
				},
				Nodes: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "node01",
					},
					{
						ClusterName: "Cluster01",
						Name:        "node02",
					},
				},
				Clusters: []model.Resource{
					{"Cluster01"},
				},
			},
		},
		{
			name:      "Return three clusters with two nodes and two instances each",
			clusters:  []string{"Cluster01", "Cluster02"},
			instances: []string{"test-instance01", "test-instance02"},
			nodes:     []string{"node01", "node02"},
			want: services.CollectResults{
				Instances: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "test-instance01",
					},
					{
						ClusterName: "Cluster02",
						Name:        "test-instance01",
					},
					{
						ClusterName: "Cluster01",
						Name:        "test-instance02",
					},
					{
						ClusterName: "Cluster02",
						Name:        "test-instance02",
					},
				},
				Nodes: []model.ClusterResource{
					{
						ClusterName: "Cluster01",
						Name:        "node01",
					},
					{
						ClusterName: "Cluster02",
						Name:        "node01",
					},
					{
						ClusterName: "Cluster01",
						Name:        "node02",
					},
					{
						ClusterName: "Cluster02",
						Name:        "node02",
					},
				},
				Clusters: []model.Resource{
					{"Cluster01"},
					{"Cluster02"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clusterRepository := mocking.NewClusterRepository()
			instanceRepository := mocking.NewInstanceRepository()
			nodeRepository := mocking.NewNodeRepository()

			clusterRepository.On("GetAllNames").Return(tt.clusters)
			for _, cluster := range tt.clusters {
				instanceRepository.On("GetAllNames", cluster).Return(tt.instances, nil)
				nodeRepository.On("GetAllNames", cluster).Return(tt.nodes, nil)
			}
			s := &services.ResourcesService{
				ClusterRepository:  clusterRepository,
				InstanceRepository: instanceRepository,
				NodeRepository:     nodeRepository,
			}
			assert.Equalf(t, tt.want, s.CollectAll(), "CollectAll()")
		})
	}
}
