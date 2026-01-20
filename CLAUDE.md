# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Goal:** Build a TUI (Terminal User Interface) using Bubbletea to monitor Proxmox homelab infrastructure (VMs, containers, and system resources) in real-time.

**Current Phase:** Phase 2 - Integrating real Proxmox API client.

**Tech Stack:**
- Go 1.25.5
- Bubbletea (TUI framework based on Elm Architecture)
- Lipgloss (styling and layout for terminal UIs)
- Development in distrobox container (golang:1.25 image)
- Module path: `github.com/talonlikeaclaw/godash`

## Development Environment

**All development happens in distrobox container.** The container is named `godash-dev` by default.

```bash
# Set up development environment (first time)
./scripts/setup-dev-env.sh

# Enter the container
distrobox enter godash-dev
cd ~/projects/godash

# With custom container name
./scripts/setup-dev-env.sh my-custom-name
distrobox enter my-custom-name
```

## Common Commands

### Building and Running
```bash
# Run the application directly
go run cmd/godash/main.go
# Or use make
make run

# Build binary
make build              # Standard build -> bin/godash
make dev                # Development build with race detector -> bin/godash-dev
make prod               # Optimized production build -> bin/godash

# Clean build artifacts
make clean
```

### Testing
```bash
# Run all tests
make test               # Verbose output
go test ./...           # Standard output

# Run with coverage
make coverage

# Run specific test
go test -v ./internal/ui -run TestSpecificTest
```

### Code Quality
```bash
# Run linter (requires golangci-lint)
make lint

# Manage dependencies
make deps               # Download and tidy dependencies
go mod tidy             # Tidy dependencies (run before commits)
```

## Architecture

### Project Structure
```
godash/
├── cmd/godash/          # Main application entry point
├── internal/
│   ├── models/          # Data models (Node, VM, Container) + fake data generators
│   ├── ui/              # Bubbletea UI components (dashboard.go)
│   ├── app/             # Application state (not implemented yet)
│   ├── config/          # Config loading from YAML
│   └── proxmox/         # Proxmox API client (in progress)
├── configs/             # Example YAML config (config.example.yaml)
└── scripts/             # Dev environment setup scripts
```

**Key Files:**
- `cmd/godash/main.go` - Application entry point, initializes Bubble Tea program
- `internal/models/models.go` - Core data structures (Node, VM, Container)
- `internal/models/fake_data.go` - Fake data generators for Phase 1 testing
- `internal/ui/dashboard.go` - Main dashboard UI component
- `internal/config/config.go` - YAML config loading and structs
- `configs/config.example.yaml` - Example configuration file

### Bubble Tea Pattern (Elm Architecture)

The application follows the Elm Architecture pattern via Bubble Tea:

1. **Model** (`internal/ui/dashboard.go`): Holds all application state including:
   - Node information
   - VMs and containers lists
   - UI state (cursor position, quit flag)

2. **Init()**: Called once at startup to return initial commands

3. **Update(msg tea.Msg)**: Handles all state changes based on messages:
   - Keyboard events (q/Ctrl+C to quit, up/down or j/k for navigation)
   - Future: view switching, data updates, API responses
   - Returns updated model and optional command

4. **View()**: Pure rendering function that returns a string representation of current state

5. **Messages**: Communication between components (keypresses, API responses, timers)

6. **Commands**: Async operations that will eventually produce messages (API calls, timers, I/O)

### Data Flow

Currently uses fake data generators in `internal/models/fake_data.go`:
- `GetFakeNode()` - Returns fake Proxmox node data
- `GetFakeVMs()` - Returns slice of fake VMs
- `GetFakeContainers()` - Returns slice of fake containers

Future phases will replace these with real Proxmox API calls.

### Key Design Patterns

- **UI State Centralization**: All UI state lives in the `ui.Model` struct
- **Message-Driven Updates**: State changes only occur via the Update() function responding to messages
- **Separation of Concerns**: Models define data structures, UI handles presentation, cmd/ handles initialization
- **Fake Data for Development**: `fake_data.go` provides realistic test data to develop UI without requiring Proxmox connection

## Current State

**What works (Phase 1 complete):**
- Dashboard TUI displaying node info (CPU, Memory, Disk usage)
- VM and LXC container lists with unified cursor navigation
- Keyboard navigation (up/down or j/k, with wrapping)
- Quit functionality (press 'q' or Ctrl+C)
- Project structure and build system
- Fake data generators for development without Proxmox connection
- BytesToGB helper for memory/disk display formatting
- Lipgloss styling throughout the UI
- Multiple views (dashboard, VM detail, container detail)
- Config loading from YAML with unit tests

**What's next (Phase 2):**
- Build Proxmox API client (`internal/proxmox/client.go`)
- Implement GetNodeStatus, GetVMs, GetContainers methods
- Wire client into Bubble Tea with periodic refresh using tea.Tick
- Replace fake data with real Proxmox API data

## Future Phases

- **Phase 3**: Add Docker integration via SSH to VMs
- **Phase 4**: Smart update checking with GitHub API integration for changelog awareness

## Important Notes

- **Fake Data**: Using anonymous/fake data in `fake_data.go`, not real homelab data
- **Development Approach**: Building UI-first with fake data; APIs come later
- **Go Practices**: Following idiomatic Go (gofmt, conventional naming, defer adding dependencies until needed)
- **Branching Strategy**: Feature branches off main, merge when complete
- **Dependencies**: Only add new dependencies when actually needed, not speculatively

## Coding Guidelines

When working on this project:

1. **Follow Elm Architecture**: All state changes through Update(), pure View() function
2. **Idiomatic Go**: Use gofmt, follow Go naming conventions, keep it simple
3. **Incremental Changes**: Make small steps, test often
4. **Guidance Over Implementation**: Provide guidance and explanations; let the user implement the code themselves unless explicitly asked to write it
5. **Write Unit Tests**: Add tests for new functionality, especially API client methods
