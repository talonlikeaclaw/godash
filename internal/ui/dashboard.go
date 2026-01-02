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
