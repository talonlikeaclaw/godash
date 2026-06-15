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

func TestGetVMs(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api2/json/nodes/pve/qemu" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"data": [
				{
					"vmid": 100,
					"name": "ubuntu-private",
					"status": "running",
					"cpu": 0.024,
					"mem": 2147483648,
					"maxmem": 8589934592,
					"uptime": 3600
				}
			]
		}`)
	}

	srv, cfg := testServer(handler)
	defer srv.Close()

	c := New(cfg)
	c.baseURL = srv.URL + "/api2/json"

	vms, err := c.GetVMs()
	if err != nil {
		t.Fatalf("GetVMs() error: %v", err)
	}
	if len(vms) != 1 {
		t.Fatalf("expected 1 VM, got %d", len(vms))
	}
	vm := vms[0]
	if vm.ID != 100 {
		t.Errorf("ID: got %d, want 100", vm.ID)
	}
	if vm.Name != "ubuntu-private" {
		t.Errorf("Name: got %s, want ubuntu-private", vm.Name)
	}
	if vm.Status != "running" {
		t.Errorf("Status: got %s, want running", vm.Status)
	}
	if vm.CPU != 2.4 {
		t.Errorf("CPU: got %.2f, want 2.4", vm.CPU)
	}
	if vm.MemUsed != 2147483648 {
		t.Errorf("MemUsed: got %d, want 2147483648", vm.MemUsed)
	}
	if vm.MemTotal != 8589934592 {
		t.Errorf("MemTotal: got %d, want 8589934592", vm.MemTotal)
	}
	if vm.Uptime != 3600 {
		t.Errorf("Uptime: got %d, want 3600", vm.Uptime)
	}
}

func TestGetContainers(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api2/json/nodes/pve/lxc" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"data": [
				{
					"vmid": 101,
					"name": "adguard",
					"status": "running",
					"cpu": 0.002,
					"mem": 536870912,
					"maxmem": 1073741824,
					"uptime": 7200
				}
			]
		}`)
	}

	srv, cfg := testServer(handler)
	defer srv.Close()

	c := New(cfg)
	c.baseURL = srv.URL + "/api2/json"

	containers, err := c.GetContainers()
	if err != nil {
		t.Fatalf("GetContainers() error: %v", err)
	}
	if len(containers) != 1 {
		t.Fatalf("expected 1 container, got %d", len(containers))
	}
	ct := containers[0]
	if ct.ID != 101 {
		t.Errorf("ID: got %d, want 101", ct.ID)
	}
	if ct.Name != "adguard" {
		t.Errorf("Name: got %s, want adguard", ct.Name)
	}
	if ct.Status != "running" {
		t.Errorf("Status: got %s, want running", ct.Status)
	}
	if ct.CPU != 0.2 {
		t.Errorf("CPU: got %.4f, want 0.2", ct.CPU)
	}
	if ct.MemUsed != 536870912 {
		t.Errorf("MemUsed: got %d, want 536870912", ct.MemUsed)
	}
	if ct.MemTotal != 1073741824 {
		t.Errorf("MemTotal: got %d, want 1073741824", ct.MemTotal)
	}
	if ct.Uptime != 7200 {
		t.Errorf("Uptime: got %d, want 7200", ct.Uptime)
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
