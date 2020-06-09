package rapi

import (
	"gnt-cc/config"
	"testing"
)

func TestNewGanetiInstance(t *testing.T) {
	p := CreateInstanceParameters{
		InstanceName:      "my-test-instance",
		DiskTemplate:      "plain",
		Vcpus:             4,
		MemoryInMegabytes: 1024,
	}
	instCreateStruct := NewGanetiInstance(p)
	if instCreateStruct.Version != 1 {
		t.Errorf("instanceCreateObject contains the wrong version (expected: '1', got: '%d')", instCreateStruct.Version)
	}

	if instCreateStruct.Hypervisor != "fake" {
		t.Errorf("instanceCreateObject contains the wrong hypervisor (expected: 'fake', got: '%s')", instCreateStruct.Hypervisor)
	}

	if instCreateStruct.DiskTemplate != "plain" {
		t.Errorf("instanceCreateObject contains the wrong disk template (expected: 'plain', got: '%s')", instCreateStruct.DiskTemplate)
	}

	if instCreateStruct.InstanceName != "my-test-instance" {
		t.Errorf("instanceCreateObject contains the wrong instance name (expected: 'my-test-instance', got: '%s')", instCreateStruct.InstanceName)
	}

	// add many more checks
}

func TestGetRapiConnection(t *testing.T) {
	config.Parse()
	url, _ := getRapiConnection("production-cluster")
	if url != "https://gnt-cc:somepassword@prod-cluster.example.com:5080" {
		t.Errorf("Expected URL 'https://gnt-cc:somepassword@prod-cluster.example.com:5080', got: '%s'", url)
	}
}
