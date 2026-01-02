package main

import (
	"fmt"
	"log"

	"github.com/ravindran-dev/linux-health/internal/system"
)

func main() {
	stats, err := system.GetSystemStats()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LINUX SYSTEM HEALTH")
	fmt.Println("-------------------")
	fmt.Printf("CPU Usage     : %.2f%%\n", stats.CPUUsage)
	fmt.Printf("Memory Usage  : %dMB / %dMB (%.2f%%)\n",
		stats.MemUsedMB, stats.MemTotalMB, stats.MemUsagePct)
	fmt.Printf("Load Average  : %.2f %.2f %.2f\n",
		stats.Load1, stats.Load5, stats.Load15)
	fmt.Printf("Disk Usage    : %.2f%%\n", stats.DiskUsage)
	fmt.Printf("Uptime        : %s\n", stats.Uptime)
}
