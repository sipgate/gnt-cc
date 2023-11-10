package testcluster

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
)

type TestCluster struct {
	ctx         context.Context
	container   testcontainers.Container
	containerIP string
}

func New() (*TestCluster, error) {
	ctx := context.Background()

	ganeticC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "ghcr.io/sipgate/ganeti-docker/image:latest",
			ExposedPorts: []string{"5080/tcp"},
			HostConfigModifier: func(config *container.HostConfig) {
				config.CapAdd = []string{"NET_ADMIN"}
			},
			WaitingFor: wait.ForLog(".*ganeti-rapi daemon startup").AsRegexp(),
		},
		Started: true,
	})
	if err != nil {
		return nil, err
	}

	containerIP, err := ganeticC.ContainerIP(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot determine test cluster container IP %s\n", err)
	}

	testCluster := TestCluster{
		ctx:         ctx,
		container:   ganeticC,
		containerIP: containerIP,
	}

	return &testCluster, nil
}

func (testCluster *TestCluster) GetClusterHost() string {
	return testCluster.containerIP
}

func (testCluster *TestCluster) Terminate() {
	_ = testCluster.container.Terminate(testCluster.ctx)
}

func (testCluster *TestCluster) CreateInstance(instanceName string) error {
	_, _, err := testCluster.container.Exec(testCluster.ctx, []string{
		"/opt/ganeti-vcluster/node1/cmd",
		"gnt-instance",
		"add",
		"-t",
		"diskless",
		"--no-ip-check",
		"--no-name-check",
		"--no-install",
		"-o", "noop",
		"--no-nics",
		instanceName,
	})
	if err != nil {
		return fmt.Errorf("cannot create instance '%s' in cluster: %e\n", instanceName, err)
	}

	return nil
}

func (testCluster TestCluster) RemoveInstance(instanceName string) error {
	_, _, err := testCluster.container.Exec(testCluster.ctx, []string{
		"/opt/ganeti-vcluster/node1/cmd",
		"gnt-instance",
		"remove",
		"--ignore-failures",
		instanceName,
	})
	if err != nil {
		return fmt.Errorf("cannot remove instance '%s' in cluster: %e\n", instanceName, err)
	}

	return nil
}

func (testCluster *TestCluster) DebugListInstances() error {
	_, reader, err := testCluster.container.Exec(testCluster.ctx, []string{
		"/opt/ganeti-vcluster/node1/cmd",
		"gnt-instance",
		"list",
	})
	if err != nil {
		return errors.New("cannot list instances in cluster: %s\n")
	}

	out, err := io.ReadAll(reader)
	if err != nil {
		return errors.New("cannot list instances in cluster: %s\n")
	}

	fmt.Printf("%s", out)

	return nil
}
