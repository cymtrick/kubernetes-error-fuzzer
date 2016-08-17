package procfs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// Originally, this USER_HZ value was dynamically retrieved via a sysconf call which
// required cgo. However, that caused a lot of problems regarding
// cross-compilation. Alternatives such as running a binary to determine the
// value, or trying to derive it in some other way were all problematic.
// After much research it was determined that USER_HZ is actually hardcoded to
// 100 on all Go-supported platforms as of the time of this writing. This is
// why we decided to hardcode it here as well. It is not impossible that there
// could be systems with exceptions, but they should be very exotic edge cases,
// and in that case, the worst outcome will be two misreported metrics.
//
// See also the following discussions:
//
// - https://github.com/prometheus/node_exporter/issues/52
// - https://github.com/prometheus/procfs/pull/2
// - http://stackoverflow.com/questions/17410841/how-does-user-hz-solve-the-jiffy-scaling-issue
const userHZ = 100

// ProcStat provides status information about the process,
// read from /proc/[pid]/stat.
type ProcStat struct {
	// The process ID.
	PID int
	// The filename of the executable.
	Comm string
	// The process state.
	State string
	// The PID of the parent of this process.
	PPID int
	// The process group ID of the process.
	PGRP int
	// The session ID of the process.
	Session int
	// The controlling terminal of the process.
	TTY int
	// The ID of the foreground process group of the controlling terminal of
	// the process.
	TPGID int
	// The kernel flags word of the process.
	Flags uint
	// The number of minor faults the process has made which have not required
	// loading a memory page from disk.
	MinFlt uint
	// The number of minor faults that the process's waited-for children have
	// made.
	CMinFlt uint
	// The number of major faults the process has made which have required
	// loading a memory page from disk.
	MajFlt uint
	// The number of major faults that the process's waited-for children have
	// made.
	CMajFlt uint
	// Amount of time that this process has been scheduled in user mode,
	// measured in clock ticks.
	UTime uint
	// Amount of time that this process has been scheduled in kernel mode,
	// measured in clock ticks.
	STime uint
	// Amount of time that this process's waited-for children have been
	// scheduled in user mode, measured in clock ticks.
	CUTime uint
	// Amount of time that this process's waited-for children have been
	// scheduled in kernel mode, measured in clock ticks.
	CSTime uint
	// For processes running a real-time scheduling policy, this is the negated
	// scheduling priority, minus one.
	Priority int
	// The nice value, a value in the range 19 (low priority) to -20 (high
	// priority).
	Nice int
	// Number of threads in this process.
	NumThreads int
	// The time the process started after system boot, the value is expressed
	// in clock ticks.
	Starttime uint64
	// Virtual memory size in bytes.
	VSize int
	// Resident set size in pages.
	RSS int

	fs FS
}

// NewStat returns the current status information of the process.
func (p Proc) NewStat() (ProcStat, error) {
	f, err := p.open("stat")
	if err != nil {
		return ProcStat{}, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return ProcStat{}, err
	}

	var (
		ignore int

		s = ProcStat{PID: p.PID, fs: p.fs}
		l = bytes.Index(data, []byte("("))
		r = bytes.LastIndex(data, []byte(")"))
	)

	if l < 0 || r < 0 {
		return ProcStat{}, fmt.Errorf(
			"unexpected format, couldn't extract comm: %s",
			data,
		)
	}

	s.Comm = string(data[l+1 : r])
	_, err = fmt.Fscan(
		bytes.NewBuffer(data[r+2:]),
		&s.State,
		&s.PPID,
		&s.PGRP,
		&s.Session,
		&s.TTY,
		&s.TPGID,
		&s.Flags,
		&s.MinFlt,
		&s.CMinFlt,
		&s.MajFlt,
		&s.CMajFlt,
		&s.UTime,
		&s.STime,
		&s.CUTime,
		&s.CSTime,
		&s.Priority,
		&s.Nice,
		&s.NumThreads,
		&ignore,
		&s.Starttime,
		&s.VSize,
		&s.RSS,
	)
	if err != nil {
		return ProcStat{}, err
	}

	return s, nil
}

// VirtualMemory returns the virtual memory size in bytes.
func (s ProcStat) VirtualMemory() int {
	return s.VSize
}

// ResidentMemory returns the resident memory size in bytes.
func (s ProcStat) ResidentMemory() int {
	return s.RSS * os.Getpagesize()
}

// StartTime returns the unix timestamp of the process in seconds.
func (s ProcStat) StartTime() (float64, error) {
	stat, err := s.fs.NewStat()
	if err != nil {
		return 0, err
	}
	return float64(stat.BootTime) + (float64(s.Starttime) / userHZ), nil
}

// CPUTime returns the total CPU user and system time in seconds.
func (s ProcStat) CPUTime() float64 {
	return float64(s.UTime+s.STime) / userHZ
}
