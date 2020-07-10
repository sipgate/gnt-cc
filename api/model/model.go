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

type GntCluster struct {
	Name        string `json:"name"`
	Hostname    string `json:"hostname"`
	Description string `json:"description"`
	Port        int    `json:"port"`
}

type GntJob struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type GntNode struct {
	Name               string        `json:"name"`
	MemoryTotal        int           `json:"memoryTotal"`
	MemoryFree         int           `json:"memoryFree"`
	PrimaryInstances   []GntInstance `json:"primaryInstances"`
	SecondaryInstances []GntInstance `json:"secondaryInstances"`
}

type JobStatusCount struct {
	Canceled int `json:"canceled"`
	Error    int `json:"error"`
	Pending  int `json:"pending"`
	Queued   int `json:"queued"`
	Success  int `json:"success"`
	Waiting  int `json:"waiting"`
}
