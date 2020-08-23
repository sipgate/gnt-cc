package router

import (
	"crypto/tls"
	auth2 "gnt-cc/auth"
	"gnt-cc/config"
	"gnt-cc/controllers"
	"gnt-cc/middleware"
	"gnt-cc/query"
	"gnt-cc/rapi_client"
	"gnt-cc/repository"
	"net"
	"net/http"
	"time"

	_ "gnt-cc/docs"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func createHTTPClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
}

func createRAPIClientFromConfig(configs []config.ClusterConfig) (rapi_client.Client, error) {
	return rapi_client.New(createHTTPClient(), configs)
}

func createCORSConfig(url string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*", url},
		AllowCredentials: true,
		AllowWebSockets:  true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	})
}

func Routes(r *gin.Engine, developmentMode bool) {
	rapiClient, err := createRAPIClientFromConfig(config.Get().Clusters)
	if err != nil {
		panic(err)
	}

	instanceRepository := repository.InstanceRepository{RAPIClient: rapiClient, QueryPerformer: &query.Performer{}}
	nodeRepository := repository.NodeRepository{RAPIClient: rapiClient}

	clusterController := controllers.ClusterController{}
	instanceController := controllers.InstanceController{
		Repository: &instanceRepository,
	}
	nodeController := controllers.NodeController{
		Repository:         &nodeRepository,
		InstanceRepository: &instanceRepository,
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if developmentMode {
		r.Use(createCORSConfig("http://localhost:8080"))
	}

	authMiddleware := auth2.GetMiddleware()

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Warningf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	v1.POST("/login", authMiddleware.LoginHandler)

	auth := v1.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/clusters", clusterController.GetAll)
	}

	withCluster := auth.Group("/clusters/:cluster")
	withCluster.Use(middleware.RequireCluster())
	{
		withCluster.GET("/nodes", nodeController.GetAll)
		withCluster.GET("/nodes/:node", nodeController.Get)
		withCluster.GET("/instances", instanceController.GetAll)
		withCluster.GET("/instances/:instance", instanceController.Get)
	}
}
