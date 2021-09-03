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

func TestGroupRepoGetAllFuncReturnsError_WhenRAPIClientReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{}, errors.New("expected error"))
	repo := repository.GroupRepository{RAPIClient: client}
	_, err := repo.GetAll("test")

	assert.EqualError(t, err, "expected error")
}

func TestGroupRepoGetAllFuncReturnsError_WhenJSONResponseIsInvalid(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: "{"}, nil)
	repo := repository.GroupRepository{RAPIClient: client}
	_, err := repo.GetAll("test")

	assert.NotNil(t, err)
}

func TestGroupRepoGetFuncReturnsGroup(t *testing.T) {
	validResponse, _ := ioutil.ReadFile("../testfiles/rapi_responses/valid_groups_response.json")
	client := mocking.NewRAPIClient()
	client.On("Get", mock.Anything, mock.Anything).
		Once().Return(rapi_client.Response{Status: 200, Body: string(validResponse)}, nil)
	repo := repository.GroupRepository{RAPIClient: client}
	result, err := repo.GetAll("test")

	assert.NoError(t, err)
	assert.EqualValues(t, []model.GntGroup{
		{
			Name: "groupname1",
			UUID: "uuid1",
		},
		{
			Name: "groupname2",
			UUID: "uuid2",
		},
	}, result)
}
