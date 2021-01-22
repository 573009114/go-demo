package hao

import (
	"github.com/shirou/gopsutil/cpu"
)

// GetCPUInfo is a function
func GetCPUInfo() ([]cpu.InfoStat, error) {
	cpuinfos, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return cpuinfos, nil
}
