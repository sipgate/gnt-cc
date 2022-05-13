package repository

import (
	"encoding/json"
	"fmt"
	"gnt-cc/model"
	"gnt-cc/query"
	"gnt-cc/rapi_client"
)

type (
	queryPerformer interface {
		Perform(client rapi_client.Client, config query.RequestConfig) ([]query.Resource, error)
	}

	InstanceRepository struct {
		RAPIClient     rapi_client.Client
		QueryPerformer queryPerformer
	}
)

func (repo *InstanceRepository) Get(clusterName string, instanceName string) (model.InstanceResult, error) {
	slug := fmt.Sprintf("/2/instances/%s", instanceName)
	response, err := repo.RAPIClient.Get(clusterName, slug)

	if err != nil {
		return model.InstanceResult{}, err
	}

	if response.Status == 404 {
		return model.InstanceResult{
			Found: false,
		}, nil
	}

	var parsedInstance rapiInstanceResponse
	err = json.Unmarshal([]byte(response.Body), &parsedInstance)

	if err != nil {
		return model.InstanceResult{}, err
	}

	return model.InstanceResult{
		Found:       true,
		NetworkPort: parsedInstance.NetworkPort,
		Instance: model.GntInstance{
			Name:           parsedInstance.Name,
			PrimaryNode:    parsedInstance.Pnode,
			SecondaryNodes: parsedInstance.Snodes,
			CpuCount:       parsedInstance.BeParams.VCPUs,
			MemoryTotal:    parsedInstance.BeParams.Memory,
			IsRunning:      parsedInstance.OperState,
			OffersVNC:      parsedInstance.HvParams.VncBindAddress != "",
			Disks:          extractDisks(parsedInstance),
			Nics:           extractNics(parsedInstance),
			Tags:           parsedInstance.Tags,
		},
	}, nil
}

func (repo *InstanceRepository) GetAll(clusterName string) ([]model.GntInstance, error) {
	resources, err := repo.QueryPerformer.Perform(repo.RAPIClient, query.RequestConfig{
		ClusterName:  clusterName,
		ResourceType: "instance",
		Fields: []string{
			"name",
			"pnode",
			"snodes",
			"beparams",
			"oper_state",
			"hvparams",
		},
	})

	if err != nil {
		return nil, err
	}

	return parseInstanceResourceArray(resources)
}

func (repo *InstanceRepository) GetAllNames(clusterName string) ([]string, error) {
	response, err := repo.RAPIClient.Get(clusterName, "/2/instances")

	if err != nil {
		return nil, err
	}

	var instanceList rapiInstanceNamesResponse
	err = json.Unmarshal([]byte(response.Body), &instanceList)

	if err != nil {
		return nil, err
	}

	instanceNames := make([]string, len(instanceList))
	for i, instance := range instanceList {
		instanceNames[i] = instance.ID
	}

	return instanceNames, nil
}

func parseInstanceResourceArray(resources []query.Resource) ([]model.GntInstance, error) {
	instances := make([]model.GntInstance, len(resources))

	for i, resource := range resources {
		instance, err := parseInstanceResource(resource)

		if err != nil {
			return nil, err
		}

		instances[i] = instance
	}

	return instances, nil
}

func parseInstanceResource(resource query.Resource) (model.GntInstance, error) {
	parsed, err := convertQueryResourceToStruct(resource)

	if err != nil {
		return model.GntInstance{}, err
	}

	return model.GntInstance{
		Name:           parsed.Name,
		PrimaryNode:    parsed.Pnode,
		SecondaryNodes: parsed.Snodes,
		CpuCount:       parsed.BeParams.VCPUs,
		MemoryTotal:    parsed.BeParams.MaxMem,
		IsRunning:      parsed.OperState,
		OffersVNC:      parsed.HvParams.VncBindAddress != "",
	}, nil
}

func convertQueryResourceToStruct(resource query.Resource) (instanceQueryResource, error) {
	var parsed instanceQueryResource
	data, err := json.Marshal(resource)

	if err != nil {
		return parsed, err
	}

	err = json.Unmarshal(data, &parsed)

	return parsed, err
}

func extractDisks(instance rapiInstanceResponse) []model.GntDisk {
	disks := []model.GntDisk{}
	diskNames := instance.DiskNames
	diskSizes := instance.DiskSizes

	for i, uuid := range instance.DiskUuids {
		var name string

		if diskNameAsString, ok := diskNames[i].(string); ok {
			name = diskNameAsString
		} else {
			name = fmt.Sprintf("Disk %d", i)
		}

		disks = append(disks, model.GntDisk{
			Uuid:     uuid,
			Name:     name,
			Template: instance.DiskTemplate,
			Capacity: diskSizes[i],
		})
	}

	return disks
}

func extractNics(instance rapiInstanceResponse) []model.GntNic {
	nics := []model.GntNic{}

	for i, uuid := range instance.NicUuids {
		mode := instance.NicModes[i]
		mac := instance.NicMacs[i]
		vlan := instance.CustomNicParams[i].Vlan
		var name string

		if nicNameAsString, ok := instance.NicNames[i].(string); ok {
			name = nicNameAsString
		} else {
			name = fmt.Sprintf("NIC %d", i)
		}

		bridge := ""
		if nicBridgeAsString, ok := instance.NicBridges[i].(string); ok {
			bridge = nicBridgeAsString
		}

		nics = append(nics, model.GntNic{
			Uuid:   uuid,
			Mode:   mode,
			Mac:    mac,
			Name:   name,
			Bridge: bridge,
			Vlan:   vlan,
		})
	}

	return nics
}
