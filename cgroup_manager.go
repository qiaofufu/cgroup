package cgroup

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"github.com/qiaofufu/cgroup/subsystem"
)


type CgroupManager struct {
	Name string
	mountPath string
	Subsystems  []subsystem.Subsystem
}

// NewCgroupManager name is cgroup name
func NewCgroupManager(name string) (*CgroupManager, error) {
	path := path.Join("/sys/fs/cgroup", name)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil  {
			return nil, err
		}
	}
	subsystems := []subsystem.Subsystem {
		&subsystem.MemorySubSystem{},
		&subsystem.CpuQuota{},
		&subsystem.Cpuset{},
	}
	return &CgroupManager{Name: name, mountPath: path, Subsystems: subsystems}, nil
}



// SetMountPoint set cgroup mount point
func (cm *CgroupManager) SetMountPoint(path string) error {
	cmd := exec.Command("mount", []string{"-t", "none", path}...)

	err := cmd.Start()

	if err != nil {
		return fmt.Errorf("mount point fail. path:%v", path)
	}

	cm.mountPath = path
	return nil
}

func (cm *CgroupManager) Set(res *subsystem.ResourceConfig) error {
	for _, v := range cm.Subsystems {
		if err := v.Set(cm.GetCgroupPath(), res); err != nil {
			return err
		}
	}
	println(res.MemoryLimit)
	return nil
}

func (cm *CgroupManager) Apply(pid int) error {
	for _, v := range cm.Subsystems {
		if err := v.Apply(cm.GetCgroupPath(), pid); err != nil {
			return err
		}
	}
	return nil
}

func (cm *CgroupManager) Destroy()  {
	for _, v := range cm.Subsystems {
		v.Destroy(cm.GetCgroupPath())
	}
}

func (cm *CgroupManager) GetCgroupPath() string {
	
	return cm.mountPath
}