package subsystem


type ResourceConfig struct {
	MemoryLimit string
	CpuQuota string
	Cpuset string
}

type Subsystem interface {
	Set(cgroupPath string, resourceConfig *ResourceConfig) error
	Apply(cgroupPath string, pid int) error
	Destroy(cgroupPath string) 
}

