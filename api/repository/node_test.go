package repository_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gnt-cc/mocking"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
	"gnt-cc/repository"
	"io/ioutil"
	"testing"
)

func TestNodeRepoGetAllFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected error"))
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.GetAll("test")

	assert.EqualError(t, err, "expected error")
}

func TestNodeRepoGetFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected error"))
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.Get("test", "node")

	assert.EqualError(t, err, "expected error")
}

func TestNodeRepoGetFuncReturnsNotFoundResult_WhenRAPIClientReturns404(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 404, Body: ""}, nil)
	repo := repository.NodeRepository{RAPIClient: client}
	result, err := repo.Get("test", "node")

	assert.NoError(t, err)
	assert.False(t, result.Found)
}

func TestNodeRepoGetFuncReturnsError_WhenJSONResponseIsInvalid(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: "{"}, nil)
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.Get("test", "node")

	assert.NotNil(t, err)
}

func TestNodeRepoGetAllFuncReturnsError_WhenJSONResponseIsInvalid(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: "{"}, nil)
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.GetAll("test")

	assert.NotNil(t, err)
}

func TestNodeRepoGetAllFuncReturnsNodes(t *testing.T) {
	validResponse, _ := ioutil.ReadFile("../testfiles/rapi_responses/valid_nodes_response.json")
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: string(validResponse)}, nil)
	repo := repository.NodeRepository{RAPIClient: client}
	nodes, err := repo.GetAll("test")

	assert.NoError(t, err)
	assert.EqualValues(t, []model.GntNode{{
		Name:               "node1",
		MemoryTotal:        15709,
		MemoryFree:         10230,
		CPUCount:           8,
		PrimaryInstances:   nil,
		SecondaryInstances: nil,
		IsMasterCandidate:  true,
		IsMasterCapable:    true,
		IsMaster:           true,
		IsVMCapable:        true,
	}, {
		Name:               "node2",
		MemoryTotal:        15709,
		MemoryFree:         10229,
		CPUCount:           8,
		PrimaryInstances:   nil,
		SecondaryInstances: nil,
		IsMasterCandidate:  true,
		IsMasterCapable:    true,
		IsVMCapable:        true,
	}}, nodes)
}

func TestNodeRepoGetFuncReturnsNode(t *testing.T) {
	validResponse, _ := ioutil.ReadFile("../testfiles/rapi_responses/valid_node_response.json")
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: string(validResponse)}, nil)
	repo := repository.NodeRepository{RAPIClient: client}
	result, err := repo.Get("test", "node")

	assert.NoError(t, err)
	assert.True(t, result.Found)
	assert.EqualValues(t, model.GntNode{
		Name:               "node1",
		MemoryTotal:        15709,
		MemoryFree:         10244,
		CPUCount:           8,
		PrimaryInstances:   []string{"burns", "milhouse"},
		SecondaryInstances: []string{},
		IsMasterCandidate:  true,
		IsMasterCapable:    true,
		IsMaster:           true,
		IsVMCapable:        true,
	}, result.Node)
}
