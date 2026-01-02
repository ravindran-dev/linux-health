package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	topName, topPID, topMB := "", 0, uint64(0)
	if topProc, _ := process.TopMemoryProcess(); topProc != nil {
		topName = topProc.Cmd
		topPID = topProc.PID
		topMB = topProc.MemKB / 1024
	}

	hotspot := ""
	if home := os.Getenv("HOME"); home != "" {
		dirs, _ := disk.Scan(
			filepath.Dir(home),
			map[string]bool{
				"/proc": true,
				"/sys":  true,
				"/dev":  true,
			},
		)
		if len(dirs) > 0 {
			hotspot = fmt.Sprintf(
				"%s (%d MB)",
				dirs[0].Path,
				dirs[0].Size/1024/1024,
			)
		}
	}

	netInfo := "inactive"
	if nets, _ := network.ActiveInterfaces(); len(nets) > 0 {
		netInfo = "active"
	}

	failedSvcs := 0
	if failed, err := service.FailedServices(); err == nil {
		failedSvcs = len(failed)
	}

	output.PrintBlock(
		stats,
		health,
		topName,
		topPID,
		topMB,
		hotspot,
		netInfo,
		failedSvcs,
	)

	os.Exit(health.ExitCode)
}
