package process

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type CPUProcess struct {
	PID int
	CPU float64
	MEM float64
	Cmd string
}

func TopCPUProcesses(limit int) ([]CPUProcess, error) {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	var procs []CPUProcess

	for _, e := range entries {
		pid, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}

		statPath := filepath.Join("/proc", e.Name(), "stat")
		statusPath := filepath.Join("/proc", e.Name(), "status")
		cmdPath := filepath.Join("/proc", e.Name(), "comm")

		statData, err := os.ReadFile(statPath)
		if err != nil {
			continue
		}

		fields := strings.Fields(string(statData))
		if len(fields) < 17 {
			continue
		}

		utime, _ := strconv.ParseFloat(fields[13], 64)
		stime, _ := strconv.ParseFloat(fields[14], 64)
		cpu := utime + stime

		var mem float64
		statusData, err := os.ReadFile(statusPath)
		if err == nil {
			for _, line := range strings.Split(string(statusData), "\n") {
				if strings.HasPrefix(line, "VmRSS:") {
					parts := strings.Fields(line)
					val, _ := strconv.ParseFloat(parts[1], 64)
					mem = val / 1024
					break
				}
			}
		}

		cmdData, _ := os.ReadFile(cmdPath)

		procs = append(procs, CPUProcess{
			PID: pid,
			CPU: cpu,
			MEM: mem,
			Cmd: strings.TrimSpace(string(cmdData)),
		})
	}

	sort.Slice(procs, func(i, j int) bool {
		return procs[i].CPU > procs[j].CPU
	})

	if len(procs) > limit {
		procs = procs[:limit]
	}

	return procs, nil
}
