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
	stats, err := system.GetSystemStats()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}

	health := system.GenerateHealth(stats)

	// ---- TOP MEMORY PROCESS ----
	topName := ""
	topPID := 0
	var topMB uint64

	if topProc, _ := process.TopMemoryProcess(); topProc != nil {
		topName = topProc.Cmd
		topPID = topProc.PID
		topMB = topProc.MemKB / 1024
	}

	// ---- DISK HOTSPOTS ----
	hotspots := []string{}
	dirs, _ := disk.TopDirs(os.Getenv("HOME"), 2)
	for _, d := range dirs {
		hotspots = append(
			hotspots,
			fmt.Sprintf("%s %s", d.Path, d.SizeHuman),
		)
	}

	// ---- SERVICES ----
	services := []string{}
	failed, _ := service.FailedServices()
	if len(failed) == 0 {
		services = append(services, "✓ No failed services")
	} else {
		for _, s := range failed {
			services = append(
				services,
				fmt.Sprintf("%s (%s)", s.Name, s.State),
			)
		}
	}

	// ---- NETWORK ----
	networkState := "inactive"
	if nets, _ := network.ActiveInterfaces(); len(nets) > 0 {
		networkState = "active"
	}

	// ---- FINAL BOX OUTPUT ----
	output.PrintBlock(
		stats,
		health,
		topName,
		topPID,
		topMB,
		networkState,
		services,
		hotspots,
	)

	os.Exit(health.ExitCode)
}
