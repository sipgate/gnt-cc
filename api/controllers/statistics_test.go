package controllers_test

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gnt-cc/controllers"
	"gnt-cc/mocking"
	"gnt-cc/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type (
	nodeReturnValues struct {
		nodes []model.GntNode
		err   error
	}

	instanceReturnValues struct {
		instances []model.GntInstance
		err       error
	}
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
	tests := []struct {
		name string

		instanceRepositoryReturns instanceReturnValues
		nodeRepositoryReturns     nodeReturnValues

		expectedReturnCode int
		expectedData       interface{}
	}{
		{
			name: "Returns statistics and 200 OK",

			instanceRepositoryReturns: instanceReturnValues{
				instances: validInstances,
			},
			nodeRepositoryReturns: nodeReturnValues{
				nodes: validNodes,
			},

			expectedReturnCode: http.StatusOK,
			expectedData: model.StatisticsResponse{
				Instances: model.StatisticsElement{
					Count:       2,
					MemoryTotal: 1024,
					CPUCount:    4,
				},
				Nodes: model.StatisticsElement{
					Count:       2,
					MemoryTotal: 2048,
					CPUCount:    8,
				},
			},
		},
		{
			name: "Expect 500 when instance repository returns error",

			instanceRepositoryReturns: instanceReturnValues{
				err: errors.New("error"),
			},
			nodeRepositoryReturns: nodeReturnValues{},

			expectedReturnCode: http.StatusInternalServerError,
			expectedData:       model.ErrorResponse{Message: "internal server error"},
		},
		{
			name: "Expect 500 when node repository returns error",

			instanceRepositoryReturns: instanceReturnValues{},
			nodeRepositoryReturns: nodeReturnValues{
				err: errors.New("error"),
			},

			expectedReturnCode: http.StatusInternalServerError,
			expectedData:       model.ErrorResponse{Message: "internal server error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			instanceRepository := mocking.NewInstanceRepository()
			nodeRepository := mocking.NewNodeRepository()

			instanceRepository.On("GetAll", mock.Anything).Return(tt.instanceRepositoryReturns.instances, tt.instanceRepositoryReturns.err)
			nodeRepository.On("GetAll", mock.Anything).Return(tt.nodeRepositoryReturns.nodes, tt.nodeRepositoryReturns.err)

			controller := &controllers.StatisticsController{
				InstanceRepository: instanceRepository,
				NodeRepository:     nodeRepository,
			}

			controller.Get(c)

			expectedJson, err := json.Marshal(tt.expectedData)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedReturnCode, rr.Code)
			assert.JSONEq(t, string(expectedJson), rr.Body.String())
		})
	}
}
