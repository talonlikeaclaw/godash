package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/talonlikeaclaw/godash/internal/models"
)

// Model represents the application state
type Model struct {
	node       models.Node
	vms        []models.VM
	containers []models.Container
	cursor     int // which VM is selected
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
			// Move cursor up
			m.cursor = (m.cursor - 1 + len(m.vms)) % len(m.vms)

		case "down", "j":
			// Move cursor down
			m.cursor = (m.cursor + 1) % len(m.vms)
		}
	}

	return m, nil
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
	fmt.Fprintf(&b, "CPU: %.1f%% | Memory: %.1fGB / %.1fGB\n\n",
		m.node.CPU,
		models.BytesToGB(m.node.MemUsed),
		models.BytesToGB(m.node.MemTotal))

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

	b.WriteString("\nControls: ↑/↓ or j/k to navigate | q to quit\n")

	return b.String()
}
