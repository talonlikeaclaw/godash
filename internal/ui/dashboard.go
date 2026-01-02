package ui

import (
	"fmt"

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
		}
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	s := "Godash - Homelab Dashboard\n\n"

	// Node info
	s += fmt.Sprintf("Node: %s (%s)\n", m.node.Name, m.node.Status)
	s += fmt.Sprintf("CPU: %.1f%% | Memory: %dGB / %dGB\n\n",
		m.node.CPU,
		m.node.MemUsed/(1024*1024*1024),
		m.node.MemTotal/(1024*1024*1024))

	// VM list
	s += "Virtual Machines:\n"
	for _, vm := range m.vms {
		s += fmt.Sprintf("  [%d] %s - %s\n", vm.ID, vm.Name, vm.Status)
	}

	s += "\nPress q to quit.\n"
	return s
}
