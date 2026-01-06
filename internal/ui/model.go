package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/talonlikeaclaw/godash/internal/models"
)

// Model represents the application state
type Model struct {
	node        models.Node
	vms         []models.VM
	containers  []models.Container
	cursor      int    // which VM or container is selected
	currentView string // current view (dashboard or detail)
	quitting    bool
}

// NewModel creates a new model with fake data
func NewModel() Model {
	return Model{
		node:        models.GetFakeNode(),
		vms:         models.GetFakeVMs(),
		containers:  models.GetFakeContainers(),
		currentView: ViewDashboard,
		cursor:      0,
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
			totalItems := len(m.vms) + len(m.containers)
			if totalItems > 0 && m.currentView == ViewDashboard {
				// Move cursor up
				m.cursor = (m.cursor - 1 + totalItems) % totalItems
			}

		case "down", "j":
			totalItems := len(m.vms) + len(m.containers)
			if totalItems > 0 && m.currentView == ViewDashboard {
				// Move cursor down
				m.cursor = (m.cursor + 1) % totalItems
			}

		case "enter":
			if m.currentView == ViewDashboard {
				m.currentView = ViewDetail
			}

		case "esc":
			if m.currentView == ViewDetail {
				m.currentView = ViewDashboard
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

	if m.currentView == ViewDetail {
		return m.renderDetailView()
	}

	return m.renderDashboardView()
}
