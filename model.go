package main

import (
	"os/exec"
	"strings"
)

type device struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getDevices() ([]device, error) {
	devices := []device{}

	cmd := "ifconfig -a | sed 's/[ \t].*//;/^$/d'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	// cmd := "wmic nic get NetConnectionID"
	// out, err := exec.Command("powershell", "-c", cmd).Output()

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	for i, line := range lines {
		var dev device
		dev.ID = i
		dev.Name = line
		devices = append(devices, dev)
	}

	return devices, nil
}

func (dev *device) getDevice() error {
	out, err := exec.Command("bash", "-c", "ifconfig -a "+dev.Name).Output()
	if err != nil {
		return err
	}

	dev.Description = string(out)

	return nil
}
