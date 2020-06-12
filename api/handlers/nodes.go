package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gnt-cc/config"
	"gnt-cc/dummy"
	"gnt-cc/httputil"
	"gnt-cc/model"
	"gnt-cc/rapi"
	"gnt-cc/utils"
)

// FindAllNodes godoc
// @Summary Find all nodes
// @Description ...
// @Produce json
// @Success 200 {object} model.AllNodesResponse
// @Failure 404 {object} httputil.HTTPError404
// @Failure 502 {object} httputil.HTTPError502
// @Router /clusters/{cluster}/nodes [get]
func FindAllNodes(context *gin.Context) {
	clusterName := context.Param("cluster")
	if !utils.IsValidCluster(clusterName) {
		httputil.NewError(context, 404, errors.New("cluster not found"))
		return
	}

	var nodes []model.GntNode

	if config.Get().DummyMode {
		nodes = dummy.GetNodes(20)
	} else {
		var err error
		nodes, err = rapi.GetNodes(clusterName)

		if err != nil {
			httputil.NewError(context, 500, errors.New(fmt.Sprintf("RAPI Backend Error: %s", err)))
			return
		}
	}

	context.JSON(200, model.AllNodesResponse{
		Cluster:       clusterName,
		NumberOfNodes: len(nodes),
		Nodes:         nodes,
	})
}
