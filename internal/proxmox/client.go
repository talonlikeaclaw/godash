package proxmox

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/talonlikeaclaw/godash/internal/config"
	"github.com/talonlikeaclaw/godash/internal/models"
)

// Client holds the HTTP client and connection details for the Proxmox API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	node       string
}

// New creates a Proxmox API client from the provided config.
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

type apiResponse[T any] struct {
	Data T `json:"data"`
}

type nodeStatusResponse struct {
	CPU     float64 `json:"cpu"`
	CPUInfo struct {
		CPUs int `json:"cpus"`
	} `json:"cpuinfo"`
	Memory struct {
		Used  int64 `json:"used"`
		Total int64 `json:"total"`
	} `json:"memory"`
	RootFS struct {
		Used  int64 `json:"used"`
		Total int64 `json:"total"`
	} `json:"rootfs"`
	Uptime int64 `json:"uptime"`
}

// get performs an authenticated GET request and returns the raw response body.
func (c *Client) get(path string) ([]byte, error) {
	url := c.baseURL + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Authorization", "PVEAPIToken="+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get %s: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("proxmox api %s: status %d", path, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}
	return body, nil
}

type vmResponse struct {
	VMID   int     `json:"vmid"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
	CPU    float64 `json:"cpu"`
	Mem    int64   `json:"mem"`
	MaxMem int64   `json:"maxmem"`
	Uptime int64   `json:"uptime"`
}

// GetVMs returns all QEMU virtual machines on the configured node.
func (c *Client) GetVMs() ([]models.VM, error) {
	body, err := c.get("/nodes/" + c.node + "/qemu")
	if err != nil {
		return nil, err
	}

	var resp apiResponse[[]vmResponse]
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse vms: %w", err)
	}

	vms := make([]models.VM, len(resp.Data))
	for i, v := range resp.Data {
		vms[i] = models.VM{
			ID:       v.VMID,
			Name:     v.Name,
			Status:   v.Status,
			CPU:      v.CPU * 100,
			MemUsed:  v.Mem,
			MemTotal: v.MaxMem,
			Uptime:   v.Uptime,
		}
	}
	return vms, nil
}

// GetNodeStatus fetches CPU, memory, disk, and uptime for the configured node.
func (c *Client) GetNodeStatus() (models.Node, error) {
	body, err := c.get("/nodes/" + c.node + "/status")
	if err != nil {
		return models.Node{}, err
	}

	var resp apiResponse[nodeStatusResponse]
	if err := json.Unmarshal(body, &resp); err != nil {
		return models.Node{}, fmt.Errorf("parse node status: %w", err)
	}

	d := resp.Data
	return models.Node{
		Name:      c.node,
		Status:    "online",
		CPU:       d.CPU * 100,
		CPUCores:  d.CPUInfo.CPUs,
		MemUsed:   d.Memory.Used,
		MemTotal:  d.Memory.Total,
		DiskUsed:  d.RootFS.Used,
		DiskTotal: d.RootFS.Total,
		Uptime:    d.Uptime,
	}, nil
}
