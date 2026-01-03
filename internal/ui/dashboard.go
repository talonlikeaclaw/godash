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
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			// Move cursor down
			if m.cursor < len(m.vms)-1 {
				m.cursor++
			}
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
	fmt.Fprintf(&b, "CPU: %.1f%% | Memory: %dGB / %dGB\n\n",
		m.node.CPU,
		m.node.MemUsed/(1024*1024*1024),
		m.node.MemTotal/(1024*1024*1024))

	// VM list with cursor
	b.WriteString("Virtual Machines:\n")
	for i, vm := range m.vms {
		// Cursor indicator
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s [%d] %s - %s (CPU: %.1f%%, Mem: %dGB/%dGB)\n",
			cursor,
			vm.ID,
			vm.Name,
			vm.Status,
			vm.CPU,
			vm.MemUsed/(1024*1024*1024),
			vm.MemTotal/(1024*1024*1024))
	}

	b.WriteString("\nControls: ↑/↓ or j/k to navigate | q to quit\n")

	return b.String()
}
