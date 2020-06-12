package dummy

import (
	"fmt"
	"gnt-cc/model"
)

func GetInstances(count int) []model.GntInstance {
	instances := make([]model.GntInstance, count)

	for i := 0; i < count; i++ {
		instances[i] = model.GntInstance{
			Name:        fmt.Sprintf("%d.instances.dummy.example.com", i),
			PrimaryNode: "ganeti-node03.example.com",
			SecondaryNodes: []string{
				"ganeti-node01.example.com",
				"ganeti-node02.example.com",
			},
			Disks: []model.Disk{
				{
					Name: "disk0",
					Size: 500000,
					Uuid: "12345678-12345678-12345678",
				},
			},
			MemoryTotal: 4096,
			CpuCount:    6,
		}
	}

	return instances
}

func GetInstance(name string) model.GntInstance {
	return model.GntInstance{
		Name:        name,
		PrimaryNode: "dummy02.example.com",
		SecondaryNodes: []string{
			"dummy01.example.com",
			"dummy03.example.com",
		},
		Disks: []model.Disk{
			{
				Name: "disk0",
				Size: 25000,
				Uuid: "12345678-12345678-12345678",
			},
			{
				Name: "disk1",
				Size: 30000,
				Uuid: "12345678-12345678-12w45678",
			},
		},
		MemoryTotal: 4096,
		CpuCount:    6,
	}
}
