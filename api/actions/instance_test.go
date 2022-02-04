package actions_test

import (
	"errors"
	"gnt-cc/actions"
	"gnt-cc/mocking"
	"gnt-cc/rapi_client"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rapiActions = []string{
	"startup",
	"reboot",
	"shutdown",
}

func TestInstanceMethodReturnsError_WhenRAPIReturnsError(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Post", "testClusterName", mock.Anything, nil).
		Return(rapi_client.Response{}, errors.New("expected error"))
	client.On("Put", "testClusterName", mock.Anything, nil).
		Return(rapi_client.Response{}, errors.New("expected error"))

	actions := actions.InstanceActions{RAPIClient: client}

	for _, rapiAction := range rapiActions {
		_, err := actions.PerformSimpleInstanceAction("testClusterName", "testInstanceName", rapiAction)
		assert.EqualError(t, err, "expected error")
	}
}

func TestRAPIEndpointIsCalled_WhenInvokingInstanceMethod(t *testing.T) {
	client := mocking.NewRAPIClient()
	client.On("Put", "testClusterName", "/2/instances/testInstanceName/startup", nil).
		Once().Return(rapi_client.Response{
		Body: "423458",
	}, nil)
	client.On("Post", "testClusterName", "/2/instances/testInstanceName/reboot", nil).
		Once().Return(rapi_client.Response{
		Body: "423458",
	}, nil)
	client.On("Put", "testClusterName", "/2/instances/testInstanceName/shutdown", nil).
		Once().Return(rapi_client.Response{
		Body: "423458",
	}, nil)

	actions := actions.InstanceActions{RAPIClient: client}

	for _, rapiAction := range rapiActions {
		jobId, err := actions.PerformSimpleInstanceAction("testClusterName", "testInstanceName", rapiAction)
		assert.Nil(t, err)
		assert.Equal(t, jobId, 423458)
	}
}
