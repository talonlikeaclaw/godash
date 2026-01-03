package models

type Node struct {
	Name      string
	Status    string
	CPU       float64
	CPUCores  int
	MemUsed   int64
	MemTotal  int64
	DiskUsed  int64
	DiskTotal int64
	Uptime    int64
}

type VM struct {
	ID       int
	Name     string
	Status   string
	CPU      float64
	MemUsed  int64
	MemTotal int64
	Uptime   int64
}

type Container struct {
	ID       int
	Name     string
	Status   string
	CPU      float64
	MemUsed  int64
	MemTotal int64
	Uptime   int64
}
