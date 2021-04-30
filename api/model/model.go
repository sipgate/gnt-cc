package model

type GntDisk struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Template string `json:"template"`
	Capacity int    `json:"capacity"`
}

type GntNic struct {
	Uuid   string `json:"uuid"`
	Mode   string `json:"mode"`
	Name   string `json:"name"`
	Mac    string `json:"mac"`
	Bridge string `json:"bridge"`
}

type GntInstance struct {
	Name           string    `json:"name"`
	PrimaryNode    string    `json:"primaryNode"`
	SecondaryNodes []string  `json:"secondaryNodes"`
	CpuCount       int       `json:"cpuCount"`
	MemoryTotal    int       `json:"memoryTotal"`
	IsRunning      bool      `json:"isRunning"`
	OffersVNC      bool      `json:"offersVnc"`
	Disks          []GntDisk `json:"disks"`
	Nics           []GntNic  `json:"nics"`
	Tags           []string  `json:"tags"`
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
	DiskFree           int      `json:"diskFree"`
	CPUCount           int      `json:"cpuCount"`
	PrimaryInstances   []string `json:"primaryInstances"`
	SecondaryInstances []string `json:"secondaryInstances"`
	IsMaster           bool     `json:"isMaster"`
	IsMasterCandidate  bool     `json:"isMasterCandidate"`
	IsMasterCapable    bool     `json:"isMasterCapable"`
	IsDrained          bool     `json:"isDrained"`
	IsOffline          bool     `json:"isOffline"`
	IsVMCapable        bool     `json:"isVmCapable"`
}

type GntJob struct {
	ID         int    `json:"id"`
	Summary    string `json:"summary"`
	ReceivedAt int    `json:"receivedAt"`
	StartedAt  int    `json:"startedAt"`
	EndedAt    int    `json:"endedAt"`
	Status     string `json:"status"`
}
