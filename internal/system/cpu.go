package system

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type cpuStat struct {
	idle  uint64
	total uint64
}

func readCPU() (cpuStat, error) {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return cpuStat{}, err
	}

	fields := strings.Fields(strings.Split(string(data), "\n")[0])
	var idle, total uint64

	for i, v := range fields[1:] {
		val, _ := strconv.ParseUint(v, 10, 64)
		total += val
		if i == 3 {
			idle = val
		}
	}

	return cpuStat{idle, total}, nil
}

func CPUUsage() (float64, error) {
	s1, err := readCPU()
	if err != nil {
		return 0, err
	}

	time.Sleep(500 * time.Millisecond)

	s2, err := readCPU()
	if err != nil {
		return 0, err
	}

	idle := s2.idle - s1.idle
	total := s2.total - s1.total

	return 100 * (1 - float64(idle)/float64(total)), nil
}
