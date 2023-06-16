package subsystem

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct {

}

// Set set memory max size 
func (m *MemorySubSystem) Set(cgroupPath string, resourceConfig *ResourceConfig) error {
	
	filepath := path.Join(cgroupPath, "memory.max")
	// println(filepath, resourceConfig.MemoryLimit)
	err := os.WriteFile(filepath, []byte(resourceConfig.MemoryLimit), os.ModePerm)
	if err != nil {
		return fmt.Errorf("set memory.max fail, write file fail\n%v",err)
	}
	return nil
}

func (m *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	filepath := path.Join(cgroupPath, "cgroup.procs")

	// println(filepath)
	err := os.WriteFile(filepath, []byte(strconv.Itoa(pid)), os.ModePerm)
	if err != nil {
		return fmt.Errorf("set cgroup.proc fail, write file fail\n%v", err)
	}
	return nil
}

func (m *MemorySubSystem) Destroy(cgroupPath string) {
	if err := os.Remove(cgroupPath); err != nil {
		panic(fmt.Errorf("destory memory fail.\n%v", err))
	}
}

