package router

import (
	"crypto/tls"
	rice "github.com/GeertJohan/go.rice"
	auth2 "gnt-cc/auth"
	"gnt-cc/config"
	"gnt-cc/controllers"
	"gnt-cc/middleware"
	"gnt-cc/query"
	"gnt-cc/rapi_client"
	"gnt-cc/repository"
	"html/template"
	"net"
	"net/http"
	"strings"
	"time"

	_ "gnt-cc/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func InitTemplates(r *gin.Engine, box *rice.Box) {
	var err error
	var tmpl string
	var message *template.Template

	if tmpl, err = box.String("index.html"); err != nil {
		panic(err)
	}

	if message, err = template.New("index.html").Parse(tmpl); err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(message)
}

func APIRoutes(r *gin.Engine, staticBox *rice.Box, developmentMode bool) {
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
	statisticsController := controllers.StatisticsController{
		InstanceRepository: &instanceRepository,
		NodeRepository:     &nodeRepository,
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if developmentMode {
		r.Use(createCORSConfig("http://localhost:8080"))
	}

	authMiddleware := auth2.GetMiddleware()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
	}

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
		withCluster.GET("/instances/:instance/console", instanceController.OpenInstanceConsole)
		withCluster.GET("/statistics", statisticsController.Get)
	}

	r.StaticFS("/static", staticBox.HTTPBox())
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api") {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
