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
	url, _ := getRapiConnection(config.ClusterConfig{
		Username: "test",
		SSL:      true,
		Password: "supersecret",
		Hostname: "test-cluster.example.com",
		Port:     5080,
	})
	assert.Equal(t, url, "https://test:supersecret@test-cluster.example.com:5080")

	url, _ = getRapiConnection(config.ClusterConfig{
		Username: "test",
		SSL:      false,
		Password: "supersecret",
		Hostname: "test-cluster.example.com",
		Port:     5090,
	})
	assert.Equal(t, url, "http://test:supersecret@test-cluster.example.com:5090")
}
