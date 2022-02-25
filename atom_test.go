package goAtom

import (
	"sync"
	"testing"
)

func TestSwap(t *testing.T) {
	a := NewAtom[int]()
	a.Swap(func(x int) int {
		return x + 1
	})
	if a.value.Load().(int) != 1 {
		t.Errorf("a.value.Load().(int) != 1")
		return
	}
}

func TestGet(t *testing.T) {
	a := NewAtom[int]()
	a.value.Store(2)
	if a.Get() != 2 {
		t.Errorf("a.Get() != 2")
		return
	}
}

func TestAtomCounter(t *testing.T) {
	a := NewAtom[int]()

	w := sync.WaitGroup{}
	w.Add(100000)
	for i := 0; i < 100000; i++ {
		go func() {
			a.Swap(func(x int) int {
				return x + 1
			})
			w.Done()
		}()
	}
	w.Wait()

	if a.Get() != 100000 {
		t.Errorf("a.Get() != 100000")
		return
	}
}
