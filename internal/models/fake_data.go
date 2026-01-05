package models

// GetFakeNode returns fake node data for testing
func GetFakeNode() Node {
	return Node{
		Name:      "node-01",
		Status:    "online",
		CPU:       67.5,                      // Warning level (60-80%)
		CPUCores:  16,
		MemUsed:   95 * 1024 * 1024 * 1024,   // 95GB (~74% - warning)
		MemTotal:  128 * 1024 * 1024 * 1024,  // 128GB
		DiskUsed:  900 * 1024 * 1024 * 1024,  // 900GB (~88% - critical)
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
			CPU:      85.3,                   // Critical level (80%+)
			MemUsed:  4 * 1024 * 1024 * 1024, // 4GB (50% - normal)
			MemTotal: 8 * 1024 * 1024 * 1024, // 8GB
			Uptime:   1296000,                // 15 days
		},
		{
			ID:       101,
			Name:     "database-server",
			Status:   "running",
			CPU:      45.8,                     // Normal level
			MemUsed:  14 * 1024 * 1024 * 1024,  // 14GB (87.5% - critical)
			MemTotal: 16 * 1024 * 1024 * 1024,  // 16GB
			Uptime:   1296000,                  // 15 days
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
			CPU:      72.4,               // Warning level (60-80%)
			MemUsed:  128 * 1024 * 1024,  // 128MB (25% - normal)
			MemTotal: 512 * 1024 * 1024,  // 512MB
			Uptime:   2592000,            // 30 days
		},
		{
			ID:       201,
			Name:     "monitoring",
			Status:   "running",
			CPU:      15.2,               // Normal level
			MemUsed:  800 * 1024 * 1024,  // 800MB (78% - warning)
			MemTotal: 1024 * 1024 * 1024, // 1GB
			Uptime:   2592000,            // 30 days
		},
		{
			ID:       202,
			Name:     "vpn-gateway",
			Status:   "running",
			CPU:      3.7,               // Normal level
			MemUsed:  96 * 1024 * 1024,  // 96MB (18.75% - normal)
			MemTotal: 512 * 1024 * 1024, // 512MB
			Uptime:   2592000,           // 30 days
		},
	}
}
