package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/talonlikeaclaw/godash/internal/models"
	"github.com/talonlikeaclaw/godash/internal/proxmox"
)

// dataRefreshMsg carries fresh API data (or an error) back to Update.
type dataRefreshMsg struct {
	node       models.Node
	vms        []models.VM
	containers []models.Container
	err        error
}

type tickMsg time.Time

// Model represents the application state.
type Model struct {
	node        models.Node
	vms         []models.VM
	containers  []models.Container
	cursor      int
	currentView string
	quitting    bool
	lastErr     error

	client      *proxmox.Client
	refreshSecs int

	cpuProgress  progress.Model
	memProgress  progress.Model
	diskProgress progress.Model
}

func newProgressBars() (progress.Model, progress.Model, progress.Model) {
	cpu := progress.New(progress.WithGradient(string(tokyoNightGreen), string(tokyoNightRed)))
	cpu.Width = 50
	mem := progress.New(progress.WithGradient(string(tokyoNightGreen), string(tokyoNightRed)))
	mem.Width = 50
	disk := progress.New(progress.WithGradient(string(tokyoNightGreen), string(tokyoNightRed)))
	disk.Width = 50
	return cpu, mem, disk
}

// NewModel creates a model backed by a real Proxmox client.
func NewModel(client *proxmox.Client, refreshSecs int) Model {
	cpu, mem, disk := newProgressBars()
	return Model{
		currentView:  ViewDashboard,
		client:       client,
		refreshSecs:  refreshSecs,
		cpuProgress:  cpu,
		memProgress:  mem,
		diskProgress: disk,
	}
}

// NewModelWithFakeData creates a model using fake data for development and testing.
func NewModelWithFakeData() Model {
	cpu, mem, disk := newProgressBars()
	return Model{
		node:         models.GetFakeNode(),
		vms:          models.GetFakeVMs(),
		containers:   models.GetFakeContainers(),
		currentView:  ViewDashboard,
		refreshSecs:  10,
		cpuProgress:  cpu,
		memProgress:  mem,
		diskProgress: disk,
	}
}

// Init fires the first data fetch when the program starts.
func (m Model) Init() tea.Cmd {
	if m.client == nil {
		return nil
	}
	return m.fetchData()
}

func (m Model) fetchData() tea.Cmd {
	return func() tea.Msg {
		node, err := m.client.GetNodeStatus()
		if err != nil {
			return dataRefreshMsg{err: err}
		}
		vms, err := m.client.GetVMs()
		if err != nil {
			return dataRefreshMsg{err: err}
		}
		containers, err := m.client.GetContainers()
		if err != nil {
			return dataRefreshMsg{err: err}
		}
		return dataRefreshMsg{node: node, vms: vms, containers: containers}
	}
}

func (m Model) tickCmd() tea.Cmd {
	return tea.Tick(time.Duration(m.refreshSecs)*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles all state changes.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dataRefreshMsg:
		if msg.err != nil {
			m.lastErr = msg.err
			return m, m.tickCmd()
		}
		m.lastErr = nil
		m.node = msg.node
		m.vms = msg.vms
		m.containers = msg.containers
		return m, m.tickCmd()

	case tickMsg:
		if m.client == nil {
			return m, nil
		}
		return m, m.fetchData()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			totalItems := len(m.vms) + len(m.containers)
			if totalItems > 0 && m.currentView == ViewDashboard {
				m.cursor = (m.cursor - 1 + totalItems) % totalItems
			}

		case "down", "j":
			totalItems := len(m.vms) + len(m.containers)
			if totalItems > 0 && m.currentView == ViewDashboard {
				m.cursor = (m.cursor + 1) % totalItems
			}

		case "enter":
			if m.currentView == ViewDashboard {
				m.currentView = ViewDetail
			}

		case "q":
			if m.currentView == ViewDetail {
				m.currentView = ViewDashboard
			}
		}
	}

	return m, nil
}

// View renders the UI.
func (m Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	if m.currentView == ViewDetail {
		return m.renderDetailView()
	}

	return m.renderDashboardView()
}
