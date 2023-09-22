package procfsgo

import (
	"io/fs"
	"strconv"
)

type PNode struct {
	c   map[uint32]*PNode
	pid uint32
	pn  *PNode
}

func BuildTree(fsys fs.FS) (t PNode, err error) {
	pm = map[uint32]*PNode{}
	dirE, _e := fs.ReadDir(fsys, "./")
	if err != nil {
		err = _e
		return
	}
	dirE = prepTreeRoot(dirE)
	for _, e := range dirE {
		stat, _e := fs.ReadFile(fsys, e.Name()+"/status")
		if _e == nil {
			continue
		}

	}
}

func prepTreeRoot(dirE []fs.DirEntry) (out []fs.DirEntry) {
	for _, e := range dirE {
		if !e.IsDir() {
			continue
		}
		_, err := strconv.ParseUint(e.Name(), 10, 32)
		if err == nil {
			out = append(out, e)
		}
	}
	return
}
