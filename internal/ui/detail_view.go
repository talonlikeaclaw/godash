package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/talonlikeaclaw/godash/internal/models"
)

// renderDetailView renders the detail view for a selected VM or container
func (m Model) renderDetailView() string {
	var b strings.Builder
	cpuProg := progress.New(progress.WithGradient(string(tokyoNightGreen), string(tokyoNightRed)))
	cpuProg.Width = 20

	memProg := progress.New(progress.WithGradient(string(tokyoNightGreen), string(tokyoNightRed)))
	memProg.Width = 20

	if m.selectedIsVM {
		var vm models.VM
		found := false
		for _, v := range m.vms {
			if v.ID == m.selectedID {
				vm = v
				found = true
				break
			}
		}
		if !found {
			b.WriteString(criticalUsageStyle.Render("Selected VM no longer available."))
			b.WriteString("\n")
			b.WriteString(helpStyle.Render("Controls: q to go back | Esc to quit"))
			return b.String()
		}

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
		fmt.Fprintf(&detailContent, "%s %s\n", nameLabel, vm.Name)

		// Status
		statusLabel := sectionHeaderStyle.Render("Status:")
		statusStyled := getStatusStyle(vm.Status).Render(vm.Status)
		fmt.Fprintf(&detailContent, "%s %s", statusLabel, statusStyled)

		// Uptime
		uptimeLabel := sectionHeaderStyle.Render("Uptime:")
		fmt.Fprintf(&detailContent, "%s %s", uptimeLabel, models.FormatUptime(vm.Uptime))

		// Resource usage
		detailContent.WriteString(sectionHeaderStyle.Render("Resource Usage:\n"))
		detailContent.WriteString("\n")

		// CPU with progress bar
		cpuPercent := vm.CPU / 100.0
		cpuBar := cpuProg.ViewAs(cpuPercent)
		cpuStyled := getCPUUsageStyle(vm.CPU).Render(fmt.Sprintf("%.1f%%", vm.CPU))
		fmt.Fprintf(&detailContent, "CPU:    %s %s\n", cpuBar, cpuStyled)

		// Memory with progress bar
		memUsed := models.BytesToGB(vm.MemUsed)
		memTotal := models.BytesToGB(vm.MemTotal)
		memPercent := (float64(vm.MemUsed) / float64(vm.MemTotal)) * 100.0
		memBar := memProg.ViewAs(memPercent / 100.0)
		memStyled := getMemoryUsageStyle(vm.MemUsed, vm.MemTotal).Render(fmt.Sprintf("%.1fGB/%.1fGB", memUsed, memTotal))
		fmt.Fprintf(&detailContent, "Memory: %s %s\n", memBar, memStyled)

		// Wrap in box
		b.WriteString(nodeBoxStyle.Render(detailContent.String()))
		b.WriteString("\n")

	} else {
		var container models.Container
		found := false
		for _, ct := range m.containers {
			if ct.ID == m.selectedID {
				container = ct
				found = true
				break
			}
		}
		if !found {
			b.WriteString(criticalUsageStyle.Render("Selected container no longer available."))
			b.WriteString("\n")
			b.WriteString(helpStyle.Render("Controls: q to go back | Esc to quit"))
			return b.String()
		}

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
		fmt.Fprintf(&detailContent, "%s %s\n", nameLabel, container.Name)

		// Status
		statusLabel := sectionHeaderStyle.Render("Status:")
		statusStyled := getStatusStyle(container.Status).Render(container.Status)
		fmt.Fprintf(&detailContent, "%s %s", statusLabel, statusStyled)

		// Uptime
		uptimeLabel := sectionHeaderStyle.Render("Uptime:")
		fmt.Fprintf(&detailContent, "%s %s", uptimeLabel, models.FormatUptime(container.Uptime))

		// Resource usage
		detailContent.WriteString(sectionHeaderStyle.Render("Resource Usage:\n"))
		detailContent.WriteString("\n")

		// CPU with progress bar
		cpuPercent := container.CPU / 100.0
		cpuBar := cpuProg.ViewAs(cpuPercent)
		cpuStyled := getCPUUsageStyle(container.CPU).Render(fmt.Sprintf("%.1f%%", container.CPU))
		fmt.Fprintf(&detailContent, "CPU:    %s %s\n", cpuBar, cpuStyled)

		// Memory with progress bar
		memUsed := models.BytesToGB(container.MemUsed)
		memTotal := models.BytesToGB(container.MemTotal)
		memPercent := (float64(container.MemUsed) / float64(container.MemTotal)) * 100.0
		memBar := memProg.ViewAs(memPercent / 100.0)
		memStyled := getMemoryUsageStyle(container.MemUsed, container.MemTotal).Render(fmt.Sprintf("%.1fGB/%.1fGB", memUsed, memTotal))
		fmt.Fprintf(&detailContent, "Memory: %s %s\n", memBar, memStyled)

		// Wrap in box
		b.WriteString(nodeBoxStyle.Render(detailContent.String()))
		b.WriteString("\n")
	}

	// Help text
	b.WriteString(helpStyle.Render("Controls: q to go back | Esc to quit"))
	b.WriteString("\n")

	return b.String()
}
