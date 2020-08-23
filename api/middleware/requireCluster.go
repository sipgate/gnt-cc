package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gnt-cc/config"
	"gnt-cc/model"
)

func RequireCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		clusterName := c.Param("cluster")

		if clusterName == "" {
			c.AbortWithStatusJSON(400, model.ErrorResponse{Message: "cluster is required"})
			return
		}

		if !config.ClusterExists(clusterName) {
			c.AbortWithStatusJSON(404, model.ErrorResponse{
				Message: fmt.Sprintf("cluster cannot be found: %s", clusterName),
			})
			return
		}

		c.Next()
	}
}
