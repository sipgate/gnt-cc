package model

type Disk struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Uuid string `json:"uuid"`
}

type GntInstance struct {
	Name           string   `json:"name"`
	PrimaryNode    string   `json:"primaryNode"`
	SecondaryNodes []string `json:"secondaryNodes"`
	Disks          []Disk   `json:"disks"`
	CpuCount       int      `json:"cpuCount"`
	MemoryTotal    int      `json:"memoryTotal"`
}
