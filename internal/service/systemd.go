package service

import (
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Name  string
	State string
}

var (
	cacheMu     sync.Mutex
	cached      []Service
	lastChecked time.Time
	cacheTTL    = 10 * time.Second
)

func FailedServices() ([]Service, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	if time.Since(lastChecked) < cacheTTL {
		return cached, nil
	}

	cmd := exec.Command("systemctl", "--failed", "--no-legend", "--plain")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	var failed []Service

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			failed = append(failed, Service{
				Name:  fields[0],
				State: fields[3],
			})
		}
	}

	// ✅ update cache
	cached = failed
	lastChecked = time.Now()

	return failed, nil
}
