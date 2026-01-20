package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type Config struct {
	Proxmox ProxmoxConfig `yaml:"proxmox"`
	Refresh RefreshConfig `yaml:"refresh"`
}

type ProxmoxConfig struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Token string `yaml:"token"`
	Node  string `yaml:"node"`
	SSL   bool   `yaml:"verify_ssl"`
}
type RefreshConfig struct {
	NodeStats int `yaml:"node_stats"`
}
