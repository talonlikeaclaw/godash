#!/bin/bash
set -euo pipefail

CONTAINER_NAME="${1:-godash-dev}"

echo "Setting up Godash development environment..."
echo "Container name: $CONTAINER_NAME"

# Check if distrobox is installed
if ! command -v distrobox &>/dev/null; then
  echo "Error: distrobox is not installed"
  echo "Please install distrobox first: https://github.com/89luca89/distrobox"
  exit 1
fi

# Check if container already exists
if distrobox list | grep -q "^[[:space:]]*${CONTAINER_NAME}[[:space:]]*$"; then
  echo "$CONTAINER_NAME container already exists"
  read -p "Do you want to recreate it? (y/n) " -r
  echo
  if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "Removing existing container..."
    distrobox rm -f "$CONTAINER_NAME"
  else
    echo "Skipping container creation"
    exit 0
  fi
fi

# Create container
IMAGE_TAG="golang:1.25"
echo "Creating $CONTAINER_NAME container with Go 1.25..."
distrobox create --image "$IMAGE_TAG" --name "$CONTAINER_NAME"

echo "Container created successfully!"

echo ""
echo "Next steps:"
echo "  1) Enter the container: distrobox enter $CONTAINER_NAME"
echo "  2) Navigate to project: cd ~/projects/godash"
echo "  3) Install Go tools: go install github.com/cosmtrek/air@latest"
echo "  4) Run the project: go run cmd/godash/main.go"
