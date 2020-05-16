package handlers

import (
	"github.com/gin-gonic/gin"
	"gnt-cc/config"
	"gnt-cc/model"
)

// FindAllClusters godoc
// @Summary Get all clusters
// @Description get all clusters
// @Produce  json
// @Success 200 {object} model.AllClustersResponse
// @Router /clusters [get]
func FindAllClusters(context *gin.Context) {
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

	context.JSON(200, gin.H{
		"clusters": clusters,
	})
}

// FindCluster godoc
// @Summary Get cluster details
// @Description get details of a cluster with a given name (Currently only returns its name)
// @Produce  json
// @Success 200 {object} model.ClusterResponse
//// @Router /clusters/{cluster} [get]
//func FindCluster(c *gin.Context) {
//	name := c.Param("cluster")
//	if !utils.IsValidCluster(name) {
//		httputil.NewError(c, 404, errors.New("cluster not found"))
//	} else {
//		c.JSON(200, gin.H{
//			"cluster": ,
//		})
//	}
//}
