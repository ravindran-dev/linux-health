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

	topName := ""
	topPID := 0
	var topMB uint64

	if topProc, _ := process.TopMemoryProcess(); topProc != nil {
		topName = topProc.Cmd
		topPID = topProc.PID
		topMB = topProc.MemKB / 1024
	}

	hotspots := []string{}
	if dirs, _ := disk.TopDirs(os.Getenv("HOME"), 2); len(dirs) > 0 {
		for _, d := range dirs {
			hotspots = append(hotspots, d.Path+" "+d.SizeHuman)
		}
	}

	// ---- SERVICES ----
	services := []string{}
	if failed, _ := service.FailedServices(); len(failed) > 0 {
		for _, s := range failed {
			services = append(services, s.Name+" ("+s.State+")")
		}
	}

	networkState := "inactive"
	if nets, _ := network.ActiveInterfaces(); len(nets) > 0 {
		networkState = "active"
	}

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
