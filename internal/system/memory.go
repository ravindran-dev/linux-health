
package system

import (
	"os"
	"strconv"
	"strings"
)

func MemoryUsage() (totalMB, usedMB uint64, percent float64, err error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}

	var total, available uint64

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			total, _ = strconv.ParseUint(fields[1], 10, 64)
		case "MemAvailable:":
			available, _ = strconv.ParseUint(fields[1], 10, 64)
		}
	}

	used := total - available
	return total / 1024, used / 1024, float64(used) / float64(total) * 100, nil
}
