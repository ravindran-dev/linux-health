package output

import (
	"fmt"

	"github.com/ravindran-dev/linux-health/internal/system"
)

const boxWidth = 40
const inner = boxWidth - 2

func PrintBlock(
	stats system.SystemStats,
	health system.HealthReport,
	topName string,
	topPID int,
	topMB uint64,
	network string,
	failedServices []string,
	diskHotspots []string,
) {
	top()
	center("LINUX SYSTEM HEALTH")
	sep()

	row(fmt.Sprintf("SCORE   : %3d / 100  [ %-4s ]", health.Score, healthLabel(health)))
	row(fmt.Sprintf("UPTIME  : %s", stats.Uptime))

	sep()

	metric("CPU", stats.CPUUsage, health.CPU.Status, "%")
	metric("MEMORY", stats.MemUsagePct, health.Memory.Status, "%")
	metric("DISK", stats.DiskUsage, health.Disk.Status, "%")
	metricVal("LOAD", fmt.Sprintf("%.2f", stats.Load1), health.Load.Status)

	sep()

	if topName != "" {
		row(fmt.Sprintf("TOP MEM : %s", topName))
		row(fmt.Sprintf("          PID %-5d %4d MB", topPID, topMB))
	}

	sep()

	if len(diskHotspots) > 0 {
		row("HOTSPOTS:")
		for _, d := range diskHotspots {
			row("  " + trim(d))
		}
		sep()
	}

	row("SERVICES:")
	if len(failedServices) == 0 {
		row("  ✓ No failed services")
	} else {
		for _, s := range failedServices {
			row("  ● " + trim(s))
		}
	}

	sep()
	row(fmt.Sprintf("NETWORK : %s", network))

	bottom()
}

func top() {
	fmt.Printf("┌%s┐\n", repeat("─", inner))
}

func bottom() {
	fmt.Printf("└%s┘\n", repeat("─", inner))
}

func sep() {
	fmt.Printf("├%s┤\n", repeat("─", inner))
}

func center(t string) {
	fmt.Printf("│%s│\n", padCenter(t))
}

func row(t string) {
	fmt.Printf("│%-*s│\n", inner, t)
}

func metric(name string, val float64, status system.Status, unit string) {
	row(fmt.Sprintf("%-7s: %5.1f %s   [ %-2s ]", name, val, unit, status))
}

func metricVal(name, val string, status system.Status) {
	row(fmt.Sprintf("%-7s: %5s     [ %-2s ]", name, val, status))
}

func healthLabel(h system.HealthReport) string {
	switch h.ExitCode {
	case 2:
		return "CRIT"
	case 1:
		return "WARN"
	default:
		return "GOOD"
	}
}

func padCenter(s string) string {
	pad := inner - len(s)
	l := pad / 2
	r := pad - l
	return repeat(" ", l) + s + repeat(" ", r)
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}

func trim(s string) string {
	if len(s) > inner-4 {
		return s[:inner-7] + "..."
	}
	return s
}
