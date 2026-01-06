package ui

import "github.com/charmbracelet/lipgloss"

// View constants
const (
	ViewDashboard = "dashboard"
	ViewDetail    = "detail"
)

// Tokyo Night color palette
var (
	tokyoNightPurple = lipgloss.Color("#bb9af7")
	tokyoNightBlue   = lipgloss.Color("#7aa2f7")
	tokyoNightCyan   = lipgloss.Color("#7dcfff")
	tokyoNightGreen  = lipgloss.Color("#9ece6a")
	tokyoNightYellow = lipgloss.Color("#e0af68")
	tokyoNightRed    = lipgloss.Color("#f7768e")
	tokyoNightGray   = lipgloss.Color("#565f89")
)

// UI Styles
var (
	// Title styling
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(tokyoNightPurple)

	// Node header style
	nodeHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(tokyoNightYellow)

	// Section headers
	sectionHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(tokyoNightCyan).
				MarginTop(1)

	// Node info box
	nodeBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(tokyoNightPurple).
			Padding(0, 1).
			MarginBottom(1)

	// Selected item highlight
	selectedItemStyle = lipgloss.NewStyle().
				Background(tokyoNightBlue).
				Foreground(lipgloss.Color("#1a1b26")).
				Bold(true)

	// Status colors
	runningStatusStyle = lipgloss.NewStyle().
				Foreground(tokyoNightGreen).
				Bold(true)

	stoppedStatusStyle = lipgloss.NewStyle().
				Foreground(tokyoNightRed).
				Bold(true)

	// Usage level colors (for CPU/Memory/Disk)
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
