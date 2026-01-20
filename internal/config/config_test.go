package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	config, err := Load("../../testdata/valid_config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if config.Proxmox.Host != "192.168.0.123" {
		t.Errorf("Host = %s, want %s", config.Proxmox.Host, "192.168.0.123")
	}

	if config.Proxmox.Port != 8005 {
		t.Errorf("Port = %d, want %d", config.Proxmox.Port, 8005)
	}

	if config.Proxmox.Token != "username@pam!token-name=your-token" {
		t.Errorf("Token = %s, want %s", config.Proxmox.Token, "username@pam!token-name=your-token")
	}

	if config.Proxmox.Node != "pve" {
		t.Errorf("Node = %s, want %s", config.Proxmox.Node, "pve")
	}

	if config.Proxmox.SSL != false {
		t.Errorf("SSL = %t, want %t", config.Proxmox.SSL, false)
	}

	if config.Refresh.NodeStats != 10 {
		t.Errorf("NodeStats = %d, want %d", config.Refresh.NodeStats, 10)
	}
}
