package proxmox

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/talonlikeaclaw/godash/internal/config"
)

func testServer(handler http.HandlerFunc) (*httptest.Server, *config.Config) {
	srv := httptest.NewServer(handler)
	cfg := &config.Config{
		Proxmox: config.ProxmoxConfig{
			Host:  "127.0.0.1",
			Port:  0,
			Token: "root@pam!test=abc123",
			Node:  "pve",
			SSL:   true,
		},
	}
	return srv, cfg
}

func TestNew(t *testing.T) {
	srv, cfg := testServer(func(w http.ResponseWriter, r *http.Request) {})
	defer srv.Close()

	c := New(cfg)
	if c == nil {
		t.Fatal("New() returned nil")
	}
}

func TestGetNodeStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api2/json/nodes/pve/status" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "PVEAPIToken=root@pam!test=abc123" {
			t.Errorf("missing or wrong auth header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"data": {
				"cpu": 0.05,
				"cpuinfo": {"cpus": 4},
				"memory": {"used": 2147483648, "total": 8589934592},
				"rootfs":  {"used": 10737418240, "total": 107374182400},
				"uptime": 3600
			}
		}`)
	}

	srv, cfg := testServer(handler)
	defer srv.Close()

	c := New(cfg)
	c.baseURL = srv.URL + "/api2/json"

	node, err := c.GetNodeStatus()
	if err != nil {
		t.Fatalf("GetNodeStatus() error: %v", err)
	}
	if node.CPU != 5.0 {
		t.Errorf("CPU: got %.2f, want 5.0", node.CPU)
	}
	if node.CPUCores != 4 {
		t.Errorf("CPUCores: got %d, want 4", node.CPUCores)
	}
	if node.MemUsed != 2147483648 {
		t.Errorf("MemUsed: got %d, want 2147483648", node.MemUsed)
	}
	if node.MemTotal != 8589934592 {
		t.Errorf("MemTotal: got %d, want 8589934592", node.MemTotal)
	}
	if node.Uptime != 3600 {
		t.Errorf("Uptime: got %d, want 3600", node.Uptime)
	}
}
