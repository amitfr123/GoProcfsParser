package procfsgo

import (
	"errors"
	"strconv"
	"strings"
)

/*
meminfo contains many different metric types.
Some of them should always be there and some are dependent on kernel version or build options
meminfo line format:
	name value <kB is optional>
*/

type MeminfoSnap struct {
	fieldMap map[string]uint64
}

const minCol = 2

func ParseMeminfoString(str string) (snap MeminfoSnap, err error) {
	m := MeminfoSnap{make(map[string]uint64)}
	l := strings.Split(str, "\n")
	if len(l) == 0 {
		err = errors.New("invalid format")
		return
	}
	for _, e := range l {
		lf := strings.Fields(e)
		if len(lf) < minCol {
			err = errors.New("invalid format")
			return
		}
		k := strings.TrimRight(lf[0], ":")
		v, _e := strconv.ParseUint(lf[1], 10, 64)
		if _e != nil {
			err = errors.New("invalid format")
			return
		}
		m.fieldMap[k] = v
	}
	snap = m
	return
}
