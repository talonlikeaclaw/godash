package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/talonlikeaclaw/godash/internal/models"
)

// Model represents the application state
type Model struct {
	node        models.Node
	vms         []models.VM
	containers  []models.Container
	cursor      int // which VM or container is selected
	currentView string
	quitting    bool
}

// NewModel creates a new model with fake data
func NewModel() Model {
	return Model{
		node:        models.GetFakeNode(),
		vms:         models.GetFakeVMs(),
		containers:  models.GetFakeContainers(),
		currentView: "dashboard",
		cursor:      0,
	}
}

var (
	// Tokyo Night color palette
	tokyoNightPurple = lipgloss.Color("#bb9af7")
	tokyoNightBlue   = lipgloss.Color("#7aa2f7")
	tokyoNightCyan   = lipgloss.Color("#7dcfff")
	tokyoNightGreen  = lipgloss.Color("#9ece6a")
	tokyoNightYellow = lipgloss.Color("#e0af68")
	tokyoNightRed    = lipgloss.Color("#f7768e")
	tokyoNightGray   = lipgloss.Color("#565f89")

	// Title styling
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(tokyoNightPurple).
			Underline(true)

	// Node header style
	nodeHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(tokyoNightYellow)

	// Section headers
	sectionHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(tokyoNightBlue).
				MarginTop(1)

	// Node info box
	nodeBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(tokyoNightPurple).
			Padding(0, 1).
			MarginBottom(1)

	// Selected item highlight
	selectedItemStyle = lipgloss.NewStyle().
				Background(tokyoNightCyan).
				Foreground(lipgloss.Color("#1a1b26")).
				Bold(true)

	// Status colors
	runningStatusStyle = lipgloss.NewStyle().
				Foreground(tokyoNightGreen).
				Bold(true)

	stoppedStatusStyle = lipgloss.NewStyle().
				Foreground(tokyoNightRed).
				Bold(true)

	// Usage level colors (for CPU/Memory)
	normalUsageStyle = lipgloss.NewStyle().
				Foreground(tokyoNightGreen)

	warningUsageStyle = lipgloss.NewStyle().
				Foreground(tokyoNightYellow)

	criticalUsageStyle = lipgloss.NewStyle().
				Foreground(tokyoNightRed)

	// Help text
	helpStyle = lipgloss.NewStyle().
			Foreground(tokyoNightGray).
			MarginTop(1)
)

// Init is called once at startup
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles all state changes
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			totalItems := len(m.vms) + len(m.containers)
			if totalItems > 0 {
				// Move cursor up
				m.cursor = (m.cursor - 1 + totalItems) % totalItems
			}

		case "down", "j":
			totalItems := len(m.vms) + len(m.containers)
			if totalItems > 0 {
				// Move cursor down
				m.cursor = (m.cursor + 1) % totalItems
			}

		case "enter":
			if m.currentView == "dashboard" {
				m.currentView = "detail"
			}

		case "esc":
			if m.currentView == "detail" {
				m.currentView = "dashboard"
			}
		}
	}

	return m, nil
}

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

// View renders the UI
func (m Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Godash - Homelab Dashboard"))
	b.WriteString("\n\n")

	// Build content for the box
	var boxContent strings.Builder

	// Node info with styling
	nodeLabel := nodeHeaderStyle.Render("Node:")
	nodeName := nodeHeaderStyle.Render(m.node.Name)
	nodeStatus := getNodeStatusStyle(m.node.Status).Render(m.node.Status)
	fmt.Fprintf(&boxContent, "%s %s (%s) | Uptime: %s\n", nodeLabel, nodeName, nodeStatus, models.FormatUptime(m.node.Uptime))

	// Style CPU, Memory, and Disk based on usage
	cpuStyled := getCPUUsageStyle(m.node.CPU).Render(fmt.Sprintf("%.1f%%", m.node.CPU))
	memUsedGB := models.BytesToGB(m.node.MemUsed)
	memTotalGB := models.BytesToGB(m.node.MemTotal)
	memStyled := getMemoryUsageStyle(m.node.MemUsed, m.node.MemTotal).Render(fmt.Sprintf("%.0fGB/%.0fGB", memUsedGB, memTotalGB))
	diskUsedGB := models.BytesToGB(m.node.DiskUsed)
	diskTotalGB := models.BytesToGB(m.node.DiskTotal)
	diskStyled := getDiskUsageStyle(m.node.DiskUsed, m.node.DiskTotal).Render(fmt.Sprintf("%.0fGB/%.0fGB", diskUsedGB, diskTotalGB))

	fmt.Fprintf(&boxContent, "CPU: %s | Memory: %s | Disk: %s\n\n",
		cpuStyled,
		memStyled,
		diskStyled)

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
	b.WriteString(helpStyle.Render("Controls: ↑/↓ or j/k to navigate | q to quit"))
	b.WriteString("\n")

	return b.String()
}
