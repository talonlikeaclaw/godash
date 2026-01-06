package ui

import "github.com/charmbracelet/lipgloss"

// getStatusStyle returns the appropriate style for a VM/container status
func getStatusStyle(status string) lipgloss.Style {
	if status == "running" {
		return runningStatusStyle
	}
	return stoppedStatusStyle
}

// getNodeStatusStyle returns the appropriate style for a node status
func getNodeStatusStyle(status string) lipgloss.Style {
	if status == "online" {
		return runningStatusStyle
	}
	return stoppedStatusStyle
}

// getCPUUsageStyle returns the appropriate style based on CPU usage percentage
func getCPUUsageStyle(cpuPercent float64) lipgloss.Style {
	if cpuPercent >= 80.0 {
		return criticalUsageStyle
	} else if cpuPercent >= 60.0 {
		return warningUsageStyle
	}
	return normalUsageStyle
}

// getMemoryUsageStyle returns the appropriate style based on memory usage percentage
func getMemoryUsageStyle(used, total int64) lipgloss.Style {
	if total == 0 {
		return normalUsageStyle
	}
	usagePercent := (float64(used) / float64(total)) * 100.0

	if usagePercent >= 85.0 {
		return criticalUsageStyle
	} else if usagePercent >= 70.0 {
		return warningUsageStyle
	}
	return normalUsageStyle
}

// getDiskUsageStyle returns the appropriate style based on disk usage percentage
func getDiskUsageStyle(used, total int64) lipgloss.Style {
	if total == 0 {
		return normalUsageStyle
	}
	usagePercent := (float64(used) / float64(total)) * 100.0

	if usagePercent >= 85.0 {
		return criticalUsageStyle
	} else if usagePercent >= 70.0 {
		return warningUsageStyle
	}
	return normalUsageStyle
}
