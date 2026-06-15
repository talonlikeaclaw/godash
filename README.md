# Godash

A terminal-based dashboard for monitoring Proxmox homelab infrastructure, VMs, LXC containers, and node resources; in real-time.

## Features

**Working now:**

- Live Proxmox node stats (CPU, memory, disk) with progress bars
- VM and LXC container list with status, CPU, memory, uptime
- Detail view for selected VM or container
- Periodic auto-refresh (configurable interval)
- Tokyo Night colour theme via Lipgloss
- Keyboard navigation (j/k or ↑/↓, Enter, q, Esc)
- Graceful fallback to fake data when no config present

**Coming soon:**

- Start/stop/restart VMs and LXCs from the TUI
- SSH into LXC containers
- Configurable list sorting

## Configuration

Create `~/.config/godash/config.yaml`:

```yaml
proxmox:
  host: "192.168.0.100"
  port: 8006
  token: "root@pam!tokenid=<uuid>"
  node: "your-node-name" # check via GET /api2/json/nodes
  verify_ssl: false # false for self-signed certs

refresh:
  node_stats: 10 # seconds between refreshes
```

Create a Proxmox API token at: Datacenter > Permissions > API Tokens.

## Keyboard Controls

| Key              | Action            |
| ---------------- | ----------------- |
| `j` / `↓`        | Move down         |
| `k` / `↑`        | Move up           |
| `Enter`          | Open detail view  |
| `q`              | Back to dashboard |
| `Esc` / `Ctrl+C` | Quit              |

## Development Setup

### Prerequisites

- [distrobox](https://github.com/89luca89/distrobox)
- podman or docker

### Quick Start

```bash
git clone https://github.com/talonlikeaclaw/godash.git
cd godash

./scripts/setup-dev-env.sh
distrobox enter godash-dev
cd ~/projects/godash

make run
```

### Common Commands

```bash
make run        # Run app
make build      # Build binary → bin/godash
make test       # Run all tests
make coverage   # Test coverage report
make lint       # Run linter
```

## Project Structure

```
godash/
├── cmd/godash/          # Entry point
├── internal/
│   ├── proxmox/         # Proxmox API client
│   ├── config/          # YAML config loading
│   ├── models/          # Node, VM, Container structs + fake data
│   └── ui/              # Bubble Tea TUI components
└── configs/             # config.example.yaml
```
