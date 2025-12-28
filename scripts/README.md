# Development Scripts

## setup-dev-env.sh

Sets up the distrobox development environment for Godash.

**Usage:**
```bash
# Use default name (godash-dev)
./scripts/setup-dev-env.sh

# Use custom name
./scripts/setup-dev-env.sh my-go-container
```

This will:
- Check if distrobox is installed
- Create a container with Go 1.25
- Provide next steps for getting started

**Parameters:**
- `$1` (optional): Container name (default: `godash-dev`)

**Requirements:**
- distrobox installed on your system
- podman or docker as container runtime
