package proxmox

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/talonlikeaclaw/godash/internal/config"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	node       string
}

func New(cfg *config.Config) *Client {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: !cfg.Proxmox.SSL,
	}
	transport := &http.Transport{TLSClientConfig: tlsCfg}

	return &Client{
		httpClient: &http.Client{Transport: transport},
		baseURL:    fmt.Sprintf("https://%s:%d/api2/json", cfg.Proxmox.Host, cfg.Proxmox.Port),
		token:      cfg.Proxmox.Token,
		node:       cfg.Proxmox.Node,
	}
}
