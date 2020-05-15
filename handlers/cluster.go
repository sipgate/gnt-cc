package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gnt-cc/config"
	"gnt-cc/httputil"
	"gnt-cc/utils"
)

// FindAllClusters godoc
// @Summary Get all cluster names
// @Description get all clusters (currently only returns names of clusters)
// @Produce  json
// @Success 200 {object} model.AllClustersResponse
// @Router /clusters [get]
func FindAllClusters(context *gin.Context) {
	var names []string
	for _, cluster := range config.Get().Clusters {
		names = append(names, cluster.Name)
	}
	context.JSON(200, gin.H{
		"clusters": names,
	})
}

// FindCluster godoc
// @Summary Get cluster details
// @Description get details of a cluster with a given name (Currently only returns its name)
// @Produce  json
// @Success 200 {object} model.ClusterResponse
// @Router /clusters/{cluster} [get]
func FindCluster(c *gin.Context) {
	name := c.Param("cluster")
	if !utils.IsValidCluster(name) {
		httputil.NewError(c, 404, errors.New("cluster not found"))
	} else {
		// TODO: return actual information, not just name
		c.JSON(200, gin.H{
			"cluster": name,
		})
	}
}
