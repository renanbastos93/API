package config

// Input are all needed values to compute all parameters
type Input struct {
	OS              string  `json:"os"`
	Arch            string  `json:"arch"`
	TotalRAM        int     `json:"total_ram"`
	Profile         string  `json:"profile"`
	DiskType        string  `json:"disk_type"`
	MaxConnections  int     `json:"max_connections"`
	PostgresVersion float32 `json:"postgres_version"`
}

// NewInput creates a Input
func NewInput(os string, arch string, totalRAM int, profile string, diskType string, maxConnections int, postgresVersion float32) *Input {
	return &Input{
		OS:              os,
		Arch:            arch,
		TotalRAM:        totalRAM,
		Profile:         profile,
		DiskType:        diskType,
		MaxConnections:  maxConnections,
		PostgresVersion: postgresVersion,
	}
}
