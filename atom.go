package goAtom

import (
	"sync/atomic"
)

type Atom[Data any] struct {
	value atomic.Value
}

func NewAtom[Data any]() *Atom[Data] {
	a := &Atom[Data]{}
	var d Data
	a.value.Store(d)
	return a
}

func (a *Atom[Data]) Swap(f func(Data) Data) {
	ok := false
	for !ok {
		old := a.value.Load().(Data)
		newValue := f(old)
		ok = a.value.CompareAndSwap(old, newValue)
	}
}

func (a *Atom[Data]) Get() Data {
	return a.value.Load().(Data)
}
