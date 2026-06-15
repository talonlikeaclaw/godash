package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/talonlikeaclaw/godash/internal/config"
	"github.com/talonlikeaclaw/godash/internal/proxmox"
	"github.com/talonlikeaclaw/godash/internal/ui"
)

func main() {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "godash", "config.yaml")

	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load config (%v) - using fake data\n", err)
		runProgram(ui.NewModelWithFakeData())
		return
	}

	client := proxmox.New(cfg)
	runProgram(ui.NewModel(client, cfg.Refresh.NodeStats))
}

func runProgram(m ui.Model) {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
