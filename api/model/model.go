package model

type GntInstance struct {
	Name           string   `json:"name"`
	PrimaryNode    string   `json:"primaryNode"`
	SecondaryNodes []string `json:"secondaryNodes"`
	CpuCount       int      `json:"cpuCount"`
	MemoryTotal    int      `json:"memoryTotal"`
}

type GntCluster struct {
	Name        string `json:"name"`
	Hostname    string `json:"hostname"`
	Description string `json:"description"`
	Port        int    `json:"port"`
}

type GntNode struct {
	Name               string   `json:"name"`
	MemoryTotal        int      `json:"memoryTotal"`
	MemoryFree         int      `json:"memoryFree"`
	PrimaryInstances   []string `json:"primaryInstances"`
	SecondaryInstances []string `json:"secondaryInstances"`
}
