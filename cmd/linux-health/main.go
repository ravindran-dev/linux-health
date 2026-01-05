package main

import (
	"fmt"
	"os"

	"github.com/ravindran-dev/linux-health/internal/disk"
	"github.com/ravindran-dev/linux-health/internal/network"
	"github.com/ravindran-dev/linux-health/internal/process"
	"github.com/ravindran-dev/linux-health/internal/service"
	"github.com/ravindran-dev/linux-health/internal/system"
	"github.com/ravindran-dev/linux-health/pkg/output"
)

func main() {
	// ---- SYSTEM STATS ----
	stats, err := system.GetSystemStats()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}

	health := system.GenerateHealth(stats)

	// ---- TOP MEMORY PROCESS ----
	topName := ""
	topPID := 0
	topMB := uint64(0)

	topProc, _ := process.TopMemoryProcess()
	if topProc != nil {
		topName = topProc.Cmd
		topPID = topProc.PID
		topMB = topProc.MemKB / 1024
	}

	// ---- DISK HOTSPOTS (FAST) ----
	hotspots := []string{}
	dirs, _ := disk.TopDirs(os.Getenv("HOME"), 3)
	for _, d := range dirs {
		hotspots = append(hotspots,
			fmt.Sprintf("%s %s", d.Path, d.SizeHuman),
		)
	}

	// ---- FAILED SERVICES ----
	failedServices := []string{}
	failed, _ := service.FailedServices()
	for _, s := range failed {
		failedServices = append(failedServices,
			fmt.Sprintf("%s (%s)", s.Name, s.State),
		)
	}

	// ---- NETWORK ----
	networkState := network.Summary()

	// ---- OUTPUT ----
	output.PrintBlock(
		stats,
		health,
		topName,
		topPID,
		topMB,
		networkState,
		failedServices,
		hotspots,
	)

	os.Exit(health.ExitCode)
}
