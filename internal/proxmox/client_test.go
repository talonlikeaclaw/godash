package proxmox

import (
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
