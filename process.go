package process

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/net"
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

type CpuInfoStat struct {
	Percent float64 `json:"percent"`
}

type ProcStat struct {
	Pid int32 `json:"pid"`
	Percent float64
	CpuInfo *CpuInfoStat
	MemInfo *process.MemoryInfoStat
	MemInfoEx *process.MemoryInfoExStat
	NumThreads int32
	Threads map[int32]*cpu.TimesStat
	Fd int32 `json:"fd"`
	Nice int32
	NumCtxSwitches *process.NumCtxSwitchesStat
	IOCounters *process.IOCountersStat
	NetIOCounters []net.IOCountersStat
	NetConnection []net.ConnectionStat
	CreateTime int64
	Status string
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

	// log.Info(p.Status())

	// log.Info(p.Uids())

	// log.Info(p.Gids())

	// log.Info(p.Tgid())

	// log.Info(p.Terminal())

	// log.Info(p.Kill())

	var pmi *process.MemoryInfoStat
	if pmi, err = p.MemoryInfo();err != nil {
		return nil, err
	}

	var pcp float64
	if pcp, err = p.CPUPercent();err != nil {
		return nil, err
	}

	var pmie *process.MemoryInfoExStat
	if pmie, err = p.MemoryInfoEx();err != nil {
		return nil, err
	}

	// fmt.Println(p.MemoryMaps(true))

	// fmt.Println(p.Rlimit())

	// fmt.Println(p.RlimitUsage(true))

	// fmt.Println(p.Children())

	// fmt.Println(p.Times())

	var pnt int32
	if pnt, err = p.NumThreads(); err != nil {
		return nil, err
	}

	var pthreads map[int32]*cpu.TimesStat
	if pthreads, err = p.Threads(); err != nil {
		return nil, err
	}

	var pnf int32
	if pnf, err = p.NumFDs(); err != nil {
		return nil, err
	}

	// fmt.Println(p.OpenFiles())

	var pnice int32
	if pnice, err = p.Nice(); err != nil {
		return nil, err
	}

	var pncs *process.NumCtxSwitchesStat
	if pncs, err = p.NumCtxSwitches(); err != nil {
		return nil, err
	}

	var pioc *process.IOCountersStat
	if pioc, err = p.IOCounters(); err != nil {
		return nil, err
	}

	var pnioc []net.IOCountersStat
	if pnioc, err = p.NetIOCounters(true); err != nil {
		return nil, err
	}

	var pconn []net.ConnectionStat
	if pconn, err = p.Connections(); err != nil {
		return nil, err
	}

	var pctime int64
	if pctime, err = p.CreateTime(); err != nil {
		return nil, err
	}

	ps = &ProcStat{
		Pid: pid,
		Percent: pp,
		CpuInfo: &CpuInfoStat{
			Percent: pcp,
		},
		MemInfo: pmi,
		MemInfoEx: pmie,
		NumThreads: pnt,
		Threads: pthreads,
		Fd: pnf,
		Nice: pnice,
		NumCtxSwitches: pncs,
		IOCounters: pioc,
		NetIOCounters: pnioc,
		NetConnection: pconn,
		CreateTime: pctime,
	}
	return
}
