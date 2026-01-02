package system

type SystemStats struct {
	CPUUsage    float64
	MemTotalMB  uint64
	MemUsedMB   uint64
	MemUsagePct float64
	Load1       float64
	Load5       float64
	Load15      float64
	DiskUsage   float64
	Uptime      string
}
