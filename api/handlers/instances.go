package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/dummy"
	"gnt-cc/httputil"
	"gnt-cc/model"
	"gnt-cc/rapi"
	"gnt-cc/websocket"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// FindAllInstances godoc
// @Summary Get all instances in a given cluster
// @Description ...
// @Produce  json
// @Success 200 {object} model.AllInstancesResponse
// @Failure 404 {object} httputil.HTTPError
// @Failure 502 {object} httputil.HTTPError
// @Router /clusters/{cluster}/instances [get]
func FindAllInstances(context *gin.Context) {
	clusterName := context.Param("cluster")
	if !config.ClusterExists(clusterName) {
		httputil.NewError(context, 404, errors.New("cluster not found"))
	} else {
		var instances []model.GntInstance

		if config.Get().DummyMode {
			instances = dummy.GetInstances(20)
		} else {
			var err error
			instances, err = rapi.GetInstances(clusterName)

			if err != nil {
				httputil.NewError(context, 500, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
				return
			}
		}

		context.JSON(200, model.AllInstancesResponse{
			Cluster:           clusterName,
			NumberOfInstances: len(instances),
			Instances:         instances,
		})
	}
}

// FindInstance godoc
// @Summary Get an instance in a given cluster with the given name
// @Description ...
// @Produce  json
// @Success 200 {object} model.InstanceResponse
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/instances/{instance} [get]
func FindInstance(context *gin.Context) {
	clusterName := context.Param("cluster")
	instanceName := context.Param("instance")
	if !config.ClusterExists(clusterName) {
		context.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Cluster not found"})
		return
	}

	var instance model.GntInstance

	if config.Get().DummyMode {
		instance = dummy.GetInstance(instanceName)
	} else {
		var err error
		instance, err = rapi.GetInstance(clusterName, instanceName)

		if err != nil {
			httputil.NewError(context, 500, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
	}

	context.JSON(200, model.InstanceResponse{
		Cluster:  clusterName,
		Instance: instance,
	})
}

// OpenInstanceConsole godoc
// @Summary Open a websocket connection to a cluster instance console
// @Description ...
// @Failure 400 {object} httputil.HTTPError400
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/console/{instance} [get]
func OpenInstanceConsole(context *gin.Context) {
	name := context.Param("cluster")
	instanceName := context.Param("instance")
	if !config.ClusterExists(name) {
		httputil.NewError(context, 404, errors.New("cluster not found"))
	} else {
		content, err := rapi.Get(name, "/2/instances/"+instanceName)
		if err != nil {
			log.Errorf("RAPI Backend Error: %s", err)

			httputil.NewError(context, 502, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
		var instanceData rapi.Instance
		err = json.Unmarshal([]byte(content), &instanceData)
		if err != nil {
			log.Errorf("Could not parse JSON result into RapiInstance struct: %s", err)

			httputil.NewError(context, 502, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
		// overwrite host/port for the websocket proxy "backend" for debugging:
		//instanceData.Pnode = "127.0.0.1"
		//instanceData.NetworkPort = 3333
		if instanceData.OperState && instanceData.NetworkPort > 0 {
			err = websocket.Handler(context.Writer, context.Request, instanceData.Pnode, instanceData.NetworkPort)
			return
		}
		log.Infof("Cannot request console for shutdown instance '%s'", instanceData.Name)

		httputil.NewError(context, 400, errors.New(fmt.Sprintf("Cannot request console for shutdown instance '%s'", instanceData.Name)))
		return
	}
}

// CreateInstance godoc
// @Summary Create an instance on a given cluster
// @Description ...
// @Accept json
// @Produce json
// @param instance body model.CreateInstanceRequest true "..."
// @Success 200 {object} model.CreateInstanceResponse
// @Failure 400 {object} httputil.HTTPError400
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/instance [post]
func CreateInstance(context *gin.Context) {
	name := context.Param("cluster")
	if !config.ClusterExists(name) {
		httputil.NewError(context, 404, errors.New("cluster not found"))
	} else {
		var json model.CreateInstanceRequest

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := rapi.Get(name, "/2/instances/"+json.Name)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Cannot create instance '%s', it already exists", json.Name)})
			return
		}

		rapiParams := rapi.CreateInstanceParameters{
			InstanceName:      json.Name,
			DiskTemplate:      "", // TODO
			Vcpus:             json.VCpuCores,
			MemoryInMegabytes: json.MemoryMegaBytes,
			Nics:              nil, // TODO
			Disks:             nil, // TODO
		}

		newInstance := rapi.NewGanetiInstance(rapiParams)

		status, err := rapi.Post(name, "/2/instances", newInstance)
		if err != nil {
			log.Errorf("RAPI Backend Error: %s", err)

			context.JSON(http.StatusBadGateway, gin.H{"error": "RAPI Error"})
			return
		}

		gntJobID, err := strconv.Atoi(strings.TrimSpace(status))
		if err != nil {
			log.Errorf("Could not parse/understand RAPI JobID: %s", err)

			context.JSON(http.StatusBadGateway, gin.H{"error": "RAPI Error"})
			return
		}

		context.JSON(200, model.CreateInstanceResponse{
			Cluster:  name,
			Instance: newInstance.InstanceName,
			Status:   "jobSubmitted",
			JobId:    gntJobID,
		})
	}
}

func CreateMultipleInstancesHandler(c *gin.Context) {
	// TODO
}
