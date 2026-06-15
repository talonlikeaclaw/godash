# CLAUDE.md

Guidance for Claude Code (claude.ai/code) when working in this repo.

## Project Overview

**Goal:** Build TUI (Terminal User Interface) using Bubbletea to monitor Proxmox homelab infrastructure (VMs, containers, system resources) in real-time.

**Current Phase:** Phase 3 - Proxmox control actions (start/stop/restart VMs and LXCs).

**Tech Stack:**
- Go 1.25.5
- Bubbletea (TUI framework based on Elm Architecture)
- Lipgloss (styling and layout for terminal UIs)
- Development in distrobox container (golang:1.25 image)
- Module path: `github.com/talonlikeaclaw/godash`

## Development Environment

**All development happens in distrobox container.** Container named `godash-dev` by default.

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
│   ├── ui/              # Bubbletea UI components
│   ├── app/             # Application state (not implemented yet)
│   ├── config/          # Config loading from YAML
│   └── proxmox/         # Proxmox API client
├── configs/             # Example YAML config (config.example.yaml)
└── scripts/             # Dev environment setup scripts
```

**Key Files:**
- `cmd/godash/main.go` - Entry point; loads config, creates proxmox client, falls back to fake data
- `internal/models/models.go` - Core data structures (Node, VM, Container)
- `internal/models/fake_data.go` - Fake data generators for offline dev/testing
- `internal/proxmox/client.go` - Proxmox HTTP client (auth, GetNodeStatus, GetVMs, GetContainers)
- `internal/ui/model.go` - Bubble Tea model; dataRefreshMsg/tickMsg wiring, async refresh
- `internal/ui/dashboard_view.go` - Dashboard rendering
- `internal/ui/detail_view.go` - VM/container detail rendering (ID-based, not cursor-based)
- `internal/config/config.go` - YAML config loading and structs
- `configs/config.example.yaml` - Example configuration file

### Bubble Tea Pattern (Elm Architecture)

App follows Elm Architecture via Bubble Tea:

1. **Model** (`internal/ui/model.go`): All app state — node, VMs, containers, cursor, selected item identity (`selectedID`/`selectedIsVM`), client, error state.

2. **Init()**: Fires first `fetchData()` command on startup.

3. **Update(msg tea.Msg)**: All state changes:
   - `dataRefreshMsg` — updates node/VM/container data, re-syncs cursor to tracked ID, schedules next tick
   - `tickMsg` — fires next `fetchData()` command
   - Keyboard events (navigation, view switching, quit)

4. **View()**: Pure rendering — delegates to `renderDashboardView()` or `renderDetailView()`.

5. **Async refresh**: `fetchData()` returns command calling all three Proxmox API methods concurrently-safe (runs off main goroutine). `tea.Tick` drives periodic re-fetch.

### Data Flow

```
main.go → config.Load() → proxmox.New(cfg) → ui.NewModel(client, refreshSecs)
    → Init() → fetchData() cmd
    → [API calls] → dataRefreshMsg
    → Update() → model updated → tickCmd()
    → [N seconds] → tickMsg → fetchData() → ...
```

Fallback: config missing → `ui.NewModelWithFakeData()` used (fake data generators in `internal/models/fake_data.go`).

### Key Design Patterns

- **UI State Centralization**: All UI state in `ui.Model` struct
- **Message-Driven Updates**: State changes only via Update() responding to messages
- **Separation of Concerns**: Models define data, UI handles presentation, cmd/ handles init
- **Fake Data for Development**: `fake_data.go` provides realistic test data without Proxmox connection

## Current State

**What works (Phases 1 & 2 complete):**
- Dashboard TUI showing live Proxmox node info (CPU, Memory, Disk usage)
- VM and LXC container lists with unified cursor navigation, sorted by ID
- Keyboard navigation (up/down or j/k, with wrapping)
- Quit (Ctrl+C or Esc; 'q' returns from detail view)
- Detail view for selected VM or LXC (tracks item by ID — stable across refreshes)
- Periodic async refresh via `tea.Tick` (interval from `config.refresh.node_stats`)
- API errors displayed in red in dashboard (no silent failures)
- Config from `~/.config/godash/config.yaml`; falls back to fake data with warning
- Fake data generators for development without Proxmox connection
- Lipgloss Tokyo Night styling throughout
- Unit tests for config, proxmox client, and UI model

**What's next (Phase 3):**
- Add start/stop/restart actions for VMs and LXCs via Proxmox API
- Keybindings in detail view with confirmation prompt for destructive actions
- Explore SSH into LXC containers from TUI

## Future Phases

- **Phase 3**: Proxmox control actions (start/stop/restart VMs and LXCs)
- **Phase 4**: Smart update checking with GitHub API integration for changelog awareness
- **Backlog**: Configurable list sorting (by name, CPU, memory, status)

## Important Notes

- **Fake Data**: Anonymous/fake data in `fake_data.go`, not real homelab data
- **Development Approach**: UI-first with fake data; APIs come later
- **Go Practices**: Idiomatic Go (gofmt, conventional naming, defer adding dependencies until needed)
- **Branching Strategy**: Feature branches off main, merge when complete
- **Dependencies**: Add only when actually needed, not speculatively

## Coding Guidelines

1. **Follow Elm Architecture**: All state changes through Update(), pure View()
2. **Idiomatic Go**: Use gofmt, Go naming conventions, keep simple
3. **Incremental Changes**: Small steps, test often
4. **Guidance Over Implementation**: Provide guidance; let user implement unless explicitly asked to write code
5. **Write Unit Tests**: Add tests for new functionality, especially API client methods