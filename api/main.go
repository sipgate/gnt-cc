package main

import (
	"gnt-cc/config"
	"gnt-cc/router"
	"net/http"
	"strconv"

	_ "gnt-cc/docs"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	config.Init()

	if !config.Get().DevelopmentMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	router.Routes(r, config.Get().DevelopmentMode)

	bindInfo := config.Get().Bind + ":" + strconv.Itoa(config.Get().Port)
	log.Infof("Starting HTTP server on %s", bindInfo)
	if err := http.ListenAndServe(bindInfo, r); err != nil {
		log.Fatal(err)
	}
}
