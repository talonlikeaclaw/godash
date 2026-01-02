package models

// GetFakeNode returns fake node data for testing
func GetFakeNode() Node {
	return Node{
		Name:      "node-01",
		Status:    "online",
		CPU:       12.5,
		CPUCores:  16,
		MemUsed:   48 * 1024 * 1024 * 1024,   // 48GB
		MemTotal:  128 * 1024 * 1024 * 1024,  // 128GB
		DiskUsed:  250 * 1024 * 1024 * 1024,  // 250GB
		DiskTotal: 1024 * 1024 * 1024 * 1024, // 1TB
		Uptime:    2592000,                   // 30 days in seconds
	}
}

// GetFakeVMs returns fake VM data for testing
func GetFakeVMs() []VM {
	return []VM{
		{
			ID:       100,
			Name:     "web-server",
			Status:   "running",
			CPU:      2.5,
			MemUsed:  4 * 1024 * 1024 * 1024, // 4GB
			MemTotal: 8 * 1024 * 1024 * 1024, // 8GB
			Uptime:   1296000,                // 15 days
		},
		{
			ID:       101,
			Name:     "database-server",
			Status:   "running",
			CPU:      8.3,
			MemUsed:  12 * 1024 * 1024 * 1024, // 12GB
			MemTotal: 16 * 1024 * 1024 * 1024, // 16GB
			Uptime:   1296000,                 // 15 days
		},
		{
			ID:       102,
			Name:     "dev-environment",
			Status:   "stopped",
			CPU:      0.0,
			MemUsed:  0,
			MemTotal: 4 * 1024 * 1024 * 1024, // 4GB
			Uptime:   0,
		},
	}
}

// GetFakeContainers returns fake container data for testing
func GetFakeContainers() []Container {
	return []Container{
		{
			ID:       200,
			Name:     "reverse-proxy",
			Status:   "running",
			CPU:      0.5,
			MemUsed:  128 * 1024 * 1024, // 128MB
			MemTotal: 512 * 1024 * 1024, // 512MB
			Uptime:   2592000,           // 30 days
		},
		{
			ID:       201,
			Name:     "monitoring",
			Status:   "running",
			CPU:      1.2,
			MemUsed:  256 * 1024 * 1024,  // 256MB
			MemTotal: 1024 * 1024 * 1024, // 1GB
			Uptime:   2592000,            // 30 days
		},
		{
			ID:       202,
			Name:     "vpn-gateway",
			Status:   "running",
			CPU:      0.3,
			MemUsed:  96 * 1024 * 1024,  // 96MB
			MemTotal: 512 * 1024 * 1024, // 512MB
			Uptime:   2592000,           // 30 days
		},
	}
}
