package main

import (
	auth2 "gnt-cc/auth"
	"gnt-cc/config"
	"gnt-cc/handlers"
	"net/http"
	"strconv"
	"time"

	_ "gnt-cc/docs"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO
// @title GNT-CC
// @version 0.0
// @description An API wrapper with local/ldap authentication around one or more Ganeti RAPI backends

// @contact.name API Support
// @contact.url https://github.com/sipgate/gnt-cc/issues

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	config.Parse()

	if !config.Get().DevelopmentMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if config.Get().DevelopmentMode {
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
		auth.GET("/clusters", handlers.FindAllClusters)
		auth.GET("/clusters/:cluster", handlers.FindCluster)
		auth.GET("/clusters/:cluster/nodes", handlers.FindAllNodes)
		auth.GET("/clusters/:cluster/instances", handlers.FindAllInstances)
		auth.GET("/clusters/:cluster/instances/:instance", handlers.FindInstance)
		auth.POST("/clusters/:cluster/instance", handlers.CreateInstance)
		auth.POST("/clusters/:cluster/instances", handlers.CreateMultipleInstancesHandler)
		auth.GET("/clusters/:cluster/console/:instance", handlers.OpenInstanceConsole)
		auth.GET("/clusters/:cluster/jobs", handlers.FindAllJobs)
		auth.GET("/clusters/:cluster/job/:jobId", handlers.FindJob)
	}

	bindInfo := config.Get().Bind + ":" + strconv.Itoa(config.Get().Port)
	log.Infof("Starting HTTP server on %s", bindInfo)
	if err := http.ListenAndServe(bindInfo, r); err != nil {
		log.Fatal(err)
	}

}
