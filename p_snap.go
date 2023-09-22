package procfsgo

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type PStatSnap struct {
	pid                   int64
	comm                  string
	state                 byte
	ppid                  int
	pgrp                  int
	session               int
	tty_nr                int
	tpgid                 int
	flags                 uint
	minflt                uint64
	cminflt               uint64
	utime                 uint64
	stime                 uint64
	cutime                int64
	cstime                int64
	priority              int64
	nice                  int64
	num_threads           int64
	itrealvalue           int64
	starttime             uint64
	vsize                 uint64
	rss                   int64
	rsslim                uint64
	startcode             uint64
	endcode               uint64
	startstack            uint64
	kstkeip               uint64
	signal                uint64
	blocked               uint64
	sigignore             uint64
	sigcatch              uint64
	wchan                 uint64
	nswap                 uint64
	cnswap                uint64
	exit_signal           int
	processor             int
	rt_priority           uint
	policy                uint
	delayacct_blkio_ticks uint64
	guest_time            uint64
	cguest_time           int64
	start_data            uint64
	end_data              uint64
	start_brk             uint64
	arg_start             uint64
	arg_end               uint64
	env_start             uint64
	env_end               uint64
	exit_code             int
}

func ParsePStatString(str string) (snap PStatSnap, err error) {
	var s PStatSnap
	n, err := fmt.Sscan(str, &s.pid,
		&s.comm,
		&s.state,
		&s.ppid,
		&s.pgrp,
		&s.session,
		&s.tty_nr,
		&s.tpgid,
		&s.flags,
		&s.minflt,
		&s.cminflt,
		&s.utime,
		&s.stime,
		&s.cutime,
		&s.cstime,
		&s.priority,
		&s.nice,
		&s.num_threads,
		&s.itrealvalue,
		&s.starttime,
		&s.vsize,
		&s.rss,
		&s.rsslim,
		&s.startcode,
		&s.endcode,
		&s.startstack,
		&s.kstkeip,
		&s.signal,
		&s.blocked,
		&s.sigignore,
		&s.sigcatch,
		&s.wchan,
		&s.nswap,
		&s.cnswap,
		&s.exit_signal,
		&s.processor,
		&s.rt_priority,
		&s.policy,
		&s.delayacct_blkio_ticks,
		&s.guest_time,
		&s.cguest_time,
		&s.start_data,
		&s.end_data,
		&s.start_brk,
		&s.arg_start,
		&s.arg_end,
		&s.env_start,
		&s.env_end,
		&s.exit_code)
	if err != nil || n != reflect.TypeOf(snap).NumField() {
		err = errors.New("invalid format")
		return
	}
	s.comm = strings.Trim(s.comm, "()")
	snap = s
	return
}
