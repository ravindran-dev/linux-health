package network

import (
	"os"
	"path/filepath"
	"strings"
)

type InterfaceState struct {
	Name  string
	State string
}

func ActiveInterfaces() ([]InterfaceState, error) {
	entries, err := os.ReadDir("/sys/class/net")
	if err != nil {
		return nil, err
	}

	var active []InterfaceState

	for _, e := range entries {
		iface := e.Name()

		// ignore loopback
		if iface == "lo" {
			continue
		}

		stateFile := filepath.Join("/sys/class/net", iface, "operstate")
		data, err := os.ReadFile(stateFile)
		if err != nil {
			continue
		}

		state := strings.TrimSpace(string(data))
		if state == "up" {
			active = append(active, InterfaceState{
				Name:  iface,
				State: state,
			})
		}
	}

	return active, nil
}
