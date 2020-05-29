package model

type CreateInstanceRequest struct {
	Name            string `json:"name" binding:"required"`
	MemoryMegaBytes int    `json:"memory_mega_bytes" binding:"required"`
	VCpuCores       int    `json:"v_cpu_cores" binding:"required"`
}
