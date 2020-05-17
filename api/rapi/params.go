package rapi

type CreateInstanceParameters struct {
	InstanceName      string       `json:"instance_name"`
	DiskTemplate      string       `json:"disk_template"`
	Vcpus             int          `json:"vcpus,omitempty"`
	MemoryInMegabytes int          `json:"memoryInMegabytes,omitempty"`
	Nics              []GanetiNic  `json:"nics,omitempty"`
	Disks             []GanetiDisk `json:"disks,omitempty"`
}
