package stat

import (
	"log"

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

// Server is the machine stat type
type Server struct {
	CPU
	Mem
	Disk
}

// CalCPUUsage calculates CPU usage
func CalCPUUsage() CPU {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var st CPU
	for _, s := range stat.CPUStats {
		st.Total += s.IOWait + s.IRQ + s.Idle + s.Nice + s.SoftIRQ + s.Steal + s.System + s.User
		st.User += s.User
		st.System += s.System
		st.Idle += s.Idle
	}
	return st
}

// CalMemUsage calculates memory usage
func CalMemUsage() Mem {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("stat read fail")
	}
	return Mem{Used: stat.MemTotal - stat.MemAvailable, Available: stat.MemAvailable, Total: stat.MemTotal}
}

// CalDiskUsage calculates disk usage
func CalDiskUsage() Disk {
	stat, err := linuxproc.ReadDisk("/")
	if err != nil {
		log.Fatal("stat read fail")
	}
	return Disk{Used: stat.Used, Total: stat.All}
}

// GetServer returns a server instance for machine stat
func GetServer() *Server {
	return &Server{CalCPUUsage(), CalMemUsage(), CalDiskUsage()}
}
