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
	node       models.Node
	vms        []models.VM
	containers []models.Container
	cursor     int // which VM or container is selected
	quitting   bool
}

// NewModel creates a new model with fake data
func NewModel() Model {
	return Model{
		node:       models.GetFakeNode(),
		vms:        models.GetFakeVMs(),
		containers: models.GetFakeContainers(),
		cursor:     0,
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

// View renders the UI
func (m Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	var b strings.Builder

	b.WriteString("Godash - Homelab Dashboard\n\n")

	// Node info
	fmt.Fprintf(&b, "Node: %s (%s)\n", m.node.Name, m.node.Status)
	fmt.Fprintf(&b, "CPU: %.1f%% | Memory: %.0fGB / %.0fGB | Disk: %.0fGB / %.0fGB\n\n",
		m.node.CPU,
		models.BytesToGB(m.node.MemUsed),
		models.BytesToGB(m.node.MemTotal),
		models.BytesToGB(m.node.DiskUsed),
		models.BytesToGB(m.node.DiskTotal))

	// VM list with cursor
	b.WriteString("Virtual Machines:\n")
	for i, vm := range m.vms {
		// Cursor indicator
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s [%d] %s - %s (CPU: %.1f%%, Mem: %.1fGB/%.1fGB)\n",
			cursor,
			vm.ID,
			vm.Name,
			vm.Status,
			vm.CPU,
			models.BytesToGB(vm.MemUsed),
			models.BytesToGB(vm.MemTotal))
	}

	// Container list with cursor
	b.WriteString("\nLXC Containers:\n")
	for i, container := range m.containers {
		// Cursor indicator
		cursor := " "
		if m.cursor == i+len(m.vms) {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s [%d] %s - %s (CPU: %.1f%%, Mem: %.1fGB/%.1fGB)\n",
			cursor,
			container.ID,
			container.Name,
			container.Status,
			container.CPU,
			models.BytesToGB(container.MemUsed),
			models.BytesToGB(container.MemTotal))
	}

	b.WriteString("\nControls: ↑/↓ or j/k to navigate | q to quit\n")

	return b.String()
}
