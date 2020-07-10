package router

import (
	auth2 "gnt-cc/auth"
	"gnt-cc/handlers"
	"time"

	_ "gnt-cc/docs"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine, developmentMode bool) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if developmentMode {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*", "http://localhost:8080"},
			AllowCredentials: true,
			AllowWebSockets:  true,
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			MaxAge:           12 * time.Hour,
		}))
	}

	authMiddleware := auth2.GetMiddleware()

	r.POST("/v1/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Warningf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/v1")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/clusters", handlers.GetAllClusters)
		// auth.GET("/clusters/:cluster", handlers.FindCluster)
		auth.GET("/clusters/:cluster/nodes", handlers.GetAllNodes)
		auth.GET("/clusters/:cluster/nodes/:node", handlers.GetNode)
		auth.GET("/clusters/:cluster/instances", handlers.GetAllInstances)
		auth.GET("/clusters/:cluster/instances/:instance", handlers.GetInstance)
		// auth.POST("/clusters/:cluster/instance", handlers.CreateInstance)
		// auth.POST("/clusters/:cluster/instances", handlers.CreateMultipleInstancesHandler)
		// auth.GET("/clusters/:cluster/console/:instance", handlers.OpenInstanceConsole)
		auth.GET("/clusters/:cluster/jobs", handlers.GetAllJobs)
		auth.GET("/clusters/:cluster/job/:jobId", handlers.GetJob)
	}
}
