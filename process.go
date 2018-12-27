package process

import (
	"github.com/shirou/gopsutil/process"
	log "github.com/sirupsen/logrus"
)

const (
	RLIMIT_CPU        int32 = 0
	RLIMIT_FSIZE      int32 = 1
	RLIMIT_DATA       int32 = 2
	RLIMIT_STACK      int32 = 3
	RLIMIT_CORE       int32 = 4
	RLIMIT_RSS        int32 = 5
	RLIMIT_NPROC      int32 = 6
	RLIMIT_NOFILE     int32 = 7
	RLIMIT_MEMLOCK    int32 = 8
	RLIMIT_AS         int32 = 9
	RLIMIT_LOCKS      int32 = 10
	RLIMIT_SIGPENDING int32 = 11
	RLIMIT_MSGQUEUE   int32 = 12
	RLIMIT_NICE       int32 = 13
	RLIMIT_RTPRIO     int32 = 14
	RLIMIT_RTTIME     int32 = 15
)

type MemoryInfoStat struct {
	RSS    uint64 `json:"rss"`    // bytes
	VMS    uint64 `json:"vms"`    // bytes
	Data   uint64 `json:"data"`   // bytes
	Stack  uint64 `json:"stack"`  // bytes
	Locked uint64 `json:"locked"` // bytes
	Swap   uint64 `json:"swap"`   // bytes
}

type CpuInfoStat struct {
	Percent float64 `json:"percent"`
} 

type ProcStat struct {
	Pid int32 `json:"pid"`
	Mem *MemoryInfoStat
	Cpu *CpuInfoStat
	Percent float64
	status string
}

func ProcessInfo(pid int32) (ps *ProcStat, err error) {
	// log.Debug(process.Processes())
	// log.Debug(process.Pids())

	var ok bool
	if ok, err = process.PidExists(pid); ok != true || err != nil {
		log.Errorf("Process is not running! Error: %v", err)
		return
	}
	log.Debugf("Process is running status!")

	var p *process.Process
	if p, err = process.NewProcess(pid); err != nil {
		log.Errorf("Process instance created failed! Error: %v", err)
		return
	}
	log.Debugf("Process instance instance has been created!!")

	var pname string
	if pname, err = p.Name(); err != nil {
		return
	}
	log.Debugf("Process name is %v", pname)

	var pp float64
	if pp, err = p.Percent(0);err != nil {
		return
	}

	// log.Info(p.Cwd())

	// log.Info(p.Username())

	// log.Info(p.Cmdline())

	// log.Info(p.CmdlineSlice())

	log.Info(p.Status())

	// log.Info(p.Uids())

	// log.Info(p.Gids())

	// log.Info(p.Tgid())

	// log.Info(p.Terminal())

	// log.Info(p.Kill())



	pmi, err := p.MemoryInfo()
	if err != nil {
		return nil, err
	}

	pcp, err := p.CPUPercent()
	if err != nil {
		return nil, err
	}

	// fmt.Println(p.MemoryInfoEx())

	// fmt.Println(p.MemoryMaps(true))

	// fmt.Println(p.Rlimit())

	// fmt.Println(p.RlimitUsage(true))

	// fmt.Println(p.Children())

	// fmt.Println(p.Times())

	// fmt.Println(p.NumThreads())

	// fmt.Println(p.Threads())

	// fmt.Println(p.NumFDs())

	// fmt.Println(p.OpenFiles())

	// fmt.Println(p.Nice())

	// fmt.Println(p.NumCtxSwitches())

	// fmt.Println(p.IOCounters())

	// fmt.Println(p.NetIOCounters(true))

	// fmt.Println(p.Connections())

	// fmt.Println(p.CreateTime())

	ps = &ProcStat{
		Pid: pid,
		Percent: pp,
		Mem: &MemoryInfoStat{
			RSS: pmi.RSS,
			VMS: pmi.VMS,
			Swap: pmi.Swap,
		},
		Cpu: &CpuInfoStat{
			Percent: pcp,
		},
	}
	return
}
