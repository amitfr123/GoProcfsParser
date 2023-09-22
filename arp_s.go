package procfsgo

import (
	"errors"
	"net"
	"reflect"
	"strings"
)

type ArpSnap struct {
	hw     uint64
	f      uint64
	hwAddr string
	m      string
	dev    string
}

const arpHeadLineCount = 1

func ParseArpTableString(str string) (snapMap map[string]ArpSnap, err error) {
	am := map[string]ArpSnap{}
	// The /proc/net/dev file contains a newline before eof
	str = strings.TrimRight(str, "\n")
	l := strings.Split(str, "\n")
	if len(l) < arpHeadLineCount {
		err = errors.New("invalid format")
		return
	}
	// The /proc/net/arp first line contain the table headers
	dl := l[arpHeadLineCount:]
	for _, e := range dl {
		lf := strings.Fields(e)
		if len(lf)-1 != reflect.TypeOf(ArpSnap{}).NumField() {
			err = errors.New("invalid format")
			return
		}
		n := strings.TrimRight(lf[0], ":")
		if net.ParseIP(n) == nil {
			err = errors.New("invalid format")
			return
		}
		mem := ArpSnap{}
		fi := 1
		lf[fi] = strings.TrimLeft(lf[fi], "0x")
		mem.hw, err = IParseU64(lf, &fi, 16)
		if err != nil {
			return
		}
		lf[fi] = strings.TrimLeft(lf[fi], "0x")
		mem.f, err = IParseU64(lf, &fi, 16)
		if err != nil {
			return
		}
		if _, err = net.ParseMAC(lf[fi]); err != nil {
			return
		}
		mem.hwAddr = lf[fi]
		fi++
		if lf[fi] != "*" && net.ParseIP(lf[fi]) == nil {
			err = errors.New("invalid format")
			return
		}
		mem.m = lf[fi]
		fi++
		mem.dev = lf[fi]
		am[n] = mem
	}
	snapMap = am
	return
}
