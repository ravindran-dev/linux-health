package process

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ProcInfo struct {
	PID   int
	Cmd   string
	MemKB uint64
}

func TopMemoryProcess() (*ProcInfo, error) {
	var top *ProcInfo

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		pid, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}

		statusPath := filepath.Join("/proc", e.Name(), "status")
		data, err := os.ReadFile(statusPath)
		if err != nil {
			continue
		}

		var mem uint64
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "VmRSS:") {
				fields := strings.Fields(line)
				mem, _ = strconv.ParseUint(fields[1], 10, 64)
				break
			}
		}

		if top == nil || mem > top.MemKB {
			cmdData, _ := os.ReadFile(filepath.Join("/proc", e.Name(), "comm"))
			top = &ProcInfo{
				PID:   pid,
				Cmd:   strings.TrimSpace(string(cmdData)),
				MemKB: mem,
			}
		}
	}

	return top, nil
}
