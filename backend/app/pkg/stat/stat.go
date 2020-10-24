package stat

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
)

// CPU is the cpu type
type CPU struct {
	Total, User, System, Idle uint64
}

// Mem is the memory type
type Mem struct {
	Used, Available, Total uint64
}

// Disk is the disk type
type Disk struct {
	Used, Total uint64
}

// Usage method returns memory usage
func (m *Mem) Usage() float64 {
	return 1 - float64(m.Available)/float64(m.Total)
}

// Usage method returns disk usage
func (d *Disk) Usage() float64 {
	return float64(d.Used) / float64(d.Total)
}

// ServerStat is the machine stat type
type ServerStat struct {
	*CPU
	*Mem
	*Disk
}

// CalCPUUsage calculates CPU usage
func CalCPUUsage() (*CPU, error) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return &CPU{}, err
	}
	var st CPU
	for _, s := range stat.CPUStats {
		st.Total += s.IOWait + s.IRQ + s.Idle + s.Nice + s.SoftIRQ + s.Steal + s.System + s.User
		st.User += s.User
		st.System += s.System
		st.Idle += s.Idle
	}
	return &st, nil
}

// CalMemUsage calculates memory usage
func CalMemUsage() (*Mem, error) {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return &Mem{}, err
	}
	return &Mem{
		Used:      stat.MemTotal - stat.MemAvailable,
		Available: stat.MemAvailable,
		Total:     stat.MemTotal,
	}, nil
}

// CalDiskUsage calculates disk usage
func CalDiskUsage() (*Disk, error) {
	stat, err := linuxproc.ReadDisk("/")
	if err != nil {
		return &Disk{}, err
	}
	return &Disk{Used: stat.Used, Total: stat.All}, nil
}

// GetServerStat returns a server instance for machine stat
func GetServerStat() (*ServerStat, error) {
	var cpu *CPU
	var mem *Mem
	var disk *Disk
	var err error
	cpu, err = CalCPUUsage()
	if err != nil {
		return nil, err
	}
	mem, err = CalMemUsage()
	if err != nil {
		return nil, err
	}
	disk, err = CalDiskUsage()
	if err != nil {
		return nil, err
	}
	return &ServerStat{cpu, mem, disk}, nil
}
