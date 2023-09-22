package procfsgo

import (
	"fmt"
	"io/fs"
	"os"
)

type ProcFsReader interface {
	GetProcFsSnap() (ps ProcFsSnap, err error)
	PrintMeminfo()
	PrintNetDev()
	PrintArp()
	GetProcTree()
}

type procParser struct {
	fsys fs.FS
}

type ProcFsSnap struct {
}

func GetProcReader() ProcFsReader {
	return procParser{os.DirFS("/proc")}
}

func (pp procParser) PrintMeminfo() {
	str, _e := fs.ReadFile(pp.fsys, "meminfo")
	if _e != nil {
		fmt.Printf("_e: %v\n", _e)
	} else {
		fmt.Printf("%s\n", str)
	}
}

func (pp procParser) PrintNetDev() {
	str, _e := fs.ReadFile(pp.fsys, "net/dev")
	if _e != nil {
		fmt.Printf("_e: %v\n", _e)
	} else {
		fmt.Printf("%s\n", str)
		v, _ := ParseNetDevicesString(string(str))
		l := len(v)
		l++
	}
}

func (pp procParser) PrintArp() {
	str, _e := fs.ReadFile(pp.fsys, "net/arp")
	if _e != nil {
		fmt.Printf("_e: %v\n", _e)
	} else {
		fmt.Printf("%s\n", str)
		v, _ := ParseArpTableString(string(str))
		l := len(v)
		l++
	}
}

func (pp procParser) GetProcFsSnap() (ps ProcFsSnap, err error) {
	return
}

func (pp procParser) GetProcTree() {
	pm = map[uint32]NetDeviceSnap{}
	dirE, err := fs.ReadDir(pp.fsys, "./")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for _, e := range dirE {

		}
	}
}
