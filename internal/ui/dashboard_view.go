package ui

import (
	"fmt"
	"strings"

	"github.com/talonlikeaclaw/godash/internal/models"
)

// renderDashboardView renders the main dashboard view
func (m Model) renderDashboardView() string {
	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Godash - Homelab Dashboard"))
	b.WriteString("\n\n")

	// Show API error if present
	if m.lastErr != nil {
		b.WriteString(criticalUsageStyle.Render(fmt.Sprintf("API error: %v", m.lastErr)))
		b.WriteString("\n\n")
	}

	// Build content for the box
	var boxContent strings.Builder

	// Node info with styling
	nodeLabel := nodeHeaderStyle.Render("Node:")
	nodeName := nodeHeaderStyle.Render(m.node.Name)
	nodeStatus := getNodeStatusStyle(m.node.Status).Render(m.node.Status)
	fmt.Fprintf(&boxContent, "%s %s (%s) | Uptime: %s\n\n", nodeLabel, nodeName, nodeStatus, models.FormatUptime(m.node.Uptime))

	// CPU
	cpuPercent := m.node.CPU / 100.0
	cpuBar := m.cpuProgress.ViewAs(cpuPercent)
	cpuValue := getCPUUsageStyle(m.node.CPU).Render(fmt.Sprintf("%.1f%%", m.node.CPU))
	fmt.Fprintf(&boxContent, "CPU:    %s %s\n", cpuBar, cpuValue)

	// Memory
	memUsedGB := models.BytesToGB(m.node.MemUsed)
	memTotalGB := models.BytesToGB(m.node.MemTotal)
	memPercent := float64(m.node.MemUsed) / float64(m.node.MemTotal)
	memBar := m.memProgress.ViewAs(memPercent)
	memValue := getMemoryUsageStyle(m.node.MemUsed, m.node.MemTotal).Render(fmt.Sprintf("%.0fGB/%.0fGB", memUsedGB, memTotalGB))
	fmt.Fprintf(&boxContent, "Memory: %s %s\n", memBar, memValue)

	// Disk
	diskUsedGB := models.BytesToGB(m.node.DiskUsed)
	diskTotalGB := models.BytesToGB(m.node.DiskTotal)
	diskPercent := float64(m.node.DiskUsed) / float64(m.node.DiskTotal)
	diskBar := m.diskProgress.ViewAs(diskPercent)
	diskValue := getDiskUsageStyle(m.node.DiskUsed, m.node.DiskTotal).Render(fmt.Sprintf("%.0fGB/%.0fGB", diskUsedGB, diskTotalGB))
	fmt.Fprintf(&boxContent, "Disk:   %s %s\n", diskBar, diskValue)

	// VM list with cursor
	boxContent.WriteString(sectionHeaderStyle.Render("Virtual Machines"))
	boxContent.WriteString("\n")
	for i, vm := range m.vms {
		// Build the line content
		line := fmt.Sprintf(" [%d] %s - %s (CPU: %.1f%%, Mem: %.1fGB/%.1fGB, Up: %s)",
			vm.ID,
			vm.Name,
			vm.Status,
			vm.CPU,
			models.BytesToGB(vm.MemUsed),
			models.BytesToGB(vm.MemTotal),
			models.FormatUptime(vm.Uptime))

		// Apply highlighting if selected
		if m.cursor == i {
			boxContent.WriteString(selectedItemStyle.Render(line))
		} else {
			// Color the status, CPU, and memory
			lineParts := fmt.Sprintf(" [%d] %s - ", vm.ID, vm.Name)
			statusStyled := getStatusStyle(vm.Status).Render(vm.Status)
			cpuStyled := getCPUUsageStyle(vm.CPU).Render(fmt.Sprintf("%.1f%%", vm.CPU))
			memUsed := models.BytesToGB(vm.MemUsed)
			memTotal := models.BytesToGB(vm.MemTotal)
			memStyled := getMemoryUsageStyle(vm.MemUsed, vm.MemTotal).Render(fmt.Sprintf("%.1fGB/%.1fGB", memUsed, memTotal))

			fmt.Fprintf(&boxContent, "%s%s (CPU: %s, Mem: %s, Up: %s)",
				lineParts,
				statusStyled,
				cpuStyled,
				memStyled,
				models.FormatUptime(vm.Uptime))
		}
		boxContent.WriteString("\n")
	}

	// Container list with cursor
	boxContent.WriteString("\n")
	boxContent.WriteString(sectionHeaderStyle.Render("LXC Containers"))
	boxContent.WriteString("\n")
	for i, container := range m.containers {
		// Build the line content
		line := fmt.Sprintf(" [%d] %s - %s (CPU: %.1f%%, Mem: %.1fGB/%.1fGB, Up: %s)",
			container.ID,
			container.Name,
			container.Status,
			container.CPU,
			models.BytesToGB(container.MemUsed),
			models.BytesToGB(container.MemTotal),
			models.FormatUptime(container.Uptime))

		// Apply highlighting if selected
		if m.cursor == i+len(m.vms) {
			boxContent.WriteString(selectedItemStyle.Render(line))
		} else {
			// Color the status, CPU, and memory
			lineParts := fmt.Sprintf(" [%d] %s - ", container.ID, container.Name)
			statusStyled := getStatusStyle(container.Status).Render(container.Status)
			cpuStyled := getCPUUsageStyle(container.CPU).Render(fmt.Sprintf("%.1f%%", container.CPU))
			memUsed := models.BytesToGB(container.MemUsed)
			memTotal := models.BytesToGB(container.MemTotal)
			memStyled := getMemoryUsageStyle(container.MemUsed, container.MemTotal).Render(fmt.Sprintf("%.1fGB/%.1fGB", memUsed, memTotal))

			fmt.Fprintf(&boxContent, "%s%s (CPU: %s, Mem: %s, Up: %s)",
				lineParts,
				statusStyled,
				cpuStyled,
				memStyled,
				models.FormatUptime(container.Uptime))
		}
		boxContent.WriteString("\n")
	}

	// Wrap everything in the box
	b.WriteString(nodeBoxStyle.Render(boxContent.String()))
	b.WriteString("\n")

	// Help text (outside the box)
	b.WriteString(helpStyle.Render("Controls: ↑/↓ or j/k to navigate | Enter to view details | Esc to quit"))
	b.WriteString("\n")

	return b.String()
}
