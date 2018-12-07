package process

import (
	"github.com/shirou/gopsutil/process"
)

type MemoryInfoStat struct {
	RSS    uint64 `json:"rss"`    // bytes
	VMS    uint64 `json:"vms"`    // bytes
	Data   uint64 `json:"data"`   // bytes
	Stack  uint64 `json:"stack"`  // bytes
	Locked uint64 `json:"locked"` // bytes
	Swap   uint64 `json:"swap"`   // bytes
}

type ProcStat struct {
	Pid int32 `json:"pid"`
	Mem *MemoryInfoStat
	Cpu string
}

/*
func NewProcess(pid int32) (procStats *process.Process, err error) {
	if p, err = process.NewProcess(pid); err != nil {
		log.Println(err)
		return nil, err
	}
	return p, err
}
*/

func MemoryInfo(pid int32) (*ProcStat, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return nil, err
	}

	pm, err := p.MemoryInfo()
	if err != nil {
		return nil, err
	}

	ps := &ProcStat{Pid: pid, Mem: &MemoryInfoStat{RSS: pm.RSS, VMS: pm.VMS, Swap: pm.Swap}}
	return  ps, nil
}