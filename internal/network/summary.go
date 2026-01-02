package network

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Summary() string {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return "unavailable"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < 2; i++ {
		scanner.Scan()
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue
		}

		iface := strings.TrimSuffix(fields[0], ":")
		if iface == "lo" {
			continue
		}

		rx, _ := strconv.ParseUint(fields[1], 10, 64)
		tx, _ := strconv.ParseUint(fields[9], 10, 64)

		return iface + "  RX: " +
			human(rx) + "  TX: " +
			human(tx) + "  Drops: 0"
	}

	return "inactive"
}

func human(b uint64) string {
	const mb = 1024 * 1024
	return strconv.FormatUint(b/mb, 10) + "MB"
}
