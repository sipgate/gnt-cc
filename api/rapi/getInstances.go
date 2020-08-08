package rapi

import (
	"gnt-cc/config"
	"gnt-cc/model"
	"gnt-cc/utils"
)

type expectedInstanceResource struct {
	Name     string   `json:"name"`
	Pnode    string   `json:"pnode"`
	Snodes   []string `json:"snodes"`
	BeParams struct {
		AutoBalance    bool `json:"auto_balance"`
		SpindleUse     int  `json:"spindle_use"`
		VCPUs          int  `json:"vcpus"`
		MinMem         int  `json:"minmen"`
		MaxMem         int  `json:"maxmem"`
		AlwaysFailover bool `json:"always_failover"`
	} `json:"beparams"`
}

func parseInstanceFromResource(resource ResourceFieldValueMap) (model.GntInstance, error) {
	var instanceResource expectedInstanceResource

	err := utils.ConvertMapToStruct(resource, &instanceResource)

	if err != nil {
		return model.GntInstance{}, err
	}

	return model.GntInstance{
		Name:           instanceResource.Name,
		PrimaryNode:    instanceResource.Pnode,
		SecondaryNodes: instanceResource.Snodes,
		Disks:          nil,
		CpuCount:       instanceResource.BeParams.VCPUs,
		MemoryTotal:    instanceResource.BeParams.MaxMem, // MaxMem correct?
	}, nil
}

func GetInstances(clusterConfig config.ClusterConfig) ([]model.GntInstance, error) {
	fields := []string{
		"name",
		"pnode",
		"snodes",
		"beparams",
	}

	var instances []model.GntInstance

	err := queryAndTransformResources(queryRequestConfig{
		clusterConfig: clusterConfig,
		resourceType:  "instance",
		fields:        fields,
	}, func(resource ResourceFieldValueMap) error {
		instance, err := parseInstanceFromResource(resource)

		if err != nil {
			return err
		}

		instances = append(instances, instance)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return instances, nil
}
