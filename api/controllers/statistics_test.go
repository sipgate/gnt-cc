package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gnt-cc/config"
	"gnt-cc/model"
	"gnt-cc/router"
	"gnt-cc/testcluster"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	validInstances = []model.GntInstance{
		{
			CpuCount:    2,
			MemoryTotal: 512,
		},
		{
			CpuCount:    2,
			MemoryTotal: 512,
		},
	}
	validNodes = []model.GntNode{
		{
			MemoryTotal: 1024,
			CPUCount:    4,
		},
		{
			MemoryTotal: 1024,
			CPUCount:    4,
		},
	}
)

func TestStatisticsController_Get(t *testing.T) {
	testCluster, err := testcluster.New()
	if err != nil {
		panic(err)
	}

	defer testCluster.Terminate()

	err = testCluster.CreateInstance("testinstance")
	if err != nil {
		panic(err)
	}
	defer testCluster.RemoveInstance("testinstance")

	err = testCluster.DebugListInstances()
	if err != nil {
		panic(err)
	}

	config.InitCustomConfig(config.Config{
		Bind:                 "127.0.0.1",
		Port:                 8000,
		DevelopmentMode:      false,
		PublicUrl:            "http://localhost:3000",
		JwtSigningKey:        "test",
		JwtExpire:            1 * time.Minute,
		AuthenticationMethod: "builtin",
		Loglevel:             "debug",
		Users: []config.UserConfig{{
			Username: "test",
			Password: "test",
		}},
		Clusters: []config.ClusterConfig{{
			Name:        "testcluster",
			Hostname:    testCluster.GetClusterHost(),
			Port:        5080,
			Description: "local testcluster managed by testcontainers",
			Username:    "gnt-cc",
			Password:    "gnt-cc",
			SSL:         true,
		}},
		LDAPConfig: config.LDAPConfig{},
		RapiConfig: config.RapiConfig{
			SkipCertificateVerify: true,
		},
	})

	engine := gin.New()
	r := router.New(engine)
	r.SetupAPIRoutes()

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: engine,
	}

	defer srv.Close()

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	loginRes, err := http.Post("http://localhost:8000/api/v1/login", "application/json", bytes.NewBuffer([]byte(`{
		"username": "test",
		"password": "test"
	}`)))
	if err != nil {
		fmt.Printf("error making login request: %s\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", "http://localhost:8000/api/v1/clusters/testcluster/statistics", nil)
	for _, cookie := range loginRes.Cookies() {
		req.AddCookie(cookie)
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	var responseJSON map[string]map[string]json.RawMessage
	all, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &responseJSON)
	if err != nil {
		return
	}

	assert.Equal(t, 9, responseJSON["instances"]["count"])

	log.Printf("%s\n", all)

	//tests := []struct {
	//	name string
	//
	//	instanceRepositoryReturns instanceReturnValues
	//	nodeRepositoryReturns     nodeReturnValues
	//
	//	expectedReturnCode int
	//	expectedData       interface{}
	//}{
	//	{
	//		name: "Returns statistics and 200 OK",
	//
	//		instanceRepositoryReturns: instanceReturnValues{
	//			instances: validInstances,
	//		},
	//		nodeRepositoryReturns: nodeReturnValues{
	//			nodes: validNodes,
	//		},
	//
	//		expectedReturnCode: http.StatusOK,
	//		expectedData: model.StatisticsResponse{
	//			Instances: model.StatisticsElement{
	//				Count:       2,
	//				MemoryTotal: 1024,
	//				CPUCount:    4,
	//			},
	//			Nodes: model.StatisticsElement{
	//				Count:       2,
	//				MemoryTotal: 2048,
	//				CPUCount:    8,
	//			},
	//		},
	//	},
	//	{
	//		name: "Expect 500 when instance repository returns error",
	//
	//		instanceRepositoryReturns: instanceReturnValues{
	//			err: errors.New("error"),
	//		},
	//		nodeRepositoryReturns: nodeReturnValues{},
	//
	//		expectedReturnCode: http.StatusInternalServerError,
	//		expectedData:       model.ErrorResponse{Message: "internal server error"},
	//	},
	//	{
	//		name: "Expect 500 when node repository returns error",
	//
	//		instanceRepositoryReturns: instanceReturnValues{},
	//		nodeRepositoryReturns: nodeReturnValues{
	//			err: errors.New("error"),
	//		},
	//
	//		expectedReturnCode: http.StatusInternalServerError,
	//		expectedData:       model.ErrorResponse{Message: "internal server error"},
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		rr := httptest.NewRecorder()
	//		cookie, _ := gin.CreateTestContext(rr)
	//
	//		instanceRepository := mocking.NewInstanceRepository()
	//		nodeRepository := mocking.NewNodeRepository()
	//
	//		instanceRepository.On("GetAll", mock.Anything).Return(tt.instanceRepositoryReturns.instances, tt.instanceRepositoryReturns.err)
	//		nodeRepository.On("GetAll", mock.Anything).Return(tt.nodeRepositoryReturns.nodes, tt.nodeRepositoryReturns.err)
	//
	//		controller := &controllers.StatisticsController{
	//			InstanceRepository: instanceRepository,
	//			NodeRepository:     nodeRepository,
	//		}
	//
	//		controller.Get(cookie)
	//
	//		expectedJson, err := json.Marshal(tt.expectedData)
	//
	//		assert.NoError(t, err)
	//		assert.Equal(t, tt.expectedReturnCode, rr.Code)
	//		assert.JSONEq(t, string(expectedJson), rr.Body.String())
	//	})
	//}
}
