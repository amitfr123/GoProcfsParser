package procfsgo

import (
	"errors"
	"reflect"
	"strings"
)

/*
Net dev contains data on local machine network interfaces.

Net dev format:
Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
int1:  number    number    number    number    number     number          number         number   number    number    number    number    number     number       number          number
int2:  number    number    number    number    number     number          number         number   number    number    number    number    number     number       number          number
.
.
.
<an extra \n for some reason>
*/

type NetDeviceSnap struct {
	rB    uint64
	rP    uint64
	rE    uint64
	rD    uint64
	rF    uint64
	rFr   uint64
	rComp uint64
	rM    uint64
	tB    uint64
	tP    uint64
	tE    uint64
	tD    uint64
	tF    uint64
	tCol  uint64
	tCar  uint64
	tComp uint64
}

const ndHeadLineCount = 2

func ParseNetDevicesString(str string) (snapMap map[string]NetDeviceSnap, err error) {
	ndm := map[string]NetDeviceSnap{}
	// The /proc/net/dev file contains a newline before eof
	str = strings.TrimRight(str, "\n")
	l := strings.Split(str, "\n")
	if len(l) < ndHeadLineCount {
		err = errors.New("invalid format")
		return
	}
	// The /proc/net/dev first 2 lines contain the table headers
	dl := l[ndHeadLineCount:]
	for _, e := range dl {
		lf := strings.Fields(e)
		if len(lf)-1 != reflect.TypeOf(NetDeviceSnap{}).NumField() {
			err = errors.New("invalid format")
			return
		}
		n := strings.TrimRight(lf[0], ":")
		mem := NetDeviceSnap{}
		fi := 1
		mem.rB, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rP, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rE, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rD, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rF, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rFr, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rComp, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.rM, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tB, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tP, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tE, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tD, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tF, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tCol, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tCar, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		mem.tComp, err = IParseU64(lf, &fi, 10)
		if err != nil {
			return
		}
		ndm[n] = mem
	}
	snapMap = ndm
	return
}
