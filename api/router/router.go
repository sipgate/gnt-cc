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
	"html/template"
	"net"
	"net/http"
	"strings"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	engine               *gin.Engine
	clusterController    controllers.ClusterController
	instanceController   controllers.InstanceController
	statisticsController controllers.StatisticsController
	nodeController       controllers.NodeController
	jobController        controllers.JobController
}

func New(engine *gin.Engine) *router {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Get().PublicUrl}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	engine.Use(cors.New(corsConfig))

	rapiClient, err := createRAPIClientFromConfig(config.Get().Clusters, config.Get().RapiConfig)
	if err != nil {
		panic(err)
	}

	instanceRepository := repository.InstanceRepository{RAPIClient: rapiClient, QueryPerformer: &query.Performer{}}
	groupRepository := repository.GroupRepository{RAPIClient: rapiClient}
	nodeRepository := repository.NodeRepository{RAPIClient: rapiClient, GroupRepository: groupRepository}
	jobRepository := repository.JobRepository{RAPIClient: rapiClient}

	r := router{
		engine: engine,
	}

	r.clusterController = controllers.ClusterController{}
	r.instanceController = controllers.InstanceController{
		Repository: &instanceRepository,
	}
	r.nodeController = controllers.NodeController{
		Repository:         &nodeRepository,
		InstanceRepository: &instanceRepository,
	}
	r.statisticsController = controllers.StatisticsController{
		InstanceRepository: &instanceRepository,
		NodeRepository:     &nodeRepository,
	}
	r.jobController = controllers.JobController{
		Repository: &jobRepository,
	}

	return &r
}

func createHTTPClient(skipCertificateVerify bool) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertificateVerify},
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

func createRAPIClientFromConfig(configs []config.ClusterConfig, rapiConfig config.RapiConfig) (rapi_client.Client, error) {
	return rapi_client.New(createHTTPClient(rapiConfig.SkipCertificateVerify), configs)
}

func (r *router) InitTemplates(box *rice.Box) {
	var err error
	var tmpl string
	var message *template.Template

	if tmpl, err = box.String("index.html"); err != nil {
		panic(err)
	}

	if message, err = template.New("index.html").Parse(tmpl); err != nil {
		panic(err)
	}

	r.engine.SetHTMLTemplate(message)
}

func (r *router) SetupAPIRoutes() {
	authMiddleware := auth2.GetMiddleware()

	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.engine.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
	}

	auth := v1.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/clusters", r.clusterController.GetAll)
	}

	withCluster := auth.Group("/clusters/:cluster")
	withCluster.Use(middleware.RequireCluster())
	{
		withCluster.GET("/nodes", r.nodeController.GetAll)
		withCluster.GET("/nodes/:node", r.nodeController.Get)
		withCluster.GET("/instances", r.instanceController.GetAll)
		withCluster.GET("/instances/:instance", r.instanceController.Get)
		withCluster.GET("/instances/:instance/console", r.instanceController.OpenInstanceConsole)
		withCluster.GET("/statistics", r.statisticsController.Get)
		withCluster.GET("/jobs", r.jobController.GetAll)
		withCluster.GET("/jobs/many", r.jobController.GetManyWithLogs)
		withCluster.GET("/jobs/:job", r.jobController.Get)
	}

	r.engine.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api") {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
