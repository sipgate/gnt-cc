package repository_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gnt-cc/mocking"
	"gnt-cc/model"
	"gnt-cc/query"
	"gnt-cc/rapi_client"
	"gnt-cc/repository"
	"io/ioutil"
	"testing"
)

func TestInstanceRepoGetAllFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	queryPerformer := mocking.NewQueryPerformer()
	queryPerformer.On("Perform", mock.Anything, mock.Anything).
		Once().Return([]query.Resource{}, errors.New("expected error"))
	repo := repository.InstanceRepository{QueryPerformer: queryPerformer}
	_, err := repo.GetAll("test")

	assert.EqualError(t, err, "expected error")
}

func TestInstanceRepoGetFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected error"))
	repo := repository.InstanceRepository{RAPIClient: client}
	_, err := repo.Get("test", "instance")

	assert.EqualError(t, err, "expected error")
}

func TestInstanceRepoGetFuncResultFoundFieldIsFalse_WhenRAPIClientReturns404(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 404, Body: ""}, nil)
	repo := repository.InstanceRepository{RAPIClient: client}
	result, err := repo.Get("test", "instance")

	assert.NoError(t, err)
	assert.False(t, result.Found)
}

func TestInstanceRepoGetAllFuncReturnsError_OnInvalidResourceReturned(t *testing.T) {
	queryPerformer := mocking.NewQueryPerformer()
	queryPerformer.On("Perform", mock.Anything, mock.Anything, mock.Anything).
		Once().Return([]query.Resource{{"Name": func() {}}}, nil)
	repo := repository.InstanceRepository{QueryPerformer: queryPerformer}
	_, err := repo.GetAll("test")

	assert.Error(t, err)
}

func TestInstanceRepoGetFuncReturnsError_WhenJSONResponseIsInvalid(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: "{"}, nil)
	repo := repository.InstanceRepository{RAPIClient: client}
	_, err := repo.Get("test", "instance")

	assert.NotNil(t, err)
}

func TestInstanceRepoGetFuncReturnsSuccessfulResult_OnValidResponse(t *testing.T) {
	validResponse, err := ioutil.ReadFile("../testfiles/rapi_responses/valid_instance_response.json")
	assert.NoError(t, err)

	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: string(validResponse)}, nil)
	repo := repository.InstanceRepository{RAPIClient: client}
	result, err := repo.Get("test", "instance")

	assert.NoError(t, err)
	assert.True(t, result.Found)
	assert.EqualValues(t, model.GntInstance{
		Name:           "burns",
		PrimaryNode:    "node1",
		SecondaryNodes: []string{"node4"},
		CpuCount:       1,
		MemoryTotal:    128,
		Nics:           []model.GntNic{},
		Disks:          []model.GntDisk{},
		Tags:           []string{"tag1"},
	}, result.Instance)
}

func TestInstanceRepoGetAllFuncCorrectlyReturnsInstances_WhenAValidRAPIResponseWasReceived(t *testing.T) {
	partialBeParams := struct {
		VCPUs  int
		MaxMem int
	}{
		VCPUs:  2,
		MaxMem: 1024,
	}

	queryPerformer := mocking.NewQueryPerformer()
	queryPerformer.On("Perform", mock.Anything, query.RequestConfig{
		ClusterName:  "test1",
		ResourceType: "instance",
		Fields:       []string{"name", "pnode", "snodes", "beparams", "oper_state", "hvparams"},
	}).
		Once().Return([]query.Resource{{
		"name":       "bart",
		"pnode":      "node4",
		"snodes":     []string{"node1", "node2"},
		"beparams":   partialBeParams,
		"oper_state": false,
		"hvparams": map[string]interface{}{
			"vnc_bind_address": "",
		},
	}, {
		"name":       "smithers",
		"pnode":      "node2",
		"snodes":     []string{"node4"},
		"beparams":   partialBeParams,
		"oper_state": true,
		"hvparams": map[string]interface{}{
			"vnc_bind_address": "test",
		},
	}}, nil)

	repo := repository.InstanceRepository{QueryPerformer: queryPerformer}
	allInstances, err := repo.GetAll("test1")

	assert.Nil(t, err)
	assert.EqualValues(t, []model.GntInstance{
		{
			Name:           "bart",
			PrimaryNode:    "node4",
			SecondaryNodes: []string{"node1", "node2"},
			MemoryTotal:    1024,
			CpuCount:       2,
			OffersVNC:      false,
			IsRunning:      false,
		},
		{
			Name:           "smithers",
			PrimaryNode:    "node2",
			SecondaryNodes: []string{"node4"},
			MemoryTotal:    1024,
			CpuCount:       2,
			OffersVNC:      true,
			IsRunning:      true,
		},
	}, allInstances)
}
