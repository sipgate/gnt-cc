package repository_test

import (
	"gnt-cc/mocking"
	"gnt-cc/model"
	"gnt-cc/rapi_client"
	"gnt-cc/repository"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestJobRepoGetFuncReturnsSuccessfulResult_OnValidResponses(t *testing.T) {
	tests := []struct {
		name           string
		file           string
		expectedResult model.GntJob
	}{{
		"InstanceCreate_With_Import",
		"../testfiles/rapi_responses/valid_job_instance_create_with_import_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "INSTANCE_CREATE(bart.local)",
			ReceivedAt:  1644229365,
			StartedAt:   1644229365,
			EndedAt:     -1,
			Status:      "running",
			Log: &[]model.GntJobLogEntry{
				{Serial: 1, Message: " - INFO: Selected nodes for instance bart.local via iallocator hail: ganeti-node01.local, ganeti-node02.local", StartedAt: 1644229367},
				{Serial: 2, Message: "creating instance disks...", StartedAt: 1644229368},
				{Serial: 3, Message: "adding instance bart.local to cluster config", StartedAt: 1644229375},
				{Serial: 4, Message: "adding disks to cluster config", StartedAt: 1644229375},
				{Serial: 5, Message: "wiping instance disks...", StartedAt: 1644229375},
				{Serial: 6, Message: " - INFO: * Wiping disk 0", StartedAt: 1644229375},
				{Serial: 7, Message: " - INFO:  - done: 2.5% ETA: 1m 26s", StartedAt: 1644229377},
				{Serial: 8, Message: " - INFO:  - done: 70.0% ETA: 27s", StartedAt: 1644229439},
				{Serial: 9, Message: " - INFO: Waiting for instance bart.local to sync disks", StartedAt: 1644229467},
				{Serial: 10, Message: " - INFO: Instance bart.local's disks are in sync", StartedAt: 1644229467},
				{Serial: 11, Message: " - INFO: Waiting for instance bart.local to sync disks", StartedAt: 1644229467},
				{Serial: 12, Message: " - INFO: Instance bart.local's disks are in sync", StartedAt: 1644229468},
				{Serial: 13, Message: "preparing remote import...", StartedAt: 1644229468},
				{Serial: 14, Message: "Disk 0 is now listening", StartedAt: 1644229471},
				{Serial: 15, Message: "Importing Disks from: [192.0.2.1:42443]", StartedAt: 1644229471},
				{Serial: 16, Message: "Disk 0 is now receiving data", StartedAt: 1644229479}}},
	}, {
		"InstanceCreate",
		"../testfiles/rapi_responses/valid_job_instance_create_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "INSTANCE_CREATE(homer.local)",
			ReceivedAt:  1645198959,
			StartedAt:   1645198960,
			EndedAt:     1645199064,
			Status:      "success",
			Log: &[]model.GntJobLogEntry{
				{Serial: 1, Message: " - INFO: Selected nodes for instance homer.local via iallocator hail: ganeti-node01.local, ganeti-node02.local", StartedAt: 1645198964},
				{Serial: 2, Message: "creating instance disks...", StartedAt: 1645198965},
				{Serial: 3, Message: "adding instance homer.local to cluster config", StartedAt: 1645198970},
				{Serial: 4, Message: "adding disks to cluster config", StartedAt: 1645198970},
				{Serial: 5, Message: "wiping instance disks...", StartedAt: 1645198970},
				{Serial: 6, Message: " - INFO: * Wiping disk 0", StartedAt: 1645198971},
				{Serial: 7, Message: " - INFO:  - done: 6.7% ETA: 43s", StartedAt: 1645198974},
				{Serial: 8, Message: "checking mirrors status", StartedAt: 1645199011},
				{Serial: 9, Message: " - INFO: Instance homer.local's disks are in sync", StartedAt: 1645199011},
				{Serial: 10, Message: "checking mirrors status", StartedAt: 1645199011},
				{Serial: 11, Message: " - INFO: Instance homer.local's disks are in sync", StartedAt: 1645199011},
				{Serial: 12, Message: "pausing disk sync to install instance OS", StartedAt: 1645199011},
				{Serial: 13, Message: "running the instance OS create scripts...", StartedAt: 1645199011},
				{Serial: 14, Message: "resuming disk sync", StartedAt: 1645199041},
				{Serial: 15, Message: "starting instance...", StartedAt: 1645199063}}},
	}, {
		"InstanceRemove",
		"../testfiles/rapi_responses/valid_job_instance_remove_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "INSTANCE_REMOVE(homer.local)",
			ReceivedAt:  1645198879,
			StartedAt:   1645198879,
			EndedAt:     1645198913,
			Status:      "success",
			Log:         &[]model.GntJobLogEntry{}},
	}, {
		"InstanceMigrate",
		"../testfiles/rapi_responses/valid_job_instance_migrate_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "INSTANCE_MIGRATE(lisa.local)",
			ReceivedAt:  1645197691,
			StartedAt:   1645197692,
			EndedAt:     1645197708,
			Status:      "success",
			Log: &[]model.GntJobLogEntry{
				{Serial: 1, Message: "Migrating instance lisa.local", StartedAt: 1645197694},
				{Serial: 2, Message: "checking disk consistency between source and target", StartedAt: 1645197694},
				{Serial: 3, Message: "closing instance disks on node ganeti-node01.local", StartedAt: 1645197694},
				{Serial: 4, Message: "changing into standalone mode", StartedAt: 1645197695},
				{Serial: 5, Message: "changing disks into dual-master mode", StartedAt: 1645197695},
				{Serial: 6, Message: "wait until resync is done", StartedAt: 1645197696},
				{Serial: 7, Message: "opening instance disks on node ganeti-node02.local in shared mode", StartedAt: 1645197696},
				{Serial: 8, Message: "opening instance disks on node ganeti-node01.local in shared mode", StartedAt: 1645197696},
				{Serial: 9, Message: "preparing ganeti-node01.local to accept the instance", StartedAt: 1645197696},
				{Serial: 10, Message: "migrating instance to ganeti-node01.local", StartedAt: 1645197697},
				{Serial: 11, Message: "starting memory transfer", StartedAt: 1645197697},
				{Serial: 12, Message: "memory transfer complete", StartedAt: 1645197699},
				{Serial: 13, Message: "closing instance disks on node ganeti-node02.local", StartedAt: 1645197704},
				{Serial: 14, Message: "wait until resync is done", StartedAt: 1645197704},
				{Serial: 15, Message: "changing into standalone mode", StartedAt: 1645197704},
				{Serial: 16, Message: "changing disks into single-master mode", StartedAt: 1645197705},
				{Serial: 17, Message: "wait until resync is done", StartedAt: 1645197706},
				{Serial: 18, Message: "done", StartedAt: 1645197706}}},
	}, {
		"InstanActivateDisks",
		"../testfiles/rapi_responses/valid_job_instance_activate_disks_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "INSTANCE_ACTIVATE_DISKS(homer.local)",
			ReceivedAt:  1645197738,
			StartedAt:   1645197738,
			EndedAt:     1645197741,
			Status:      "success",
			Log:         &[]model.GntJobLogEntry{}},
	}, {
		"InstanQueryData",
		"../testfiles/rapi_responses/valid_job_instance_query_data_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "INSTANCE_QUERY_DATA",
			ReceivedAt:  1645201116,
			StartedAt:   1645201117,
			EndedAt:     1645201118,
			Status:      "success",
			Log:         &[]model.GntJobLogEntry{}},
	}, {
		"ClusterVerify",
		"../testfiles/rapi_responses/valid_job_cluster_verify_response.json",
		model.GntJob{
			ID:          12345,
			ClusterName: "test",
			Summary:     "CLUSTER_VERIFY",
			ReceivedAt:  1645199102,
			StartedAt:   1645199102,
			EndedAt:     1645199104,
			Status:      "success",
			Log:         &[]model.GntJobLogEntry{}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validResponse, err := ioutil.ReadFile(tt.file)
			assert.NoError(t, err)

			client := mocking.NewRAPIClient()
			client.On("Get", mock.Anything, mock.Anything).
				Once().Return(rapi_client.Response{Status: 200, Body: string(validResponse)}, nil)
			repo := repository.JobRepository{RAPIClient: client}
			result, err := repo.Get("test", "12345")

			assert.NoError(t, err)
			assert.True(t, result.Found)
			assert.EqualValues(t, tt.expectedResult,
				result.Job)
		})
	}
}
