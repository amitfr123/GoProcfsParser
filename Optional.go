package procfsgo

import "errors"

/*
	This struct was created for getting snaps of files that require root privilege s
*/

type DefaultAble[T any] interface {
	Default() T
}

type Optional[T DefaultAble[T]] struct {
	valid bool
	data  T
}

func (o Optional[T]) Get() (T, error) {
	if o.valid {
		return o.data, nil
	}
	return o.data.Default(), errors.New("nil value")
}

func (o *Optional[T]) Set(d T) {
	o.valid = true
	o.data = d
}
