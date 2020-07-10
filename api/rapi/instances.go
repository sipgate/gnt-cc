package rapi

import (
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"gnt-cc/model"
)

func NewGanetiInstance(instanceDetails CreateInstanceParameters) InstanceCreate {
	var inst InstanceCreate
	inst.InstanceName = instanceDetails.InstanceName
	inst.DiskTemplate = instanceDetails.DiskTemplate

	if len(instanceDetails.Nics) == 0 {
		inst.Nics = make([]GanetiNic, 0)
	} else {
		for _, nic := range instanceDetails.Nics {
			inst.Nics = append(inst.Nics, nic)
		}
	}

	if len(instanceDetails.Disks) == 0 {
		inst.Disks = make([]GanetiDisk, 0)
	} else {
		for _, disk := range instanceDetails.Disks {
			inst.Disks = append(inst.Disks, disk)
		}
	}

	if instanceDetails.Vcpus > 0 {
		inst.BeParams.Vcpus = instanceDetails.Vcpus
	}

	if instanceDetails.MemoryInMegabytes > 0 {
		inst.BeParams.Memory = instanceDetails.MemoryInMegabytes
	}

	inst.Version = 1
	inst.Mode = "create"
	inst.Hypervisor = "fake"
	inst.Iallocator = "hail"
	inst.OsType = "noop"
	inst.ConflictsCheck = false
	inst.IPCheck = false
	inst.NameCheck = false
	inst.NoInstall = true
	inst.WaitForSync = false
	return inst
}

func GetInstances(clusterConfig config.ClusterConfig) ([]model.GntInstance, error) {
	response, err := Get(clusterConfig, "/2/instances?bulk=1")

	if err != nil {
		return nil, err
	}

	var instanceData []Instance
	err = json.Unmarshal([]byte(response), &instanceData)

	if err != nil {
		return nil, err
	}

	instances := make([]model.GntInstance, len(instanceData))

	for i, instance := range instanceData {
		instances[i] = model.GntInstance{
			Name:           instance.Name,
			PrimaryNode:    instance.Pnode,
			SecondaryNodes: instance.Snodes,
			Disks:          nil,
			CpuCount:       instance.BeParams.Vcpus,
			MemoryTotal:    instance.BeParams.Memory,
		}
	}

	return instances, nil
}

func GetInstance(clusterConfig config.ClusterConfig, instanceName string) (model.GntInstance, error) {
	response, err := Get(clusterConfig, fmt.Sprintf("/2/instances/%s", instanceName))

	if err != nil {
		return model.GntInstance{}, err
	}

	var instanceData Instance
	err = json.Unmarshal([]byte(response), &instanceData)

	if err != nil {
		return model.GntInstance{}, err
	}

	return model.GntInstance{
		Name:           instanceData.Name,
		PrimaryNode:    instanceData.Pnode,
		SecondaryNodes: instanceData.Snodes,
		Disks:          nil,
		CpuCount:       instanceData.BeParams.Vcpus,
		MemoryTotal:    instanceData.BeParams.Memory,
	}, nil
}
