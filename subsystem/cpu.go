package subsystem

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)


type CpuQuota struct {

}

func (c *CpuQuota) Set(cgroupPath string, resourceConfig *ResourceConfig) error {
	filepath := path.Join(cgroupPath, "cpu.max")
	args := strings.Split(resourceConfig.CpuQuota, "/")
	if len(args) < 2 {
		return fmt.Errorf("resource CpuQuota argument miss. stynx: 5000/10000")
	}
	err := os.WriteFile(filepath, []byte(fmt.Sprintf("%v %v", args[0], args[1])), os.ModePerm)
	if err != nil {
		return fmt.Errorf("set cpu.max fail, write file fail\n%v",err)
	}
	return nil
}

func (c *CpuQuota) Apply(cgroupPath string, pid int) error {
	filepath := path.Join(cgroupPath, "cgroup.procs")

	println(filepath)
	err := os.WriteFile(filepath, []byte(strconv.Itoa(pid)), os.ModePerm)
	if err != nil {
		return fmt.Errorf("set cgroup.proc fail, write file fail\n%v", err)
	}
	return nil
}

func (c *CpuQuota) Destroy(cgroupPath string) {
	if err := os.Remove(cgroupPath); err != nil {
		panic(fmt.Errorf("destory memory fail.\n%v", err))
	}
}

