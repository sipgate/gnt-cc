package rapi

import (
	"gnt-cc/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGanetiInstance(t *testing.T) {
	p := CreateInstanceParameters{
		InstanceName:      "my-test-instance",
		DiskTemplate:      "plain",
		Vcpus:             4,
		MemoryInMegabytes: 1024,
	}
	instCreateStruct := NewGanetiInstance(p)

	assert.Equal(t, instCreateStruct.Version, 1)
	assert.Equal(t, instCreateStruct.Hypervisor, "fake")
	assert.Equal(t, instCreateStruct.DiskTemplate, "plain")
	assert.Equal(t, instCreateStruct.InstanceName, "my-test-instance")

	// add many more checks
}

func TestGetRapiConnection(t *testing.T) {
	config.Parse()
	url, _ := getRapiConnection("production-cluster")

	assert.Equal(t, url, "https://gnt-cc:somepassword@prod-cluster.example.com:5080")
}
