package ui

import (
	"fmt"
	"strings"

	"github.com/talonlikeaclaw/godash/internal/models"
)

// renderDetailView renders the detail view for a selected VM or container
func (m Model) renderDetailView() string {
	var b strings.Builder

	if m.cursor < len(m.vms) {
		// Rendering VM detail
		vm := m.vms[m.cursor]

		// Title
		title := fmt.Sprintf("Virtual Machine: %s", vm.Name)
		b.WriteString(titleStyle.Render(title))
		b.WriteString("\n\n")

		// Build detail content
		var detailContent strings.Builder

		// ID and Name
		idLabel := nodeHeaderStyle.Render("ID:")
		nameLabel := nodeHeaderStyle.Render("Name:")
		fmt.Fprintf(&detailContent, "%s %d\n", idLabel, vm.ID)
		fmt.Fprintf(&detailContent, "%s %s\n\n", nameLabel, vm.Name)

		// Status
		statusLabel := sectionHeaderStyle.Render("Status:")
		statusStyled := getStatusStyle(vm.Status).Render(vm.Status)
		fmt.Fprintf(&detailContent, "%s %s\n\n", statusLabel, statusStyled)

		// Resource usage
		detailContent.WriteString(sectionHeaderStyle.Render("Resource Usage"))
		detailContent.WriteString("\n")

		cpuStyled := getCPUUsageStyle(vm.CPU).Render(fmt.Sprintf("%.1f%%", vm.CPU))
		memUsed := models.BytesToGB(vm.MemUsed)
		memTotal := models.BytesToGB(vm.MemTotal)
		memStyled := getMemoryUsageStyle(vm.MemUsed, vm.MemTotal).Render(fmt.Sprintf("%.1fGB/%.1fGB", memUsed, memTotal))
		memPercent := (float64(vm.MemUsed) / float64(vm.MemTotal)) * 100.0

		fmt.Fprintf(&detailContent, "CPU:    %s\n", cpuStyled)
		fmt.Fprintf(&detailContent, "Memory: %s (%.1f%%)\n\n", memStyled, memPercent)

		// Uptime
		uptimeLabel := sectionHeaderStyle.Render("Uptime:")
		fmt.Fprintf(&detailContent, "%s %s\n", uptimeLabel, models.FormatUptime(vm.Uptime))

		// Wrap in box
		b.WriteString(nodeBoxStyle.Render(detailContent.String()))
		b.WriteString("\n")

	} else {
		// Rendering Container detail
		container := m.containers[m.cursor-len(m.vms)]

		// Title
		title := fmt.Sprintf("LXC Container: %s", container.Name)
		b.WriteString(titleStyle.Render(title))
		b.WriteString("\n\n")

		// Build detail content
		var detailContent strings.Builder

		// ID and Name
		idLabel := nodeHeaderStyle.Render("ID:")
		nameLabel := nodeHeaderStyle.Render("Name:")
		fmt.Fprintf(&detailContent, "%s %d\n", idLabel, container.ID)
		fmt.Fprintf(&detailContent, "%s %s\n\n", nameLabel, container.Name)

		// Status
		statusLabel := sectionHeaderStyle.Render("Status:")
		statusStyled := getStatusStyle(container.Status).Render(container.Status)
		fmt.Fprintf(&detailContent, "%s %s\n\n", statusLabel, statusStyled)

		// Resource usage
		detailContent.WriteString(sectionHeaderStyle.Render("Resource Usage"))
		detailContent.WriteString("\n")

		cpuStyled := getCPUUsageStyle(container.CPU).Render(fmt.Sprintf("%.1f%%", container.CPU))
		memUsed := models.BytesToGB(container.MemUsed)
		memTotal := models.BytesToGB(container.MemTotal)
		memStyled := getMemoryUsageStyle(container.MemUsed, container.MemTotal).Render(fmt.Sprintf("%.1fGB/%.1fGB", memUsed, memTotal))
		memPercent := (float64(container.MemUsed) / float64(container.MemTotal)) * 100.0

		fmt.Fprintf(&detailContent, "CPU:    %s\n", cpuStyled)
		fmt.Fprintf(&detailContent, "Memory: %s (%.1f%%)\n\n", memStyled, memPercent)

		// Uptime
		uptimeLabel := sectionHeaderStyle.Render("Uptime:")
		fmt.Fprintf(&detailContent, "%s %s\n", uptimeLabel, models.FormatUptime(container.Uptime))

		// Wrap in box
		b.WriteString(nodeBoxStyle.Render(detailContent.String()))
		b.WriteString("\n")
	}

	// Help text
	b.WriteString(helpStyle.Render("Controls: Esc to go back | q to quit"))
	b.WriteString("\n")

	return b.String()
}
