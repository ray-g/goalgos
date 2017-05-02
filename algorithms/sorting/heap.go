package sorting

import (
	heap "github.com/ray-g/goalgos/data-structures/heaps/binary-heap"
)

type Int int

func (x Int) Less(than interface{}) bool { return x < than.(Int) }

func HeapSort(v Changeable) {
	h := heap.New(true)
	for i := 0; i < v.Len(); i++ {
		h.Insert(Int(v.Get(i).(int)))
	}

	for i := 0; i < v.Len(); i++ {
		v.Change(i, int(h.Extract().(Int)))
	}
}
