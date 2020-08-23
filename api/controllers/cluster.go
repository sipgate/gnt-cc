package controllers

import (
	"github.com/gin-gonic/gin"
	"gnt-cc/config"
	"gnt-cc/model"
)

type ClusterController struct{}

// GetAll godoc
// @Summary Get all clusters
// @Description Get all clusters configured in the config file. Sensitive values will be omitted
// @Produce json
// @Success 200 {object} model.AllClustersResponse
// @Router /clusters [get]
func (controller *ClusterController) GetAll(c *gin.Context) {
	configClusters := config.Get().Clusters
	clusters := make([]model.GntCluster, len(configClusters))

	for i, cluster := range configClusters {
		clusters[i] = model.GntCluster{
			Name:        cluster.Name,
			Hostname:    cluster.Hostname,
			Description: cluster.Description,
			Port:        cluster.Port,
		}
	}

	c.JSON(200, model.AllClustersResponse{
		Clusters: clusters,
	})
}
