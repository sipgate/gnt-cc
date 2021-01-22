package model

type GntInstance struct {
	Name           string   `json:"name"`
	PrimaryNode    string   `json:"primaryNode"`
	SecondaryNodes []string `json:"secondaryNodes"`
	CpuCount       int      `json:"cpuCount"`
	MemoryTotal    int      `json:"memoryTotal"`
	IsRunning      bool     `json:"isRunning"`
	OffersVNC      bool     `json:"offersVnc"`
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
	DiskTotal          int      `json:"diskTotal"`
	CPUCount           int      `json:"cpuCount"`
	PrimaryInstances   []string `json:"primaryInstances"`
	SecondaryInstances []string `json:"secondaryInstances"`
}

type GntJob struct {
	ID         int    `json:"id"`
	Summary    string `json:"summary"`
	ReceivedAt int    `json:"receivedAt"`
	StartedAt  int    `json:"startedAt"`
	EndedAt    int    `json:"endedAt"`
	Status     string `json:"status"`
}
