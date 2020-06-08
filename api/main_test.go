package main

import (
	"gnt-cc/config"
	"gnt-cc/rapi"
	"gnt-cc/utils"
	"testing"
)

func TestIsvalidCluster(t *testing.T) {
	config.Parse()
	validClusterName := "production-cluster"
	if !utils.IsValidCluster(validClusterName) {
		t.Errorf("IsValidCluster failed to validate a proper cluster name")
	}

	invalidClusterName := "myInvalidCluster"
	if utils.IsValidCluster(invalidClusterName) {
		t.Errorf("IsValidCluster failed to invalidate a wrong cluster name")
	}
}

func TestNewGanetiInstance(t *testing.T) {
	p := rapi.CreateInstanceParameters{
		InstanceName:      "my-test-instance",
		DiskTemplate:      "plain",
		Vcpus:             4,
		MemoryInMegabytes: 1024,
	}
	instCreateStruct := rapi.NewGanetiInstance(p)
	if instCreateStruct.Version != 1 {
		t.Errorf("instanceCreateObject contains the wrong version (expected: 1, got: %d)", instCreateStruct.Version)
	}

	// add many more checks

}
