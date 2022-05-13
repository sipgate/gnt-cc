package repository_test

import (
	"errors"
	"gnt-cc/mocking"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
	"gnt-cc/repository"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNodeRepoGetAllFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected error"))
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.GetAll("test")

	assert.EqualError(t, err, "expected error")
}

func TestNodeRepoGetAllNamesFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected error"))
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.GetAllNames("test")

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

func TestNodeRepoGetAllNamesFuncReturnsError_WhenJSONResponseIsInvalid(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: "{"}, nil)
	repo := repository.NodeRepository{RAPIClient: client}
	_, err := repo.GetAllNames("test")

	assert.NotNil(t, err)
}

func TestNodeRepoGetAllFuncReturnsNodes(t *testing.T) {
	validNodesResponse, _ := ioutil.ReadFile("../testfiles/rapi_responses/valid_nodes_response.json")
	validGroupsResponse, _ := ioutil.ReadFile("../testfiles/rapi_responses/valid_groups_response.json")
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, "/2/nodes?bulk=1").
		Once().Return(rapi_client.Response{Status: 200, Body: string(validNodesResponse)}, nil)
	client.On("Get", mock.Anything, "/2/groups?bulk=1").
		Once().Return(rapi_client.Response{Status: 200, Body: string(validGroupsResponse)}, nil)
	groupRepo := repository.GroupRepository{RAPIClient: client}
	repo := repository.NodeRepository{RAPIClient: client, GroupRepository: groupRepo}
	nodes, err := repo.GetAll("test")

	assert.NoError(t, err)
	assert.EqualValues(t, []model.GntNode{{
		Name:                    "node1",
		MemoryTotal:             64420,
		MemoryFree:              31873,
		CPUCount:                24,
		DiskTotal:               1660012,
		DiskFree:                671740,
		IsMasterCandidate:       false,
		IsMasterCapable:         true,
		IsMaster:                false,
		IsVMCapable:             true,
		PrimaryInstancesCount:   2,
		SecondaryInstancesCount: 0,
		GroupName:               "groupname1",
	}, {
		Name:                    "node2",
		MemoryTotal:             128848,
		MemoryFree:              82412,
		CPUCount:                40,
		DiskTotal:               2241324,
		DiskFree:                1119916,
		IsMasterCandidate:       true,
		IsMasterCapable:         true,
		IsVMCapable:             true,
		PrimaryInstancesCount:   0,
		SecondaryInstancesCount: 2,
		GroupName:               "groupname2",
	}}, nodes)
}

func TestNodeRepoGetAllNamesFuncReturnsNodes(t *testing.T) {
	validNodeNamessResponse, _ := ioutil.ReadFile("../testfiles/rapi_responses/valid_node_names_response.json")
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, "/2/nodes").
		Once().Return(rapi_client.Response{Status: 200, Body: string(validNodeNamessResponse)}, nil)
	groupRepo := repository.GroupRepository{RAPIClient: client}
	repo := repository.NodeRepository{RAPIClient: client, GroupRepository: groupRepo}
	nodes, err := repo.GetAllNames("test")

	assert.NoError(t, err)
	assert.EqualValues(t, []string{
		"homer",
		"marge",
		"bart",
	}, nodes)
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
	assert.EqualValues(t, model.GntNodeWithInstances{
		GntNode: model.GntNode{
			Name:              "node1",
			MemoryTotal:       15709,
			MemoryFree:        10244,
			CPUCount:          8,
			IsMasterCandidate: true,
			IsMasterCapable:   true,
			IsMaster:          true,
			IsVMCapable:       true,
		},
		PrimaryInstances:   []string{"burns", "milhouse"},
		SecondaryInstances: []string{},
	}, result.Node)
}
