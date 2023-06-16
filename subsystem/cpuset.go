package subsystem

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)
type Cpuset struct {

}

func (c *Cpuset) Set(cgroupPath string, resourceConfig *ResourceConfig) error {
	filepath := path.Join(cgroupPath, "cpu.max")
	ss := strings.Split(resourceConfig.Cpuset, "-")
	args := ""
	
	if len(ss) < 2 {
		args = ss[0]
	} else {
		args = resourceConfig.Cpuset
	}
	err := os.WriteFile(filepath, []byte(args), os.ModePerm)
	if err != nil {
		return fmt.Errorf("set cpu.max fail, write file fail\n%v",err)
	}
	return nil
}

func (c *Cpuset) Apply(cgroupPath string, pid int) error {
	filepath := path.Join(cgroupPath, "cgroup.procs")

	println(filepath)
	err := os.WriteFile(filepath, []byte(strconv.Itoa(pid)), os.ModePerm)
	if err != nil {
		return fmt.Errorf("set cgroup.proc fail, write file fail\n%v", err)
	}
	return nil
}

func (c *Cpuset) Destroy(cgroupPath string) {
	if err := os.Remove(cgroupPath); err != nil {
		panic(fmt.Errorf("destory memory fail.\n%v", err))
	}
}

