package procfsgo

import (
	"errors"
	"strconv"
)

func IParseU64(arr []string, i *int, b int) (v uint64, err error) {
	if i == nil || len(arr) <= *i {
		err = errors.New("invalid params")
		return
	}
	v, err = strconv.ParseUint(arr[*i], b, 64)
	*i++
	return
}
