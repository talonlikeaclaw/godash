package models

import "fmt"

type Node struct {
	Name      string
	Status    string
	CPU       float64
	CPUCores  int
	MemUsed   int64
	MemTotal  int64
	DiskUsed  int64
	DiskTotal int64
	Uptime    int64
}

type VM struct {
	ID       int
	Name     string
	Status   string
	CPU      float64
	MemUsed  int64
	MemTotal int64
	Uptime   int64
}

type Container struct {
	ID       int
	Name     string
	Status   string
	CPU      float64
	MemUsed  int64
	MemTotal int64
	Uptime   int64
}

// BytesToGB converts bytes to gigabytes with decimal precision
func BytesToGB(numBytes int64) float64 {
	return float64(numBytes) / (1024 * 1024 * 1024)
}

// FormatUptime converts seconds to a human-readable uptime string
func FormatUptime(seconds int64) string {
	if seconds == 0 {
		return "0s"
	}

	days := seconds / 86400
	hours := (seconds % 86400) / 3600
	minutes := (seconds % 3600) / 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh", days, hours)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}
