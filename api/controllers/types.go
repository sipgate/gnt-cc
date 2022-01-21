package controllers

import "gnt-cc/model"

type (
	nodeRepository interface {
		Get(clusterName string, nodeName string) (model.NodeResult, error)
		GetAll(clusterName string) ([]model.GntNode, error)
	}

	instanceRepository interface {
		Get(clusterName string, instanceName string) (model.InstanceResult, error)
		GetAll(clusterName string) ([]model.GntInstance, error)
	}

	jobRepository interface {
		GetAll(clusterName string) ([]model.GntJob, error)
		Get(clusterName, jobID string) (model.JobResult, error)
	}

	instanceActions interface {
		Start(clusterName string, instanceName string) (int, error)
		Restart(clusterName string, instanceName string) (int, error)
		Shutdown(clusterName string, instanceName string) (int, error)
	}
)
