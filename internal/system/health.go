package system

type Status string

const (
	StatusOK       Status = "OK"
	StatusWarning  Status = "WARNING"
	StatusCritical Status = "CRITICAL"
)

type MetricStatus struct {
	Value  float64
	Status Status
}

type HealthReport struct {
	Score    int
	CPU      MetricStatus
	Memory   MetricStatus
	Disk     MetricStatus
	Load     MetricStatus
	ExitCode int
}

func evaluate(value, warn, crit float64) MetricStatus {
	switch {
	case value >= crit:
		return MetricStatus{value, StatusCritical}
	case value >= warn:
		return MetricStatus{value, StatusWarning}
	default:
		return MetricStatus{value, StatusOK}
	}
}

func GenerateHealth(stats SystemStats) HealthReport {
	cpu := evaluate(stats.CPUUsage, 75, 90)
	mem := evaluate(stats.MemUsagePct, 70, 85)
	disk := evaluate(stats.DiskUsage, 80, 95)
	load := evaluate(stats.Load1, 2.0, 4.0)

	score := 100
	for _, s := range []MetricStatus{cpu, mem, disk, load} {
		if s.Status == StatusWarning {
			score -= 10
		}
		if s.Status == StatusCritical {
			score -= 25
		}
	}

	exitCode := 0
	if cpu.Status == StatusCritical || mem.Status == StatusCritical ||
		disk.Status == StatusCritical || load.Status == StatusCritical {
		exitCode = 2
	} else if cpu.Status == StatusWarning || mem.Status == StatusWarning ||
		disk.Status == StatusWarning || load.Status == StatusWarning {
		exitCode = 1
	}

	return HealthReport{
		Score:    score,
		CPU:      cpu,
		Memory:   mem,
		Disk:     disk,
		Load:     load,
		ExitCode: exitCode,
	}
}
