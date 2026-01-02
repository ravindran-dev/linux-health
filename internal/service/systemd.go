package service

import (
	"bytes"
	"os/exec"
)

type ServiceStatus struct {
	Name  string
	State string
}

func FailedServices() ([]ServiceStatus, error) {
	cmd := exec.Command("systemctl", "--failed", "--no-legend")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return []ServiceStatus{}, nil
}
